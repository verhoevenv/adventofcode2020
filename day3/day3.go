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

type Slope struct {
	dx int
	dy int
}

var ListedSlopes []Slope = []Slope{
	Slope{1, 1},
	Slope{3, 1},
	Slope{5, 1},
	Slope{7, 1},
	Slope{1, 2},
}

func AtWithRepeatedMap(m *Map, pos *Position) byte {
	return m.terrain[pos.y][pos.x%m.width]
}

func TraverseAndCountTrees(m *Map, s *Slope) int {
	numTrees := 0
	pos := Position{0, 0}

	for pos.y < m.height {
		if AtWithRepeatedMap(m, &pos) == Tree {
			numTrees++
		}

		pos.x += s.dx
		pos.y += s.dy
	}

	return numTrees
}

func MultiplyTreesForSlopes(m *Map, slopes []Slope) int {
	multTrees := 1

	for _, slope := range slopes {
		treesForSlope := TraverseAndCountTrees(m, &slope)
		multTrees *= treesForSlope
	}

	return multTrees
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	input := make([][]byte, 0)
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}

	m := NewMap(input)

	fmt.Println(MultiplyTreesForSlopes(&m, ListedSlopes))
}
