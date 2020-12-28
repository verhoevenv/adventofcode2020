package main

import (
	"strings"
	"testing"
)

func TestArrangeTiles(t *testing.T) {
	tiles := parse(tiles)

	arrangement := arrangeTiles(tiles)

	expectedCornerTilesMult := 1951 * 3079 * 2971 * 1171

	result := 1
	result *= arrangement[0][0].tile.id
	result *= arrangement[2][0].tile.id
	result *= arrangement[0][2].tile.id
	result *= arrangement[2][2].tile.id

	if result != expectedCornerTilesMult {
		t.Errorf("Arrangement was incorrect, got: %v, want: %v.",
			arrangement, expectedCornerTilesMult)
	}
}

func TestBorder(t *testing.T) {
	tile := parseTile(`Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`)

	tables := []struct {
		side        side
		orientation orientation
		expected    border
	}{
		{top, orientation{false, original}, 0b0011010010},
		{left, orientation{false, original}, 0b0111110010},

		{left, orientation{false, rotate90}, 0b0100101100},
		{bot, orientation{false, rotate180}, 0b0100101100},
		{right, orientation{false, rotate270}, 0b0011010010},

		{right, orientation{true, original}, 0b0111110010},
		{left, orientation{true, original}, 0b0001011001},

		{top, orientation{true, original}, 0b0100101100},
		{top, orientation{true, rotate90}, 0b00111110010},
		{right, orientation{true, rotate90}, 0b0011100111},
	}

	for _, table := range tables {
		oriented := orientedTile{&tile, table.orientation}
		result := borderOf(oriented, table.side)
		if result != table.expected {
			t.Errorf("Border of side %v, orientation %v was incorrect, got: %v, want: %v.",
				table.side, table.orientation, result, table.expected)
		}
	}

}

var properlyFlippedImage = asSnapshot(`.####...#####..#...###..
#####..#..#.#.####..#.#.
.#.#...#.###...#.##.##..
#.#.##.###.#.##.##.#####
..##.###.####..#.####.##
...#.#..##.##...#..#..##
#.##.#..#.#..#..##.#.#..
.###.##.....#...###.#...
#.####.#.#....##.#..#.#.
##...#..#....#..#...####
..#.##...###..#.#####..#
....#.##.#.#####....#...
..##.##.###.....#.##..#.
#...#...###..####....##.
.#.##...#.##.#.#.###...#
#.###.#..####...##..#...
#.###...#.##...#.######.
.###.###.#######..#####.
..##.#..#..#.#######.###
#.#..##.########..#..##.
#.#####..#.#...##..#....
#....##..#.#########..##
#...#.....#..##...###.##
#..###....##.#...##.##.#`)

func asSnapshot(in string) snapshot {
	result := make([][]bool, 0)

	for i, line := range strings.Split(in, "\n") {
		result = append(result, make([]bool, len(line)))
		for j, c := range line {
			if c == '#' {
				result[i][j] = true
			}
		}
	}

	return snapshot(result)
}

func TestCountMonsters(t *testing.T) {
	result := countMonsters(&properlyFlippedImage)

	if result != 2 {
		t.Errorf("Expected 2 sea monsters, found %v",
			result)
	}
}

func TestMakeOrientedSnapshot(t *testing.T) {
	tiles := parse(tiles)
	arrangement := arrangeTiles(tiles)

	result := makeOrientedSnapshot(arrangement)

	for y := 0; y < len(properlyFlippedImage); y++ {
		for x := 0; x < len(properlyFlippedImage); x++ {
			if result.at(xy{x, y}) != properlyFlippedImage.at(xy{x, y}) {
				t.Errorf("Snapshot at (%v, %v) was incorrect, got: %v, want: %v.",
					x, y, result.at(xy{x, y}), properlyFlippedImage.at(xy{x, y}))
			}
		}
	}

}

func TestDetermineRougness(t *testing.T) {
	tiles := parse(tiles)
	arrangement := arrangeTiles(tiles)
	snapshot := makeOrientedSnapshot(arrangement)

	result := determineRoughness(snapshot)

	if result != 273 {
		t.Errorf("Expected 273 roughness, found %v",
			result)
	}

}

var tiles = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`
