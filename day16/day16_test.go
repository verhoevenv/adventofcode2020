package main

import (
	"reflect"
	"testing"
)

var notesStr = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestInvalid(t *testing.T) {
	notes := parse(notesStr)

	result := findCompletelyInvalidFields(notes.rules, notes.nearby)

	expected := []int{4, 55, 12}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("findCompletelyInvalidFields was incorrect, got: %v, want: %v.",
			result, expected)
	}
}
