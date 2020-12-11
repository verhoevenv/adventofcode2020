package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type layout struct {
	grid map[xy]state
	size xy
}

type state rune

type xy struct {
	x int
	y int
}

func makeLayout(layoutStr string) *layout {
	result := make(map[xy]state)

	maxX := 0
	maxY := 0
	for y, line := range strings.Split(layoutStr, "\n") {
		maxY++
		maxX = 0
		for x, r := range line {
			maxX++
			result[xy{x, y}] = state(r)
		}
	}

	return &layout{result, xy{maxX, maxY}}
}

type rules struct {
	neighbourRule neighbouringSeats
	tolerance     int
}

var part1 = rules{
	withCache(surroundingNeighbours),
	4,
}

var part2 = rules{
	withCache(visibleNeighbours),
	5,
}

type neighbouringSeats func(l *layout, loc *xy) []xy

func withCache(f neighbouringSeats) neighbouringSeats {
	var neighboursCache = make(map[xy][]xy)

	return func(l *layout, loc *xy) []xy {
		if v, ok := neighboursCache[*loc]; ok {
			return v
		}
		result := f(l, loc)
		neighboursCache[*loc] = result
		return result
	}
}

func surroundingNeighbours(l *layout, loc *xy) []xy {
	result := make([]xy, 0)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			candidate := xy{loc.x + dx, loc.y + dy}
			if !reflect.DeepEqual(*loc, candidate) &&
				candidate.x >= 0 && candidate.x < l.size.x &&
				candidate.y >= 0 && candidate.y < l.size.y {
				result = append(result, candidate)
			}
		}
	}
	return result
}

func visibleNeighbours(l *layout, loc *xy) []xy {
	result := make([]xy, 0)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if !(dx == 0 && dy == 0) {
				dist := 0
				stop := false
				for !stop {
					dist++
					candidate := xy{loc.x + dx*dist, loc.y + dy*dist}
					if candidate.x < 0 || candidate.x >= l.size.x ||
						candidate.y < 0 || candidate.y >= l.size.y {
						stop = true
					} else if l.grid[candidate] == '#' ||
						l.grid[candidate] == 'L' {
						result = append(result, candidate)
						stop = true
					}
				}
			}
		}
	}
	return result
}

func (l *layout) count(s state, locs []xy) int {
	counter := 0
	for _, loc := range locs {
		if l.grid[loc] == s {
			counter++
		}
	}
	return counter
}

func (l *layout) countAll(s state) int {
	counter := 0
	for _, ls := range l.grid {
		if ls == s {
			counter++
		}
	}
	return counter
}

func (l *layout) round(rules rules) {
	result := make(map[xy]state)

	for loc, s := range l.grid {
		switch {
		case s == 'L' && l.count('#', rules.neighbourRule(l, &loc)) == 0:
			result[loc] = '#'
		case s == '#' && l.count('#', rules.neighbourRule(l, &loc)) >= rules.tolerance:
			result[loc] = 'L'
		default:
			result[loc] = s
		}
	}

	l.grid = result
}

func (l *layout) untilStable(rules rules) {
	for {
		old := l.grid
		l.round(rules)
		if reflect.DeepEqual(old, l.grid) {
			break
		}
	}
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	layout := makeLayout(string(input))
	layout.untilStable(part2)

	fmt.Println(layout.countAll('#'))
}
