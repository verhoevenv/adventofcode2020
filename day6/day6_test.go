package main

import "testing"

func TestAnyYesAnswers(t *testing.T) {

	tables := []struct {
		g     group
		count int
	}{
		{group([]string{"abc"}), 3},
		{group([]string{"a", "b", "c"}), 3},
		{group([]string{"ab", "ac"}), 3},
		{group([]string{"a", "a", "a", "a"}), 1},
		{group([]string{"b"}), 1},
	}

	for _, table := range tables {
		result := table.g.allYesAnswers()
		if result != table.count {
			t.Errorf("Count of 'any' yes answers of %v was incorrect, got: %v, want: %v.",
				table.g, result, table.count)
		}
	}

}

func TestAllYesAnswers(t *testing.T) {

	tables := []struct {
		g     group
		count int
	}{
		{group([]string{"abc"}), 3},
		{group([]string{"a", "b", "c"}), 0},
		{group([]string{"ab", "ac"}), 1},
		{group([]string{"a", "a", "a", "a"}), 1},
		{group([]string{"b"}), 1},
	}

	for _, table := range tables {
		result := table.g.allYesAnswers()
		if result != table.count {
			t.Errorf("Count of 'all' yes answers of %v was incorrect, got: %v, want: %v.",
				table.g, result, table.count)
		}
	}

}
