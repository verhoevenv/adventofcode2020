package main

import (
	"bufio"
	"fmt"
	"os"
)

type Map struct {
	terrain [][]byte
	width   int
	height  int
}

const Tree = '#'

func NewMap(terrain [][]byte) Map {
	return Map{terrain, len(terrain[0]), len(terrain)}
}

type Position struct {
	x int
	y int
}

func AtWithRepeatedMap(m *Map, pos *Position) byte {
	return m.terrain[pos.y][pos.x%m.width]
}

func TraverseAndCountTrees(m *Map) int {
	numTrees := 0
	pos := Position{0, 0}

	for pos.y < m.height {
		if AtWithRepeatedMap(m, &pos) == Tree {
			numTrees++
		}

		pos.x += 3
		pos.y += 1
	}

	return numTrees
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	input := make([][]byte, 0)
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}

	m := NewMap(input)

	fmt.Println(TraverseAndCountTrees(&m))
}
