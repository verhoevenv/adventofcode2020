package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type group []string

func (g *group) anyYesAnswers() int {
	perQuestion := make(map[rune]bool)

	for _, person := range *g {
		for _, qWithAnswerYes := range person {
			perQuestion[qWithAnswerYes] = true
		}
	}

	return len(perQuestion)
}

func (g *group) allYesAnswers() int {
	perQuestion := make(map[rune]bool)

	const allQuestions = "abcdefghijklmnopqrstuvwxyz"
	for _, q := range allQuestions {
		perQuestion[q] = true
	}

	for _, person := range *g {
		seenForPerson := make(map[rune]bool)
		for _, qWithAnswerYes := range person {
			seenForPerson[qWithAnswerYes] = true
		}
		for _, q := range allQuestions {
			if !seenForPerson[q] {
				delete(perQuestion, q)
			}
		}
	}

	return len(perQuestion)
}

func parse(in string) []group {
	result := make([]group, 0)

	for _, groupStr := range strings.Split(in, "\n\n") {
		groupStrs := make([]string, 0)

		for _, line := range strings.Split(groupStr, "\n") {
			groupStrs = append(groupStrs, line)
		}

		result = append(result, group(groupStrs))
	}

	return result
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	gs := parse(string(input))

	sum := 0
	for _, g := range gs {
		sum += g.allYesAnswers()
	}

	fmt.Println(sum)
}
