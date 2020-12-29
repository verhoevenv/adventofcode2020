package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type gamestate struct {
	p1 deck
	p2 deck
}

type game struct {
	gamestate
	previousStates []*gamestate
}

const maxDeckSize = 100

type card int

type deck struct {
	cardBuffer [maxDeckSize]card
	topIdx     int
	endIdx     int
}

func (d *deck) popTopCard() card {
	card := d.cardBuffer[d.topIdx]
	d.topIdx = (d.topIdx + 1) % maxDeckSize
	return card
}

func (d *deck) addCard(c card) {
	d.cardBuffer[d.endIdx] = c
	d.endIdx = (d.endIdx + 1) % maxDeckSize
}

func (d *deck) isEmpty() bool {
	return d.endIdx == d.topIdx
}

func (d *deck) size() int {
	return (d.endIdx + maxDeckSize - d.topIdx) % maxDeckSize
}

func (d *deck) asSlice() []card {
	result := make([]card, 0)
	for i := d.topIdx; i != d.endIdx; i = (i + 1) % maxDeckSize {
		result = append(result, d.cardBuffer[i])
	}
	return result
}

func (d *deck) equals(other *deck) bool {
	for i := 0; ; i++ {
		myIdx := (d.topIdx + i) % maxDeckSize
		otherIdx := (other.topIdx + i) % maxDeckSize
		if myIdx == d.endIdx && otherIdx == other.endIdx {
			return true
		}
		if myIdx == d.endIdx || otherIdx == other.endIdx {
			return false
		}
		if d.cardBuffer[myIdx] != other.cardBuffer[otherIdx] {
			return false
		}
	}
}

func (g *game) playOneRound() {
	p1Card := g.p1.popTopCard()
	p2Card := g.p2.popTopCard()

	if p1Card > p2Card {
		g.p1.addCard(p1Card)
		g.p1.addCard(p2Card)
	} else {
		g.p2.addCard(p2Card)
		g.p2.addCard(p1Card)
	}
}

func (g *game) play() {
	for !g.p1.isEmpty() && !g.p2.isEmpty() {
		g.playOneRound()
	}
}

type player int // 1|2

func (g *game) playRecursive() player {
	for {
		for i := range g.previousStates {
			if g.equalDecks(g.previousStates[i]) {
				return 1
			}
		}

		g.previousStates = append(g.previousStates, g.copy())

		if g.p1.isEmpty() {
			return 2
		}
		if g.p2.isEmpty() {
			return 1
		}

		p1Card := int(g.p1.popTopCard())
		p2Card := int(g.p2.popTopCard())

		var winner player
		if p1Card <= g.p1.size() && p2Card <= g.p2.size() {
			subgame := game{
				*gameStateFromSlice(
					g.p1.asSlice()[0:p1Card],
					g.p2.asSlice()[0:p2Card],
				),
				make([]*gamestate, 0),
			}
			winner = subgame.playRecursive()
		} else {
			if p1Card > p2Card {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			g.p1.addCard(card(p1Card))
			g.p1.addCard(card(p2Card))
		} else {
			g.p2.addCard(card(p2Card))
			g.p2.addCard(card(p1Card))
		}
	}
}

func (g *gamestate) equalDecks(other *gamestate) bool {
	return g.p1.equals(&other.p1) && g.p2.equals(&other.p2)
}

func gameStateFromSlice(p1 []card, p2 []card) *gamestate {
	var p1d deck
	for i, card := range p1 {
		p1d.cardBuffer[i] = card
		p1d.endIdx = i + 1
	}

	var p2d deck
	for i, card := range p2 {
		p2d.cardBuffer[i] = card
		p2d.endIdx = i + 1
	}

	return &gamestate{p1d, p2d}
}

func (g *gamestate) copy() *gamestate {
	return gameStateFromSlice(g.p1.asSlice(), g.p2.asSlice())
}

func (g *game) scoreWinner() int {
	var winner *deck
	if g.p1.isEmpty() {
		winner = &g.p2
	} else {
		winner = &g.p1
	}

	cards := winner.asSlice()

	score := 0
	multiplier := len(cards)

	for _, c := range cards {
		score += int(c) * multiplier
		multiplier--
	}

	return score
}

func parse(in string) *game {
	playersStr := strings.Split(in, "\n\n")

	var p1 deck
	for i, cardStr := range strings.Split(playersStr[0], "\n")[1:] {
		p1.cardBuffer[i] = card(unsafeAtoi(cardStr))
		p1.endIdx = i + 1
	}

	var p2 deck
	for i, cardStr := range strings.Split(playersStr[1], "\n")[1:] {
		p2.cardBuffer[i] = card(unsafeAtoi(cardStr))
		p2.endIdx = i + 1
	}

	return &game{gamestate{p1, p2}, make([]*gamestate, 0)}
}

func unsafeAtoi(in string) int {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return v
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	game := parse(string(input))
	game.playRecursive()

	scoreOfWinner := game.scoreWinner()

	fmt.Println(scoreOfWinner)
}
