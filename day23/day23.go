package main

import "fmt"

type game struct {
	cups []int
}

func (g *game) playOneRound() {
	current := g.cups[0]
	pickUp := g.cups[1:4]

	afterPickup := append([]int{current}, g.cups[4:]...)

	destinationLabel := current
	for contains := true; contains; _, contains = index(pickUp, destinationLabel) {
		destinationLabel--
		if destinationLabel <= 0 {
			destinationLabel += len(g.cups)
		}
	}
	destination, _ := index(afterPickup, destinationLabel)

	newOrder := append(afterPickup[0:destination+1], append(pickUp, afterPickup[destination+1:]...)...)
	nextRound := append(newOrder[1:], newOrder[0])

	g.cups = nextRound
}

func (g *game) orderAfter1() []int {
	idx, _ := index(g.cups, 1)

	return append(g.cups[idx+1:], g.cups[0:idx]...)
}

func (g *game) play(rounds int) {
	for i := 0; i < rounds; i++ {
		g.playOneRound()
	}
}

func index(haystack []int, needle int) (int, bool) {
	for i, h := range haystack {
		if h == needle {
			return i, true
		}
	}
	return -1, false
}

func main() {
	game := game{[]int{9, 6, 3, 2, 7, 5, 4, 8, 1}}

	game.play(100)

	fmt.Println(game.orderAfter1())
}
