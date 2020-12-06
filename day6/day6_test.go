package main

import "testing"

func TestYesAnswers(t *testing.T) {

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
		result := table.g.yesAnswers()
		if result != table.count {
			t.Errorf("Count of yes answers of %v was incorrect, got: %v, want: %v.",
				table.g, result, table.count)
		}
	}

}
