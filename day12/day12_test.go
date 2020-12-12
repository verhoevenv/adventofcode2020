package main

import (
	"reflect"
	"testing"
)

var instuctionsStr = `F10
N3
F7
R90
F11`

func TestNavigate(t *testing.T) {
	ship := makeShip()
	instructions := parse(instuctionsStr)

	ship.navigate(instructions)

	expected := xy{17, -8}
	if !reflect.DeepEqual(ship.pos, expected) {
		t.Errorf("Navigate was incorrect, got: %v, want: %v.",
			ship.pos, expected)
	}
}

func TestMDist(t *testing.T) {
	given := xy{17, -8}

	result := mDist(given)

	expected := 25
	if result != expected {
		t.Errorf("Manhattan distance was incorrect, got: %v, want: %v.",
			result, expected)
	}
}
