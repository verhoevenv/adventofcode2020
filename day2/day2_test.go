package main

import "testing"

func TestParse(t *testing.T) {
	pol, pass := Parse("1-3 a: abcde")

	expectedPol := Policy{1, 3, "a"}
	expectedPass := Password("abcde")
	if pol != expectedPol && pass != expectedPass {
		t.Errorf("Parse incorrect, got: %v and %v, want: %v and %v.",
			pol, pass, expectedPol, expectedPass)
	}
}

func TestValid(t *testing.T) {
	tables := []struct {
		pass  Password
		pol   Policy
		valid bool
	}{
		{Password("abcde"), Policy{1, 3, "a"}, true},
		{Password("cdefg"), Policy{1, 3, "b"}, false},
		{Password("ccccccccc"), Policy{2, 9, "c"}, true},
	}

	for _, table := range tables {
		result := Valid(table.pass, table.pol)
		if result != table.valid {
			t.Errorf("Validation of (%v,%v) was incorrect, got: %v, want: %v.",
				table.pol, table.pass, result, table.valid)
		}
	}
}
