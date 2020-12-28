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

type booleanReadable interface {
	at(pos xy) bool
	sideSize() int //assuming square
}

func (t *tile) at(pos xy) bool {
	return t.bits[pos.y][pos.x]
}

func (t *tile) sideSize() int {
	return tileSize
}

func (t *orientedTile) at(pos xy) bool {
	return at(pos, t.orientation, t.tile)
}

func (t *orientedTile) sideSize() int {
	return tileSize
}

func at(pos xy, o orientation, b booleanReadable) bool {
	s := b.sideSize()
	switch o.rotate {
	case rotate90:
		pos = xy{s - pos.y - 1, pos.x}
	case rotate180:
		pos = xy{s - pos.x - 1, s - pos.y - 1}
	case rotate270:
		pos = xy{pos.y, s - pos.x - 1}
	}

	if o.flip {
		pos = xy{s - pos.x - 1, pos.y}
	}

	return b.at(pos)
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

func borderOf(t orientedTile, s side) border {
	result := 0

	for i := 0; i < tileSize; i++ {
		result = result << 1

		var loc xy
		switch s {
		case top:
			loc = xy{i, 0}
		case right:
			loc = xy{tileSize - 1, i}
		case bot:
			loc = xy{i, tileSize - 1}
		case left:
			loc = xy{0, i}
		}
		if t.at(loc) {
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

var seaMonster = strings.Split(`
                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `, "\n")[1:]

const pixelsInMonster = 15

func countMonsters(image booleanReadable) int {
	count := 0

	monsterH := len(seaMonster)
	monsterW := len(seaMonster[0])

	size := image.sideSize()

	for dy := 0; dy < size-monsterH; dy++ {
		for dx := 0; dx < size-monsterW; dx++ {
			if monsterAt(xy{dx, dy}, image) {
				count++
			}
		}
	}
	return count
}

func monsterAt(offset xy, image booleanReadable) bool {
	monsterH := len(seaMonster)
	monsterW := len(seaMonster[0])

	for dy := 0; dy < monsterH; dy++ {
		for dx := 0; dx < monsterW; dx++ {
			if seaMonster[dy][dx] == '#' {
				if !image.at(xy{offset.x + dx, offset.y + dy}) {
					return false
				}
			}
		}
	}

	return true
}

type snapshot [][]bool

func (s *snapshot) at(pos xy) bool {
	return [][]bool(*s)[pos.y][pos.x]
}
func (s *snapshot) sideSize() int {
	return len(*s)
}

type orientedSnapshot struct {
	snapshot    *snapshot
	orientation orientation
}

func (s *orientedSnapshot) at(pos xy) bool {
	return at(pos, s.orientation, s.snapshot)
}
func (s *orientedSnapshot) sideSize() int {
	return s.snapshot.sideSize()
}

type unborderedTile orientedTile

func (s *unborderedTile) at(pos xy) bool {
	return (*orientedTile)(s).at(xy{pos.x + 1, pos.y + 1})
}
func (s *unborderedTile) sideSize() int {
	return (*orientedTile)(s).sideSize() - 2
}

func makeSnapshot(arrangements [][]*orientedTile) snapshot {
	numTiles := len(arrangements)

	withoutBorders := make([][]*unborderedTile, numTiles)
	for i := 0; i < numTiles; i++ {
		withoutBorders[i] = make([]*unborderedTile, numTiles)
		for j := 0; j < numTiles; j++ {
			withoutBorders[i][j] = (*unborderedTile)(arrangements[i][j])
		}
	}

	newTileSize := tileSize - 2
	picSideSize := numTiles * newTileSize

	result := make([][]bool, picSideSize)
	for i := 0; i < picSideSize; i++ {
		result[i] = make([]bool, picSideSize)
	}

	for y := 0; y < picSideSize; y++ {
		for x := 0; x < picSideSize; x++ {
			tileAt := withoutBorders[y/newTileSize][x/newTileSize]
			result[y][x] = tileAt.at(xy{x % newTileSize, y % newTileSize})
		}
	}

	return result
}

func makeOrientedSnapshot(arrangement [][]*orientedTile) *orientedSnapshot {
	shot := makeSnapshot(arrangement)

	for _, o := range allOrientations {
		os := &orientedSnapshot{&shot, o}
		if countMonsters(os) > 0 {
			return os
		}
	}

	panic("No good orientation found!")
}

func determineRoughness(s *orientedSnapshot) int {
	counter := 0

	for x := 0; x < s.sideSize(); x++ {
		for y := 0; y < s.sideSize(); y++ {
			if s.at(xy{x, y}) {
				counter++
			}
		}
	}

	numMonsters := countMonsters(s)

	return counter - (numMonsters * pixelsInMonster)
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	tiles := parse(string(input))

	arrangement := arrangeTiles(tiles)

	sizeSide := len(arrangement) - 1
	multiplier := 1
	multiplier *= arrangement[0][0].tile.id
	multiplier *= arrangement[sizeSide][0].tile.id
	multiplier *= arrangement[0][sizeSide].tile.id
	multiplier *= arrangement[sizeSide][sizeSide].tile.id
	fmt.Println(multiplier)

	snapshot := makeOrientedSnapshot(arrangement)

	roughness := determineRoughness(snapshot)
	fmt.Println(roughness)
}
