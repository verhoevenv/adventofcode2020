package main

import (
	"reflect"
	"testing"
)

var rules = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestParse(t *testing.T) {
	tables := []struct {
		bagColor         color
		expectedContents []contents
	}{
		{"faded blue", nil},
		{"bright white", []contents{{1, "shiny gold"}}},
		{"dark olive", []contents{{3, "faded blue"}, {4, "dotted black"}}},
	}

	regulations := parse(rules)

	for _, table := range tables {
		result := regulations[table.bagColor]
		if !reflect.DeepEqual(result, table.expectedContents) {
			t.Errorf("Parse of %v was incorrect, got: %v, want: %v.",
				table.bagColor, result, table.expectedContents)
		}
	}

}

func TestHowManyCanHoldBag(t *testing.T) {
	regulations := parse(rules)

	result := regulations.howManyCanHoldShinyGoldBag()

	if result != 4 {
		t.Errorf("howManyCanHoldBag incorrect, got: %v, want %v.",
			result, 4)
	}

}

func TestHowManyInShinyGoldBag(t *testing.T) {
	regulations := parse(rules)

	result := regulations.howManyInShinyGoldBag()

	if result != 32 {
		t.Errorf("howManyInShinyGoldBag incorrect, got: %v, want %v.",
			result, 32)
	}

}

var rules2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

func TestHowManyInShinyGoldBag2(t *testing.T) {
	regulations := parse(rules2)

	result := regulations.howManyInShinyGoldBag()

	if result != 126 {
		t.Errorf("howManyInShinyGoldBag incorrect, got: %v, want %v.",
			result, 126)
	}

}
