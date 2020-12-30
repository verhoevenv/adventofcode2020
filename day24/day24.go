package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type void struct{}

type layout struct {
	black map[locable]void
}

type locable interface {
	calcNeighbours() []locable
}

type hexCoord struct {
	x int
	y int
	z int
}

func (loc hexCoord) calcNeighbours() []locable {
	result := make([]locable, 0)
	for _, dir := range []hexCoord{
		hexCoord{+1, -1, 0}, hexCoord{+1, 0, -1}, hexCoord{0, +1, -1},
		hexCoord{-1, +1, 0}, hexCoord{-1, 0, +1}, hexCoord{0, -1, +1},
	} {
		candidate := hexCoord{loc.x + dir.x, loc.y + dir.y, loc.z + dir.z}
		result = append(result, candidate)
	}
	return result
}

func (loc hexCoord) equals(other hexCoord) bool {
	return loc.x == other.x && loc.y == other.y && loc.z == other.z
}

func makeLayout(hexes []hexCoord) *layout {
	flips := make(map[locable]bool)

	for _, hex := range hexes {
		flips[hex] = !flips[hex]
	}

	result := make(map[locable]void)
	for k, v := range flips {
		if v {
			result[k] = void{}

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

func (l *layout) countBlack(locs []locable) int {
	counter := 0
	for _, loc := range locs {
		if _, ok := l.black[loc]; ok {
			counter++
		}
	}
	return counter
}

func (l *layout) countAllBlack() int {
	return len(l.black)
}

func (l *layout) cycle() {
	pois := make(map[locable]void)
	for black := range l.black {
		for _, loc := range surroundingNeighbours(black) {
			pois[loc] = void{}
		}
	}

	result := make(map[locable]void)
	for loc := range pois {
		_, isBlack := l.black[loc]
		blackNeightbours := l.countBlack(surroundingNeighbours(loc))

		switch {
		case isBlack && (blackNeightbours == 1 || blackNeightbours == 2):
			result[loc] = void{}
		case !isBlack && blackNeightbours == 2:
			result[loc] = void{}
		}
	}

	l.black = result
}

func parseCoordinates(in string) []hexCoord {
	result := make([]hexCoord, 0)

	for _, line := range strings.Split(in, "\n") {
		toRead := line
		hex := hexCoord{0, 0, 0}
		for toRead != "" {
			switch {
			case strings.HasPrefix(toRead, "e"):
				toRead = toRead[1:]
				hex.x++
				hex.y--
			case strings.HasPrefix(toRead, "se"):
				toRead = toRead[2:]
				hex.y--
				hex.z++
			case strings.HasPrefix(toRead, "sw"):
				toRead = toRead[2:]
				hex.x--
				hex.z++
			case strings.HasPrefix(toRead, "w"):
				toRead = toRead[1:]
				hex.x--
				hex.y++
			case strings.HasPrefix(toRead, "nw"):
				toRead = toRead[2:]
				hex.y++
				hex.z--
			case strings.HasPrefix(toRead, "ne"):
				toRead = toRead[2:]
				hex.x++
				hex.z--
			default:
				panic(fmt.Sprintf("cannot parse %v (part of %v)", toRead, line))
			}
		}
		result = append(result, hex)
	}

	return result
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	tiles := parseCoordinates(string(input))
	layout := makeLayout(tiles)

	fmt.Println(layout.countAllBlack())

	for i := 0; i < 100; i++ {
		layout.cycle()
	}

	fmt.Println(layout.countAllBlack())
}
