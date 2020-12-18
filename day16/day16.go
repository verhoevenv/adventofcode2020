package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type notes struct {
	rules  []fieldRule
	your   ticket
	nearby []ticket
}

type fieldRule struct {
	field   string
	canHave func(int) bool
}

type ticket []int

func inRange(l1 int, h1 int, l2 int, h2 int) func(int) bool {
	return func(v int) bool {
		return (v >= l1 && v <= h1) || (v >= l2 && v <= h2)
	}
}

func findCompletelyInvalidFields(rules []fieldRule, tickets []ticket) []int {
	result := make([]int, 0)
	for _, ticket := range tickets {
		for _, value := range ticket {
			noneValid := true
			for _, rule := range rules {
				if rule.canHave(value) {
					noneValid = false
					break
				}
			}
			if noneValid {
				result = append(result, value)
			}
		}
	}
	return result
}

var ruleRE = regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)

func parse(in string) *notes {
	var result notes

	parts := strings.Split(in, "\n\n")

	result.rules = make([]fieldRule, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		matches := ruleRE.FindStringSubmatch(line)

		result.rules = append(result.rules, fieldRule{
			field: matches[1],
			canHave: inRange(
				unsafeAtoi(matches[2]),
				unsafeAtoi(matches[3]),
				unsafeAtoi(matches[4]),
				unsafeAtoi(matches[5]),
			),
		})
	}

	yourStr := strings.Split(parts[1], "\n")[1]
	result.your = parseTicket(yourStr)

	result.nearby = make([]ticket, 0)
	for _, line := range strings.Split(parts[2], "\n")[1:] {
		result.nearby = append(result.nearby, parseTicket(line))
	}

	return &result
}

func parseTicket(line string) ticket {
	result := make(ticket, 0)
	for _, v := range strings.Split(line, ",") {
		result = append(result, unsafeAtoi(v))
	}
	return result
}

func unsafeAtoi(in string) int {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return v
}

func sum(ints []int) int {
	counter := 0
	for _, v := range ints {
		counter += v
	}
	return counter
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	notes := parse(string(input))

	result := findCompletelyInvalidFields(notes.rules, notes.nearby)

	fmt.Println(sum(result))
}
