package main

import (
	"reflect"
	"testing"
)

var seatingsPart1 = []string{
	`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,

	`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,

	`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`,

	`#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##`,

	`#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##`,

	`#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`,
}

func TestRound(t *testing.T) {
	layout := makeLayout(seatingsPart1[0])

	for i := 1; i < len(seatingsPart1); i++ {
		layout.round(part1)

		expected := makeLayout(seatingsPart1[i])

		if !reflect.DeepEqual(layout, expected) {
			t.Errorf("Round from %v to %v was incorrect, got: %v, want: %v.",
				i-1, i, layout, expected)
		}
	}
}

func TestUntilStable(t *testing.T) {
	layout := makeLayout(seatingsPart1[0])

	layout.untilStable(part1)

	expected := makeLayout(seatingsPart1[len(seatingsPart1)-1])

	if !reflect.DeepEqual(layout, expected) {
		t.Errorf("UntilStable was incorrect, got: %v, want: %v.",
			layout, expected)
	}

}

var seatingsPart2 = []string{
	`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,

	`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,

	`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`,

	`#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`,

	`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`,

	`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`,

	`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`,
}

func TestVisibleNeighbours(t *testing.T) {
	layout := makeLayout(seatingsPart2[0])

	result := visibleNeighbours(layout, &xy{0, 0})

	expected := []xy{xy{0, 1}, xy{2, 0}, xy{1, 1}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("visibleNeighbours incorrect, got: %v, want: %v.",
			result, expected)
	}
}

func TestRound2(t *testing.T) {
	layout := makeLayout(seatingsPart2[0])

	for i := 1; i < len(seatingsPart2); i++ {
		layout.round(part2)

		expected := makeLayout(seatingsPart2[i])

		if !reflect.DeepEqual(layout, expected) {
			t.Errorf("Round from %v to %v was incorrect, got: %v, want: %v.",
				i-1, i, layout, expected)
		}
	}
}

func TestUntilStable2(t *testing.T) {
	layout := makeLayout(seatingsPart2[0])

	layout.untilStable(part2)

	expected := makeLayout(seatingsPart2[len(seatingsPart2)-1])

	if !reflect.DeepEqual(layout, expected) {
		t.Errorf("UntilStable was incorrect, got: %v, want: %v.",
			layout, expected)
	}

}
