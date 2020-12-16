package main

import (
	"testing"
)

func TestFindSubsequentDepartTimes(t *testing.T) {
	tables := []struct {
		given    string
		expected uint64
	}{
		{"7,13,x,x,59,x,31,19", 1068781},
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	}

	for _, table := range tables {
		result := findFindSubsequentDepartTimes(table.given)
		if result != table.expected {
			t.Errorf("FindSubsequentDepartTimes of %v was incorrect, got: %v, want: %v.",
				table.given, result, table.expected)
		}
	}

}
