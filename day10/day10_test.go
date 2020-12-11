package main

import (
	"reflect"
	"testing"
)

var smallData = []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

func TestChain(t *testing.T) {
	adaps := makeAdapters(smallData)

	result := adaps.chainTo()

	expected := []int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Chain was incorrect, got: %v, want: %v.",
			result, expected)
	}

}

func TestCountDifferences(t *testing.T) {
	adaps := makeAdapters(smallData)

	result := adaps.countDifferencesInChain()

	expected := map[int]int{1: 7, 3: 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("countDifferencesInChain was incorrect, got: %v, want: %v.",
			result, expected)
	}

}

var largerData = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49,
	45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}

func TestCountDifferencesLarger(t *testing.T) {
	adaps := makeAdapters(largerData)

	result := adaps.countDifferencesInChain()

	expected := map[int]int{1: 22, 3: 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("countDifferencesInChain was incorrect, got: %v, want: %v.",
			result, expected)
	}

}

func TestCountCombinations(t *testing.T) {
	adaps := makeAdapters(smallData)

	result := adaps.countCombinations()

	if result != 8 {
		t.Errorf("TestCountCombinations was incorrect, got: %v, want: %v.",
			result, 8)
	}
}

func TestCountCombinationsLarger(t *testing.T) {
	adaps := makeAdapters(largerData)

	result := adaps.countCombinations()

	if result != 19208 {
		t.Errorf("TestCountCombinations was incorrect, got: %v, want: %v.",
			result, 19208)
	}
}
