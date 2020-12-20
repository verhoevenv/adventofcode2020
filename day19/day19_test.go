package main

import (
	"testing"
)

var rulesStr = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"`

func TestMatches(t *testing.T) {
	rules := parseRules(rulesStr)

	tables := []struct {
		message string
		matches bool
	}{
		{"ababbb", true},
		{"bababa", false},
		{"abbbab", true},
		{"aaabbb", false},
		{"aaaabbb", false},
	}

	for _, table := range tables {
		result := matches(rules[0], table.message)
		if result != table.matches {
			t.Errorf("Match of %v was incorrect, got: %v, want: %v.",
				table.message, result, table.matches)
		}
	}
}

func TestOr(t *testing.T) {
	rule := or(lit('a'), lit('b'))

	tables := []struct {
		message string
		matches bool
	}{
		{"a", true},
		{"b", true},
		{"", false},
		{"ab", false},
	}

	for _, table := range tables {
		result := matches(rule, table.message)
		if result != table.matches {
			t.Errorf("Match of %v was incorrect, got: %v, want: %v.",
				table.message, result, table.matches)
		}
	}
}

func TestFollowing(t *testing.T) {
	rule := following([]int{1, 2}, rules{1: lit('a'), 2: lit('b')})

	tables := []struct {
		message string
		matches bool
	}{
		{"a", false},
		{"b", false},
		{"", false},
		{"ab", true},
		{"aba", false},
	}

	for _, table := range tables {
		result := matches(rule, table.message)
		if result != table.matches {
			t.Errorf("Match of %v was incorrect, got: %v, want: %v.",
				table.message, result, table.matches)
		}
	}
}
