package main

import "fmt"

const maxCups = 1000000

type game struct {
	nextCup [maxCups + 1]int //read as: map[label]labelidx
	numCups int
	current int
}

func makeOrdering(initial []int, numCups int) []int {
	result := make([]int, numCups)
	for i := 0; i < numCups; i++ {
		result[i] = i + 1
	}
	for i := 0; i < len(initial); i++ {
		result[i] = initial[i]
	}
	return result
}

func makeGame(ordering []int) game {
	var result game

	for i := 0; i < len(ordering)-1; i++ {
		result.nextCup[ordering[i]] = ordering[i+1]
	}

	result.nextCup[ordering[len(ordering)-1]] = ordering[0]

	result.numCups = len(ordering)
	result.current = ordering[0]

	return result
}

func (g *game) playOneRound() {
	var pickUp [3]int

	pickUp[0] = g.nextCup[g.current]
	pickUp[1] = g.nextCup[pickUp[0]]
	pickUp[2] = g.nextCup[pickUp[1]]

	g.nextCup[g.current] = g.nextCup[pickUp[2]]

	destination := g.current - 1
	if destination <= 0 {
		destination += g.numCups
	}
	for destination == pickUp[0] ||
		destination == pickUp[1] ||
		destination == pickUp[2] {
		destination--
		if destination <= 0 {
			destination += g.numCups
		}
	}
	afterDestination := g.nextCup[destination]

	g.nextCup[destination] = pickUp[0]
	g.nextCup[pickUp[2]] = afterDestination

	g.current = g.nextCup[g.current]
}

func (g *game) orderAfter1() []int {
	result := make([]int, 0)

	cup := g.nextCup[1]
	for i := 0; i < g.numCups-1; i++ {
		result = append(result, cup)
		cup = g.nextCup[cup]
	}

	return result
}

func (g *game) twoCupsAfter1() (int, int) {
	return g.nextCup[1], g.nextCup[g.nextCup[1]]
}

func (g *game) play(rounds int) {
	for i := 0; i < rounds; i++ {
		g.playOneRound()
	}
}

func main() {
	ordering := makeOrdering([]int{9, 6, 3, 2, 7, 5, 4, 8, 1}, 1000000)
	game := makeGame(ordering)

	game.play(10000000)

	one, two := game.twoCupsAfter1()

	fmt.Println(one * two)
}
