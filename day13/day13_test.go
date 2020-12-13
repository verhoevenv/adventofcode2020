package main

import (
	"testing"
)

var notesStr = `939
7,13,x,x,59,x,31,19`

func TestFindNextDepartTime(t *testing.T) {
	schedule := parse(notesStr)

	result, _ := findMinWaitTime(&schedule)

	expected := 5
	if result != expected {
		t.Errorf("FindNextDepartTime distance was incorrect, got: %v, want: %v.",
			result, expected)
	}
}
