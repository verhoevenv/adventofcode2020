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

func TestFindIngredientsWithoutAllergens(t *testing.T) {
	result := findIngredientsWithoutAllergens(list)

	expectedIngredients := []string{"kfcds", "nhms", "sbzzf", "trh"}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	if !reflect.DeepEqual(result, expectedIngredients) {
		t.Errorf("findIngredientsWithoutAllergens was incorrect, got: %v, want: %v.",
			result, expectedIngredients)
	}
}

func TestCountIngredients(t *testing.T) {
	ingredients := findIngredientsWithoutAllergens(list)

	result := countAll(list, ingredients)

	expectedCount := 5

	if result != expectedCount {
		t.Errorf("countIngredients was incorrect, got: %v, want: %v.",
			result, expectedCount)
	}
}
