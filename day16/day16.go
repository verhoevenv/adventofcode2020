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
		result = append(result, invalidFields(rules, ticket)...)
	}
	return result
}

func invalidFields(rules []fieldRule, t ticket) []int {
	result := make([]int, 0)
	for _, value := range t {
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
	return result
}

func filterOutInvalidTickets(rules []fieldRule, tickets []ticket) []ticket {
	result := make([]ticket, 0)
	for _, ticket := range tickets {
		if len(invalidFields(rules, ticket)) == 0 {
			result = append(result, ticket)
		}
	}
	return result
}

func determineOrdering(rules []fieldRule, validTickets []ticket) []string {
	ruleOutFieldXHasRuleY := make([][]bool, len(rules))
	for i := 0; i < len(rules); i++ {
		ruleOutFieldXHasRuleY[i] = make([]bool, len(rules))
	}

	for _, ticket := range validTickets {
		for field := 0; field < len(rules); field++ {
			for rule := 0; rule < len(rules); rule++ {
				if !rules[rule].canHave(ticket[field]) {
					ruleOutFieldXHasRuleY[field][rule] = true
				}
			}
		}
	}

	knownFields := make(map[int]string)
	for len(knownFields) < len(rules) {
		for field := 0; field < len(rules); field++ {
			possibilities := make([]int, 0)
			for rule := 0; rule < len(rules); rule++ {
				if !ruleOutFieldXHasRuleY[field][rule] {
					possibilities = append(possibilities, rule)
				}
			}
			if len(possibilities) == 1 {
				knownFields[field] = rules[possibilities[0]].field
				for otherField := 0; otherField < len(rules); otherField++ {
					if otherField != field {
						ruleOutFieldXHasRuleY[otherField][possibilities[0]] = true
					}
				}
			}
		}
	}

	result := make([]string, len(rules))
	for i := 0; i < len(rules); i++ {
		result[i] = knownFields[i]
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
	valids := filterOutInvalidTickets(notes.rules, notes.nearby)
	ordering := determineOrdering(notes.rules, valids)

	mult := 1
	for i, v := range ordering {
		if strings.HasPrefix(v, "departure") {
			mult *= notes.your[i]
		}
	}

	fmt.Println(mult)
}
