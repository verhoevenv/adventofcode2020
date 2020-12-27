package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type xy struct {
	x int
	y int
}

const tileSize = 10

type tile struct {
	id   int
	bits [tileSize][tileSize]bool
}

type orientedTile struct {
	tile        *tile
	orientation orientation
}

type side int

const (
	top side = iota
	right
	bot
	left
)

var allSides = []side{top, right, bot, left}

type rotation int

type orientation struct {
	flip   bool //flip horizontally, so along vertical axis, apply flip first
	rotate rotation
}

var allOrientations = []orientation{
	orientation{false, original},
	orientation{false, rotate90},
	orientation{false, rotate180},
	orientation{false, rotate270},
	orientation{true, original},
	orientation{true, rotate90},
	orientation{true, rotate180},
	orientation{true, rotate270},
}

const (
	original rotation = iota
	rotate90
	rotate180
	rotate270
)

type borderReading struct {
	side      side
	clockwise bool
}

func originalSide(s side, o orientation) borderReading {
	side := side((int(s) + int(o.rotate)) % 4)
	readingDir := s == top || s == right

	if o.flip {
		if side == left {
			side = right
		} else if side == right {
			side = left
		}
		readingDir = !readingDir
	}

	return borderReading{side, readingDir}
}

func borderOf(t orientedTile, s side) border {
	return borderOfTile(t.tile, originalSide(s, t.orientation))
}

func borderOfTile(t *tile, b borderReading) border {
	result := 0
	for i := 0; i < tileSize; i++ {
		result = result << 1

		var loc xy
		if b.clockwise {
			switch b.side {
			case top:
				loc = xy{i, 0}
			case right:
				loc = xy{tileSize - 1, i}
			case bot:
				loc = xy{tileSize - 1 - i, tileSize - 1}
			case left:
				loc = xy{0, tileSize - 1 - i}
			}
		} else {
			switch b.side {
			case top:
				loc = xy{tileSize - 1 - i, 0}
			case right:
				loc = xy{tileSize - 1, tileSize - 1 - i}
			case bot:
				loc = xy{i, tileSize - 1}
			case left:
				loc = xy{0, i}
			}
		}

		if t.bits[loc.y][loc.x] {
			result++
		}
	}
	return border(result)
}

type border int

type searchState struct {
	arrangement   [][]*orientedTile
	nextLoc       xy
	unpickedTiles []*tile
}

func arrangeTiles(tiles []tile) [][]*orientedTile {
	sidesToTileIDs := make(map[side]map[border][]orientedTile)
	for _, side := range allSides {
		sidesToTileIDs[side] = make(map[border][]orientedTile)
	}

	for i := range tiles {
		for _, orientation := range allOrientations {
			for _, side := range allSides {
				border := borderOfTile(&tiles[i], originalSide(side, orientation))
				orientedTile := orientedTile{&tiles[i], orientation}
				sidesToTileIDs[side][border] = append(sidesToTileIDs[side][border], orientedTile)
			}
		}
	}

	tilePts := make([]*tile, len(tiles))
	for i := range tiles {
		tilePts[i] = &tiles[i]
	}

	sizeOfArrangement := int(math.Sqrt(float64(len(tiles))))
	emptyArrangement := make([][]*orientedTile, sizeOfArrangement)
	for i := 0; i*i < len(tiles); i++ {
		emptyArrangement[i] = make([]*orientedTile, sizeOfArrangement)
	}
	toSearch := make([]searchState, 0)
	toSearch = append(toSearch, searchState{
		emptyArrangement,
		xy{0, 0},
		tilePts,
	})

	for len(toSearch) > 0 {

		toExpand := toSearch[len(toSearch)-1]
		newStates := expandLast(toExpand)

		for i := range newStates {
			propagate(&newStates[i])

			if allNil(newStates[i].unpickedTiles) {
				return newStates[i].arrangement
			}
		}

		toSearch = append(toSearch[0:len(toSearch)-1], newStates...)
	}

	panic("no solution found")
}

func allNil(arr []*tile) bool {
	for _, v := range arr {
		if v != nil {
			return false
		}
	}
	return true
}

func expandLast(lastState searchState) []searchState {
	newStates := make([]searchState, 0)

	for tileIdx, nextTile := range lastState.unpickedTiles {
		if nextTile == nil {
			continue
		}
		nextUnpickTiles := make([]*tile, len(lastState.unpickedTiles))
		copy(nextUnpickTiles, lastState.unpickedTiles)
		nextUnpickTiles[tileIdx] = nil
		for _, orientation := range allOrientations {
			nextArrangement := make([][]*orientedTile, len(lastState.arrangement))
			for i := range lastState.arrangement {
				nextArrangement[i] = make([]*orientedTile, len(lastState.arrangement[i]))
				copy(nextArrangement[i], lastState.arrangement[i])
			}
			nextArrangement[lastState.nextLoc.y][lastState.nextLoc.x] = &orientedTile{nextTile, orientation}
			nextLoc := lastState.nextLoc
			nextLoc.x++
			if nextLoc.x >= len(lastState.arrangement) {
				nextLoc.x = 0
				nextLoc.y++
			}
			proposedNewState := searchState{
				nextArrangement,
				nextLoc,
				nextUnpickTiles,
			}
			if !isInvalid(proposedNewState.arrangement) {
				newStates = append(newStates, proposedNewState)
			}
		}
	}

	return newStates
}

func propagate(s *searchState) {
	//TODO if too slow, implement forward propagation
}

func isInvalid(arrangement [][]*orientedTile) bool {
	for y := len(arrangement) - 1; y >= 0; y-- {
		for x := len(arrangement[y]) - 1; x >= 0; x-- {
			if arrangement[y][x] == nil {
				continue
			}

			if y > 0 {
				borderXY := borderOf(*arrangement[y][x], top)
				borderAbove := borderOf(*arrangement[y-1][x], bot)
				if borderXY != borderAbove {
					return true
				}
			}
			if x > 0 {
				borderXY := borderOf(*arrangement[y][x], left)
				borderLeft := borderOf(*arrangement[y][x-1], right)
				if borderXY != borderLeft {
					return true
				}
			}
		}
	}
	return false
}

func parse(in string) []tile {
	tileStrs := strings.Split(in, "\n\n")

	result := make([]tile, 0)
	for _, tileStr := range tileStrs {
		result = append(result, parseTile(tileStr))
	}

	return result
}

func parseTile(in string) tile {
	var result tile
	tileLines := strings.Split(in, "\n")

	result.id = unsafeAtoi(tileLines[0][5:9])

	for x, tileLine := range tileLines[1:] {
		for y, bit := range tileLine {
			if bit == '#' {
				result.bits[x][y] = true
			}
		}
	}

	return result
}

func unsafeAtoi(in string) int {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return v
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	tiles := parse(string(input))

	arrangement := arrangeTiles(tiles)

	sizeSide := len(arrangement) - 1
	result := 1
	result *= arrangement[0][0].tile.id
	result *= arrangement[sizeSide][0].tile.id
	result *= arrangement[0][sizeSide].tile.id
	result *= arrangement[sizeSide][sizeSide].tile.id

	fmt.Println(result)
}
