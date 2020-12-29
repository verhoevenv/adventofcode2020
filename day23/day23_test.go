package main

import (
	"reflect"
	"testing"
)

func TestPlayOneRound(t *testing.T) {
	game := game{[]int{3, 8, 9, 1, 2, 5, 4, 6, 7}}

	game.playOneRound()

	expectedOrder := []int{5, 4, 6, 7, 3, 2, 8, 9}
	if !reflect.DeepEqual(game.orderAfter1(), expectedOrder) {
		t.Errorf("playOneRound was incorrect, got: %v, want: %v.",
			game.orderAfter1(), expectedOrder)
	}
}

func TestPlay100Rounds(t *testing.T) {
	game := game{[]int{3, 8, 9, 1, 2, 5, 4, 6, 7}}

	game.play(100)

	expectedOrder := []int{6, 7, 3, 8, 4, 5, 2, 9}
	if !reflect.DeepEqual(game.orderAfter1(), expectedOrder) {
		t.Errorf("play100Rounds was incorrect, got: %v, want: %v.",
			game.orderAfter1(), expectedOrder)
	}
}
