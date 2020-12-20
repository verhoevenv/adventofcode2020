package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type rule func(string) []split
type rules map[int]rule

func parse(in string) (rules, []string) {
	parts := strings.Split(in, "\n\n")

	rules := parseRules(parts[0])

	messages := strings.Split(parts[1], "\n")

	return rules, messages
}

func parseRules(in string) rules {
	result := make(map[int]rule)

	for _, ruleStr := range strings.Split(in, "\n") {
		ruleSplit := strings.Split(ruleStr, ": ")
		ruleNum := unsafeAtoi(ruleSplit[0])
		rule := parseRule(ruleSplit[1], result)

		result[ruleNum] = rule
	}

	return result
}

func parseRule(in string, allRules rules) rule {
	if strings.Contains(in, " | ") {
		parts := strings.Split(in, " | ")
		return or(parseRule(parts[0], allRules), parseRule(parts[1], allRules))
	}
	if strings.HasPrefix(in, "\"") {
		return lit(in[1])
	}
	nextRulesStr := strings.Split(in, " ")
	nextRules := make([]int, 0)
	for _, r := range nextRulesStr {
		nextRules = append(nextRules, unsafeAtoi(r))
	}
	return following(nextRules, allRules)
}

type split struct {
	pre  string
	post string
}

func matches(r rule, in string) bool {
	allSplits := r(in)
	for _, split := range allSplits {
		if split.post == "" {
			return true
		}
	}
	return false
}

func or(l rule, r rule) rule {
	return func(in string) []split {
		result := make([]split, 0)
		for _, match := range l(in) {
			result = append(result, match)
		}
		for _, match := range r(in) {
			result = append(result, match)
		}
		return result
	}
}

func following(otherRules []int, allRules rules) rule {
	return func(in string) []split {
		toMatch := []split{split{"", in}}
		for _, rIdx := range otherRules {
			rule := allRules[rIdx]

			nextToMatch := make([]split, 0)

			for _, next := range toMatch {
				newSplits := rule(next.post)
				for _, newSplit := range newSplits {
					nextToMatch = append(nextToMatch, split{next.pre + newSplit.pre, newSplit.post})
				}
			}

			toMatch = nextToMatch
		}

		return toMatch
	}
}

func lit(b byte) rule {
	return func(in string) []split {
		var result []split
		if len(in) == 0 {
			result = nil
		} else if in[0] == b {
			result = []split{split{in[0:1], in[1:]}}
		} else {
			result = nil
		}

		return result
	}
}

func unsafeAtoi(in string) int {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return v
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	rules, messages := parse(string(input))

	count := 0
	for _, m := range messages {
		if matches(rules[0], m) {
			count++
		}
	}
	fmt.Println(count)
}
