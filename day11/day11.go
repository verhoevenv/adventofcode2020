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

var neighboursCache = make(map[xy][]xy)

func (l *layout) neighbours(loc *xy) []xy {
	if v, ok := neighboursCache[*loc]; ok {
		return v
	}
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
	neighboursCache[*loc] = result
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

func (l *layout) round() {
	result := make(map[xy]state)

	for loc, s := range l.grid {
		switch {
		case s == 'L' && l.count('#', l.neighbours(&loc)) == 0:
			result[loc] = '#'
		case s == '#' && l.count('#', l.neighbours(&loc)) >= 4:
			result[loc] = 'L'
		default:
			result[loc] = s
		}
	}

	l.grid = result
}

func (l *layout) untilStable() {
	for {
		old := l.grid
		l.round()
		if reflect.DeepEqual(old, l.grid) {
			break
		}
	}
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	layout := makeLayout(string(input))
	layout.untilStable()

	fmt.Println(layout.countAll('#'))
}
