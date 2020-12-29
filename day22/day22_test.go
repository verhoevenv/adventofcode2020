package main

import (
	"reflect"
	"testing"
)

var startDecks = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestPlayOneRound(t *testing.T) {
	game := parse(startDecks)

	game.playOneRound()

	expectedP1Deck := []card{2, 6, 3, 1, 9, 5}
	expectedP2Deck := []card{8, 4, 7, 10}
	if !reflect.DeepEqual(game.p1.asSlice(), expectedP1Deck) {
		t.Errorf("playOneRound for p1 was incorrect, got: %v, want: %v.",
			game.p1.asSlice(), expectedP1Deck)
	}
	if !reflect.DeepEqual(game.p2.asSlice(), expectedP2Deck) {
		t.Errorf("playOneRound for p2 was incorrect, got: %v, want: %v.",
			game.p2.asSlice(), expectedP2Deck)
	}
}

func TestPlay(t *testing.T) {
	game := parse(startDecks)

	game.play()

	expectedP1Deck := []card{}
	expectedP2Deck := []card{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}
	if !reflect.DeepEqual(game.p1.asSlice(), expectedP1Deck) {
		t.Errorf("play for p1 was incorrect, got: %v, want: %v.",
			game.p1.asSlice(), expectedP1Deck)
	}
	if !reflect.DeepEqual(game.p2.asSlice(), expectedP2Deck) {
		t.Errorf("play for p2 was incorrect, got: %v, want: %v.",
			game.p2.asSlice(), expectedP2Deck)
	}
}

func TestPlayRecursive(t *testing.T) {
	game := parse(startDecks)

	game.playRecursive()

	expectedP1Deck := []card{}
	expectedP2Deck := []card{7, 5, 6, 2, 4, 1, 10, 8, 9, 3}
	if !reflect.DeepEqual(game.p1.asSlice(), expectedP1Deck) {
		t.Errorf("play for p1 was incorrect, got: %v, want: %v.",
			game.p1.asSlice(), expectedP1Deck)
	}
	if !reflect.DeepEqual(game.p2.asSlice(), expectedP2Deck) {
		t.Errorf("play for p2 was incorrect, got: %v, want: %v.",
			game.p2.asSlice(), expectedP2Deck)
	}
}

func TestScoreWinner(t *testing.T) {
	game := parse(startDecks)
	game.play()

	result := game.scoreWinner()

	if result != 306 {
		t.Errorf("scoreWinner was incorrect, got: %v, want: 306.",
			result)
	}
}
