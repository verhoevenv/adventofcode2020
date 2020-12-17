package main

import (
	"testing"
)

func TestMemory(t *testing.T) {
	tables := []struct {
		start []int
		n2020 int
	}{
		{[]int{0, 3, 6}, 436},
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 3}, 10},
		{[]int{1, 2, 3}, 27},
		{[]int{2, 3, 1}, 78},
		{[]int{3, 2, 1}, 438},
		{[]int{3, 1, 2}, 1836},
	}

	for _, table := range tables {
		mem := startMemory(table.start)
		result := nTimes(2020, mem)
		if result != table.n2020 {
			t.Errorf("Memory of %v was incorrect, got: %v, want: %v.",
				table.start, result, table.n2020)
		}
	}
}
