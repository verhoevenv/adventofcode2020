package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type void struct{}

type layout struct {
	activeCubes map[xyz]void
}

type xyz struct {
	x int
	y int
	z int
}

func (me *xyz) equals(other *xyz) bool {
	return me.x == other.x && me.y == other.y && me.z == other.z
}

func makeLayout(layoutStr string) *layout {
	result := make(map[xyz]void)

	for y, line := range strings.Split(layoutStr, "\n") {
		for x, r := range line {
			if r == '#' {
				result[xyz{x, y, 0}] = void{}
			}
		}
	}

	return &layout{result}
}

type neighbours func(loc *xyz) []xyz

var surroundingNeighbours = withCache(calcNeighbours)

func withCache(f neighbours) neighbours {
	var neighboursCache = make(map[xyz][]xyz)

	return func(loc *xyz) []xyz {
		if v, ok := neighboursCache[*loc]; ok {
			return v
		}
		result := f(loc)
		neighboursCache[*loc] = result
		return result
	}
}

func calcNeighbours(loc *xyz) []xyz {
	result := make([]xyz, 0)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			for _, dz := range []int{-1, 0, 1} {
				candidate := xyz{loc.x + dx, loc.y + dy, loc.z + dz}
				if !loc.equals(&candidate) {
					result = append(result, candidate)
				}
			}
		}
	}
	return result
}

func (l *layout) countActive(locs []xyz) int {
	counter := 0
	for _, loc := range locs {
		if _, ok := l.activeCubes[loc]; ok {
			counter++
		}
	}
	return counter
}

func (l *layout) countAllActive() int {
	return len(l.activeCubes)
}

func (l *layout) cycle() {
	pois := make(map[xyz]void)
	for active := range l.activeCubes {
		for _, loc := range surroundingNeighbours(&active) {
			pois[loc] = void{}
		}
	}

	result := make(map[xyz]void)
	for loc := range pois {
		_, isActive := l.activeCubes[loc]
		activeNeightbours := l.countActive(surroundingNeighbours(&loc))

		switch {
		case isActive && (activeNeightbours == 2 || activeNeightbours == 3):
			result[loc] = void{}
		case !isActive && activeNeightbours == 3:
			result[loc] = void{}
		}
	}

	l.activeCubes = result
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	layout := makeLayout(string(input))

	for i := 0; i < 6; i++ {
		layout.cycle()
	}

	fmt.Println(layout.countAllActive())
}
