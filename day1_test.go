package main

import "testing"

func Test(t *testing.T) {
	first, second := ElemsThatSumTo2020([]int{1721, 979, 366, 299, 675, 1456})
	if first != 1721 && second != 299 {
		t.Errorf("Elems was incorrect, got: %d and %d, want: %d and %d.",
			first, second, 1721, 299)
	}
}
