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
	instructions := parse(instuctionsStr)

	tables := []struct {
		nav    navigable
		endPos xy
	}{
		{makeShip(), xy{17, -8}},
	}

	for _, table := range tables {
		navigateBy(table.nav, instructions)
		result := table.nav.getPos()
		if !reflect.DeepEqual(result, table.endPos) {
			t.Errorf("Navigate of %v was incorrect, got: %v, want: %v.",
				table.nav, result, table.endPos)
		}
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
