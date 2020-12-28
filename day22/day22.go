package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type game struct {
	p1 deck
	p2 deck
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

func (d *deck) asSlice() []card {
	result := make([]card, 0)
	for i := d.topIdx; i != d.endIdx; i = (i + 1) % maxDeckSize {
		result = append(result, d.cardBuffer[i])
	}
	return result
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

	return &game{p1, p2}
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
	game.play()

	scoreOfWinner := game.scoreWinner()

	fmt.Println(scoreOfWinner)
}
