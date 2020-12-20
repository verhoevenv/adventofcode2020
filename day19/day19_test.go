package main

import (
	"strings"
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

var rulesPart2Str = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1`

var messagesPart2 = strings.Split(``, "\n")

func TestMatchesPart2Base(t *testing.T) {
	rules := parseRules(rulesPart2Str)

	tables := []struct {
		message string
		matches bool
	}{
		{"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa", false},
		{"bbabbbbaabaabba", true},
		{"babbbbaabbbbbabbbbbbaabaaabaaa", false},
		{"aaabbbbbbaaaabaababaabababbabaaabbababababaaa", false},
		{"bbbbbbbaaaabbbbaaabbabaaa", false},
		{"bbbababbbbaaaaaaaabbababaaababaabab", false},
		{"ababaaaaaabaaab", true},
		{"ababaaaaabbbaba", true},
		{"baabbaaaabbaaaababbaababb", false},
		{"abbbbabbbbaaaababbbbbbaaaababb", false},
		{"aaaaabbaabaaaaababaa", false},
		{"aaaabbaaaabbaaa", false},
		{"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa", false},
		{"babaaabbbaaabaababbaabababaaab", false},
		{"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba", false},
	}

	for _, table := range tables {
		result := matches(rules[0], table.message)
		if result != table.matches {
			t.Errorf("Match of %v was incorrect, got: %v, want: %v.",
				table.message, result, table.matches)
		}
	}
}

func TestMatchesPart2(t *testing.T) {
	rules := parseRules(rulesPart2Str)

	update8And11(rules)

	tables := []struct {
		message string
		matches bool
	}{
		{"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa", false},
		{"bbabbbbaabaabba", true},
		{"babbbbaabbbbbabbbbbbaabaaabaaa", true},
		{"aaabbbbbbaaaabaababaabababbabaaabbababababaaa", true},
		{"bbbbbbbaaaabbbbaaabbabaaa", true},
		{"bbbababbbbaaaaaaaabbababaaababaabab", true},
		{"ababaaaaaabaaab", true},
		{"ababaaaaabbbaba", true},
		{"baabbaaaabbaaaababbaababb", true},
		{"abbbbabbbbaaaababbbbbbaaaababb", true},
		{"aaaaabbaabaaaaababaa", true},
		{"aaaabbaaaabbaaa", false},
		{"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa", true},
		{"babaaabbbaaabaababbaabababaaab", false},
		{"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba", true},
	}

	for _, table := range tables {
		result := matches(rules[0], table.message)
		if result != table.matches {
			t.Errorf("Match of %v was incorrect, got: %v, want: %v.",
				table.message, result, table.matches)
		}
	}
}
