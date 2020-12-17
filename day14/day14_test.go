package main

import (
	"testing"
)

var programStr = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestApplyMask(t *testing.T) {
	mask := makeMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	tables := []struct {
		val       uint64
		maskedVal uint64
	}{
		{11, 73},
		{101, 101},
		{0, 64},
	}

	for _, table := range tables {
		result := mask.apply(table.val)
		if result != table.maskedVal {
			t.Errorf("Masking of %v was incorrect, got: %v, want: %v.",
				table.val, result, table.maskedVal)
		}
	}
}

func TestRunProgram(t *testing.T) {
	s := makeSys()

	s.runProgram(programStr)

	result := s.sumOfMemory()

	if result != 165 {
		t.Errorf("RunProgram was incorrect, got: %v, want: %v.",
			result, 165)
	}
}
