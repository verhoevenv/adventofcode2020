package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tables := []struct {
		expr   string
		result int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, table := range tables {
		result := calculate(table.expr)
		if result != table.result {
			t.Errorf("Calculate of %v was incorrect, got: %v, want: %v.",
				table.expr, result, table.result)
		}
	}
}
