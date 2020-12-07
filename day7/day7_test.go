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
