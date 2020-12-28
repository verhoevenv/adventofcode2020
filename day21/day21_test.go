package main

import (
	"reflect"
	"sort"
	"testing"
)

var list = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestFindInertIngredients(t *testing.T) {
	result, _ := findInertIngredients(list)

	expectedIngredients := []string{"kfcds", "nhms", "sbzzf", "trh"}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	if !reflect.DeepEqual(result, expectedIngredients) {
		t.Errorf("findInertIngredients was incorrect, got: %v, want: %v.",
			result, expectedIngredients)
	}
}

func TestCountIngredients(t *testing.T) {
	ingredients, _ := findInertIngredients(list)

	result := countAll(list, ingredients)

	expectedCount := 5

	if result != expectedCount {
		t.Errorf("countIngredients was incorrect, got: %v, want: %v.",
			result, expectedCount)
	}
}

func TestFindCanonicalList(t *testing.T) {
	_, result := findInertIngredients(list)

	expectedIngredients := []string{"mxmxvkd", "sqjhc", "fvjkl"}

	if !reflect.DeepEqual(result, expectedIngredients) {
		t.Errorf("findInertIngredients was incorrect, got: %v, want: %v.",
			result, expectedIngredients)
	}
}
