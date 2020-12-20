package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type void struct{}

type layout struct {
	activeCubes map[locable]void
}

type locable interface {
	calcNeighbours() []locable
}

type xyz struct {
	x int
	y int
	z int
}

func (loc xyz) calcNeighbours() []locable {
	result := make([]locable, 0)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			for _, dz := range []int{-1, 0, 1} {
				candidate := xyz{loc.x + dx, loc.y + dy, loc.z + dz}
				if !loc.equals(candidate) {
					result = append(result, candidate)
				}
			}
		}
	}
	return result
}

func (loc xyz) equals(other xyz) bool {
	return loc.x == other.x && loc.y == other.y && loc.z == other.z
}

type xyzw struct {
	x int
	y int
	z int
	w int
}

func (loc xyzw) calcNeighbours() []locable {
	result := make([]locable, 0)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			for _, dz := range []int{-1, 0, 1} {
				for _, dw := range []int{-1, 0, 1} {
					candidate := xyzw{loc.x + dx, loc.y + dy, loc.z + dz, loc.w + dw}
					if !loc.equals(candidate) {
						result = append(result, candidate)
					}
				}
			}
		}
	}
	return result
}

func (loc xyzw) equals(other xyzw) bool {
	return loc.x == other.x && loc.y == other.y && loc.z == other.z && loc.w == other.w
}

func makeLayout(layoutStr string, dim int) *layout {
	result := make(map[locable]void)

	for y, line := range strings.Split(layoutStr, "\n") {
		for x, r := range line {
			if r == '#' {
				if dim == 3 {
					result[xyz{x, y, 0}] = void{}
				} else {
					result[xyzw{x, y, 0, 0}] = void{}
				}
			}
		}
	}

	return &layout{result}
}

type neighbours func(loc locable) []locable

var surroundingNeighbours neighbours = func() neighbours {
	var neighboursCache = make(map[locable][]locable)

	return func(loc locable) []locable {
		if v, ok := neighboursCache[loc]; ok {
			return v
		}
		result := loc.calcNeighbours()
		neighboursCache[loc] = result
		return result
	}
}()

func (l *layout) countActive(locs []locable) int {
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
	pois := make(map[locable]void)
	for active := range l.activeCubes {
		for _, loc := range surroundingNeighbours(active) {
			pois[loc] = void{}
		}
	}

	result := make(map[locable]void)
	for loc := range pois {
		_, isActive := l.activeCubes[loc]
		activeNeightbours := l.countActive(surroundingNeighbours(loc))

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

	layout := makeLayout(string(input), 4)

	for i := 0; i < 6; i++ {
		layout.cycle()
	}

	fmt.Println(layout.countAllActive())
}
