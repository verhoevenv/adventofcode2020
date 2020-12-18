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

func TestFilterInvalid(t *testing.T) {
	notes := parse(notesStr)

	result := filterOutInvalidTickets(notes.rules, notes.nearby)

	expected := []ticket{ticket{7, 3, 47}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("filterOutInvalidTickets was incorrect, got: %v, want: %v.",
			result, expected)
	}
}

var notesStr2 = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

func TestDetermineOrdering(t *testing.T) {
	notes := parse(notesStr2)

	result := determineOrdering(notes.rules, notes.nearby)

	expected := []string{"row", "class", "seat"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("determineOrdering was incorrect, got: %v, want: %v.",
			result, expected)
	}
}
