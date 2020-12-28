package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
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

func findInertIngredients(in string) ([]string, []string) {
	lines := make([]line, 0)

	allIngredients := make(map[string]bool)
	allAllergens := make(map[string]bool)

	for _, lineStr := range strings.Split(in, "\n") {
		lines = append(lines, parseLine(lineStr))
	}

	for _, line := range lines {
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

	for _, line := range lines {
		for ingredient := range allIngredients {
			if !contains(line.ingredients, ingredient) {
				for _, allergen := range line.allergens {
					delete(canContain[ingredient], allergen)
				}
			}
		}
	}

	inertIgredients := make([]string, 0)
	for ingredient, allergens := range canContain {
		if len(allergens) == 0 {
			inertIgredients = append(inertIgredients, ingredient)
		}
	}

	allergenOf := make(map[string]string)
	change := true
	for change {
		change = false
		for ingredient, allergens := range canContain {
			if len(allergens) == 1 {
				allergen := singleKeyFrom(allergens)
				allergenOf[ingredient] = allergen
				for otherIngredient := range allIngredients {
					delete(canContain[otherIngredient], allergen)
				}
				change = true
			}
		}
	}

	canonicalList := make([]string, 0)
	for ingredient := range allergenOf {
		canonicalList = append(canonicalList, ingredient)
	}
	sort.Slice(canonicalList, func(i, j int) bool {
		return allergenOf[canonicalList[i]] < allergenOf[canonicalList[j]]
	})

	return inertIgredients, canonicalList
}

func singleKeyFrom(m map[string]bool) string {
	for k := range m {
		return k
	}
	panic("no elem in map")
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

	ingredients, list := findInertIngredients(string(input))
	result := countAll(string(input), ingredients)

	fmt.Println(result)
	fmt.Println(strings.Join(list, ","))
}
