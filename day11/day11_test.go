package main

import (
	"reflect"
	"testing"
)

var seatings = []string{
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
	layout := makeLayout(seatings[0])

	for i := 1; i < len(seatings); i++ {
		layout.round()

		expected := makeLayout(seatings[i])

		if !reflect.DeepEqual(layout, expected) {
			t.Errorf("Round from %v to %v was incorrect, got: %v, want: %v.",
				i-1, i, layout, expected)
		}
	}
}

func TestUntilStable(t *testing.T) {
	layout := makeLayout(seatings[0])

	layout.untilStable()

	expected := makeLayout(seatings[len(seatings)-1])

	if !reflect.DeepEqual(layout, expected) {
		t.Errorf("UntilStable was incorrect, got: %v, want: %v.",
			layout, expected)
	}

}
