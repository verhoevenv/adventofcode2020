package main

import (
	"reflect"
	"testing"
)

var initialCubes = `.#.
..#
###`

func TestCalcNeighbours(t *testing.T) {
	given := xyz{1, 2, 3}

	result := given.calcNeighbours()

	expected := []locable{
		xyz{0, 1, 2}, xyz{0, 1, 3}, xyz{0, 1, 4},
		xyz{0, 2, 2}, xyz{0, 2, 3}, xyz{0, 2, 4},
		xyz{0, 3, 2}, xyz{0, 3, 3}, xyz{0, 3, 4},
		xyz{1, 1, 2}, xyz{1, 1, 3}, xyz{1, 1, 4},
		xyz{1, 2, 2}, xyz{1, 2, 4},
		xyz{1, 3, 2}, xyz{1, 3, 3}, xyz{1, 3, 4},
		xyz{2, 1, 2}, xyz{2, 1, 3}, xyz{2, 1, 4},
		xyz{2, 2, 2}, xyz{2, 2, 3}, xyz{2, 2, 4},
		xyz{2, 3, 2}, xyz{2, 3, 3}, xyz{2, 3, 4},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("CalcNeighbours was incorrect, got: %v, want: %v.",
			result, expected)
	}

}

func TestCalcNeighbours4Dim(t *testing.T) {
	given := xyzw{1, 2, 3, 4}

	result := given.calcNeighbours()

	if len(result) != 80 {
		t.Errorf("CalcNeighbours was incorrect, got: %v, want: %v elems.",
			result, 80)
	}

}

func TestCountAllActive(t *testing.T) {
	layout := makeLayout(initialCubes, 3)

	totalActiveCubes := map[int]int{
		1: 11,
		2: 21,
		3: 38,
		6: 112,
	}

	for i := 1; i <= 6; i++ {
		layout.cycle()

		if expected, ok := totalActiveCubes[i]; ok {
			result := layout.countAllActive()
			if expected != result {
				t.Errorf("CountAllActive was incorrect, got: %v, want: %v.",
					result, expected)
			}
		}
	}
}

func TestCountAllActive4Dim(t *testing.T) {
	layout := makeLayout(initialCubes, 4)

	totalActiveCubes := map[int]int{
		1: 29,
		2: 60,
		6: 848,
	}

	for i := 1; i <= 6; i++ {
		layout.cycle()

		if expected, ok := totalActiveCubes[i]; ok {
			result := layout.countAllActive()
			if expected != result {
				t.Errorf("CountAllActive in 4 dimensions was incorrect, got: %v, want: %v.",
					result, expected)
			}
		}
	}
}
