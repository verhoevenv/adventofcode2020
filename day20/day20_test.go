package main

import (
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

func TestOriginalSide(t *testing.T) {
	tables := []struct {
		side     side
		orient   orientation
		expected borderReading
	}{
		{top, orientation{false, original}, borderReading{top, true}},
		{left, orientation{false, original}, borderReading{left, false}},
		{bot, orientation{false, original}, borderReading{bot, false}},
		{left, orientation{true, original}, borderReading{right, true}},
		{bot, orientation{false, rotate180}, borderReading{top, false}},
		{top, orientation{true, original}, borderReading{top, false}},
		{top, orientation{false, rotate90}, borderReading{right, true}},
		{right, orientation{true, rotate180}, borderReading{right, false}},
		{top, orientation{true, rotate90}, borderReading{left, false}},
		{top, orientation{true, rotate180}, borderReading{bot, false}},
		{bot, orientation{true, rotate180}, borderReading{top, true}},
	}

	for _, table := range tables {
		result := originalSide(table.side, table.orient)
		if result != table.expected {
			t.Errorf("Original side of (%v, %v) was incorrect, got: %v, want: %v.",
				table.side, table.orient, result, table.expected)
		}
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
		reading  borderReading
		expected border
	}{
		{borderReading{top, true}, 0b0011010010},
		{borderReading{top, false}, 0b0100101100},
		{borderReading{left, true}, 0b0100111110},
		{borderReading{left, false}, 0b0111110010},
		{borderReading{right, true}, 0b0001011001},
		{borderReading{right, false}, 0b1001101000},
		{borderReading{bot, true}, 0b1110011100},
		{borderReading{bot, false}, 0b0011100111},
	}

	for _, table := range tables {
		result := borderOfTile(&tile, table.reading)
		if result != table.expected {
			t.Errorf("Border of %v was incorrect, got: %v, want: %v.",
				table.reading, result, table.expected)
		}
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
