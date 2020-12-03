package main

import "testing"

var world = [][]byte{
	[]byte("..##......."),
	[]byte("#...#...#.."),
	[]byte(".#....#..#."),
	[]byte("..#.#...#.#"),
	[]byte(".#...##..#."),
	[]byte("..#.##....."),
	[]byte(".#.#.#....#"),
	[]byte(".#........#"),
	[]byte("#.##...#..."),
	[]byte("#...##....#"),
	[]byte(".#..#...#.#"),
}

func TestAtWithRepeatedMap(t *testing.T) {
	tables := []struct {
		pos      Position
		expected byte
	}{
		{Position{0, 0}, '.'},
		{Position{0, 1}, '#'},
		{Position{2, 0}, '#'},
		{Position{11, 0}, '.'},
		{Position{11, 1}, '#'},
		{Position{13, 0}, '#'},
	}

	m := NewMap(world)

	for _, table := range tables {
		result := AtWithRepeatedMap(&m, &table.pos)
		if result != table.expected {
			t.Errorf("Validation of (%v,%v) was incorrect, got: %q, want: %q.",
				table.pos.x, table.pos.y, result, table.expected)
		}
	}

}

func TestTraverseAndCountTrees(t *testing.T) {
	m := NewMap(world)

	result := TraverseAndCountTrees(&m)

	if result != 7 {
		t.Errorf("TraverseAndCountTrees incorrect, got: %v, want %v.",
			result, 7)
	}

}
