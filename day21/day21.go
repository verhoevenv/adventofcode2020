package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type line struct {
	ingredients []string
	allergens   []string
}

var lineRE = regexp.MustCompile(`(.+) \(contains (.+)\)`)

func parseLine(in string) line {
	matches := lineRE.FindStringSubmatch(in)
	ingredients := strings.Split(matches[1], " ")
	allergens := strings.Split(matches[2], ", ")
	return line{ingredients, allergens}
}

func findIngredientsWithoutAllergens(in string) []string {
	allIngredients := make(map[string]bool)
	allAllergens := make(map[string]bool)

	for _, lineStr := range strings.Split(in, "\n") {
		line := parseLine(lineStr)
		for _, ingredient := range line.ingredients {
			allIngredients[ingredient] = true
		}
		for _, allergen := range line.allergens {
			allAllergens[allergen] = true
		}
	}

	canContain := make(map[string]map[string]bool)
	for ingredient := range allIngredients {
		canContain[ingredient] = make(map[string]bool)
		for allergen := range allAllergens {
			canContain[ingredient][allergen] = true
		}
	}

	for _, lineStr := range strings.Split(in, "\n") {
		line := parseLine(lineStr)

		for ingredient := range allIngredients {
			if !contains(line.ingredients, ingredient) {
				for _, allergen := range line.allergens {
					delete(canContain[ingredient], allergen)
				}
			}
		}
	}

	result := make([]string, 0)
	for ingredient, allergens := range canContain {
		if len(allergens) == 0 {
			result = append(result, ingredient)
		}
	}

	return result
}

func contains(haystack []string, needle string) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}
	return false
}

func countAll(list string, words []string) int {
	count := 0

	for _, lineStr := range strings.Split(list, "\n") {
		line := parseLine(lineStr)
		for _, w := range words {
			if contains(line.ingredients, w) {
				count++
			}
		}
	}

	return count
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	ingredients := findIngredientsWithoutAllergens(string(input))
	result := countAll(string(input), ingredients)

	fmt.Println(result)
}
