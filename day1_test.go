package main

import "testing"

func TestPair(t *testing.T) {
	first, second := PairThatSumsTo2020([]int{1721, 979, 366, 299, 675, 1456})
	if first != 1721 && second != 299 {
		t.Errorf("Elems was incorrect, got: %d and %d, want: %d and %d.",
			first, second, 1721, 299)
	}
}

func TestTriplet(t *testing.T) {
	first, second, third := TripletThatSumsTo2020([]int{1721, 979, 366, 299, 675, 1456})
	if first != 979 && second != 366 && third != 675 {
		t.Errorf("Elems was incorrect, got: %d, %d and %d, want: %d, %d and %d.",
			first, second, third, 979, 366, 675)
	}
}
