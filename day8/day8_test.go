package main

import (
	"testing"
)

var instructions = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestAccumulateUntilLoopDetected(t *testing.T) {
	program := parse(instructions)

	result := program.accumulateUntilLoopDetected()

	if result != 5 {
		t.Errorf("accumulateUntilLoopDetected incorrect, got: %v, want %v.",
			result, 5)
	}

}
