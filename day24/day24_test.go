package main

import (
	"testing"
)

var tileList = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestCountFlippedTiles(t *testing.T) {
	tiles := parseCoordinates(tileList)
	layout := makeLayout(tiles)

	result := layout.countAllBlack()

	if result != 10 {
		t.Errorf("Expected 10 tiles flipped, got %v",
			result)
	}
}

func TestCycle(t *testing.T) {
	tables := map[int]int{
		1:   15,
		2:   12,
		3:   25,
		4:   14,
		5:   23,
		6:   28,
		7:   41,
		8:   37,
		9:   49,
		10:  37,
		20:  132,
		30:  259,
		40:  406,
		50:  566,
		60:  788,
		70:  1106,
		80:  1373,
		90:  1844,
		100: 2208,
	}

	tiles := parseCoordinates(tileList)
	layout := makeLayout(tiles)

	for day := 1; day <= 100; day++ {
		layout.cycle()

		result := layout.countAllBlack()

		if expected, present := tables[day]; present && (expected != result) {
			t.Errorf("On day %v, expected %v tiles flipped, got %v",
				day, expected, result)
		}
	}

}
