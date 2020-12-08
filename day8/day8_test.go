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

	accResult, normalTermination := program.accumulateWithLoopDetection()

	if normalTermination {
		t.Errorf("accumulateWithLoopDetection incorrect, expected loop detected.")
	}

	if accResult != 5 {
		t.Errorf("accumulateWithLoopDetection incorrect, got: %v, want %v.",
			accResult, 5)
	}

}

var terminatingProgram = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
nop -4
acc +6`

func TestAccumulateUntilLoopDetectedWithoutLoop(t *testing.T) {
	program := parse(terminatingProgram)

	accResult, normalTermination := program.accumulateWithLoopDetection()

	if !normalTermination {
		t.Errorf("accumulateWithLoopDetection incorrect, expected normal termination.")
	}

	if accResult != 8 {
		t.Errorf("accumulateWithLoopDetection incorrect, got: %v, want %v.",
			accResult, 8)
	}

}

func TestFixProgram(t *testing.T) {
	program := parse(instructions)

	result := program.fixProgram()

	if result != 8 {
		t.Errorf("fixProgram incorrect, got: %v, want %v.",
			result, 8)
	}

}
