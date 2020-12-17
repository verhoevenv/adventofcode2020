package main

import (
	"reflect"
	"sort"
	"testing"
)

var programStr = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestApplyMask(t *testing.T) {
	mask := makeMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	tables := []struct {
		val       memVal
		maskedVal memVal
	}{
		{11, 73},
		{101, 101},
		{0, 64},
	}

	for _, table := range tables {
		result := mask.applyToVal(table.val)
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

func TestApplyMaskToAddress(t *testing.T) {

	tables := []struct {
		addr   memAddr
		m      *mask
		result []memAddr
	}{
		{42, makeMask("000000000000000000000000000000X1001X"), []memAddr{
			26, 27, 58, 59,
		}},
		{26, makeMask("00000000000000000000000000000000X0XX"), []memAddr{
			16, 17, 18, 19, 24, 25, 26, 27,
		}},
	}

	for _, table := range tables {
		result := table.m.applyToAddr(table.addr)
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Masking of %v was incorrect, got: %v, want: %v.",
				table.addr, result, table.result)
		}
	}
}

var program2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestRunProgramV2(t *testing.T) {
	s := makeSys()

	s.runProgramV2(program2)

	result := s.sumOfMemory()

	if result != 208 {
		t.Errorf("RunProgramV2 was incorrect, got: %v, want: %v.",
			result, 208)
	}
}
