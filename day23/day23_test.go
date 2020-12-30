package main

import (
	"reflect"
	"testing"
)

func TestPlayOneRound(t *testing.T) {
	game := makeGame([]int{3, 8, 9, 1, 2, 5, 4, 6, 7})

	game.playOneRound()

	expectedOrder := []int{5, 4, 6, 7, 3, 2, 8, 9}
	if !reflect.DeepEqual(game.orderAfter1(), expectedOrder) {
		t.Errorf("playOneRound was incorrect, got: %v, want: %v.",
			game.orderAfter1(), expectedOrder)
	}
}

func TestPlay100Rounds(t *testing.T) {
	game := makeGame([]int{3, 8, 9, 1, 2, 5, 4, 6, 7})

	game.play(100)

	expectedOrder := []int{6, 7, 3, 8, 4, 5, 2, 9}
	if !reflect.DeepEqual(game.orderAfter1(), expectedOrder) {
		t.Errorf("play100Rounds was incorrect, got: %v, want: %v.",
			game.orderAfter1(), expectedOrder)
	}
}

func BenchmarkPlay(b *testing.B) {
	order := makeOrdering([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 1000000)
	game := makeGame(order)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.playOneRound()
	}
}

func TestPlayBigGame(t *testing.T) {
	order := makeOrdering([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 1000000)
	game := makeGame(order)

	game.play(10000000)

	result1, result2 := game.twoCupsAfter1()
	expected1 := 934001
	expected2 := 159792
	if result1 != expected1 {
		t.Errorf("play big game was incorrect for first cup, got: %v, want: %v.",
			result1, expected1)
	}
	if result2 != expected2 {
		t.Errorf("play big game was incorrect for second cup, got: %v, want: %v.",
			result2, expected2)
	}

}
