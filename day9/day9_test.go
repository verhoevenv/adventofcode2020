package main

import (
	"testing"
)

func TestExistsAsSum(t *testing.T) {
	var data = []uint64{20, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 1,
		21, 22, 23, 24, 25}

	xmas := makeXMAS(data, 25)

	tables := []struct {
		nextNum uint64
		valid   bool
	}{
		{26, true},
		{49, true},
		{100, false},
		{50, false},
	}

	for _, table := range tables {
		result := xmas.existsAsSum(table.nextNum)
		if result != table.valid {
			t.Errorf("ExistsAsSum of %v was incorrect, got: %v, want: %v.",
				table.nextNum, result, table.valid)
		}
	}

}

func TestShift(t *testing.T) {
	var data = []uint64{20, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 1,
		21, 22, 23, 24, 25, 45}

	xmas := makeXMAS(data, 25)
	xmas.shift()

	tables := []struct {
		nextNum uint64
		valid   bool
	}{
		{26, true},
		{65, false},
		{64, true},
		{66, true},
	}

	for _, table := range tables {
		result := xmas.existsAsSum(table.nextNum)
		if result != table.valid {
			t.Errorf("ExistsAsSum of %v was incorrect, got: %v, want: %v.",
				table.nextNum, result, table.valid)
		}
	}

}

func TestFindNonSumNumber(t *testing.T) {
	var data = []uint64{35, 20, 15, 25, 47, 40, 62, 55, 65,
		95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

	xmas := makeXMAS(data, 5)
	result := xmas.findNonSumNumber()

	if result != 127 {
		t.Errorf("FindNonSumNumber was incorrect, got: %v, want: %v.",
			result, 127)
	}

}
