package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type mask struct {
	orig       string
	clearMask  uint64
	setMask    uint64
	floatyBits []uint64
}

func makeMask(in string) *mask {
	clearMask := uint64(0)
	setMask := uint64(0)
	floatyBits := make([]uint64, 0)

	for i, b := range in {
		clearMask = clearMask << 1
		setMask = setMask << 1

		switch b {
		case '1':
			setMask++
		case '0':
			clearMask++
		case 'X':
			floatyBits = append(floatyBits, 1<<(35-i))
		}
	}

	return &mask{
		in,
		clearMask,
		setMask,
		floatyBits,
	}
}

func (m *mask) applyToVal(val memVal) memVal {
	asInt := uint64(val)
	return memVal((asInt | m.setMask) &^ m.clearMask)
}

func (m *mask) applyToAddr(original memAddr) []memAddr {
	result := make([]memAddr, 1)
	masked := memAddr(uint64(original) | m.setMask)

	result[0] = masked

	for _, floatMask := range m.floatyBits {
		expandedResult := make([]memAddr, 0, len(result)*2)

		for _, addr := range result {
			asInt := uint64(addr)

			v1 := memAddr((asInt | floatMask))
			v2 := memAddr((asInt &^ floatMask))
			expandedResult = append(expandedResult, v1, v2)
		}

		result = expandedResult
	}

	return result
}

type sys struct {
	currMask *mask
	mem      map[memAddr]memVal
}

type memAddr uint64
type memVal uint64

func makeSys() *sys {
	return &sys{
		nil,
		make(map[memAddr]memVal),
	}
}

var setMaskLine = regexp.MustCompile(`mask = ([X10]+)`)
var setMemLine = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func (s *sys) run(line string) {
	maskMatch := setMaskLine.FindStringSubmatch(line)
	if maskMatch != nil {
		s.currMask = makeMask(maskMatch[1])
	}
	memMatch := setMemLine.FindStringSubmatch(line)
	if memMatch != nil {
		a := memAddr(unsafeAtoi(memMatch[1]))
		v := memVal(unsafeAtoi(memMatch[2]))

		s.mem[a] = s.currMask.applyToVal(v)
	}
}

func (s *sys) runV2(line string) {
	maskMatch := setMaskLine.FindStringSubmatch(line)
	if maskMatch != nil {
		s.currMask = makeMask(maskMatch[1])
	}
	memMatch := setMemLine.FindStringSubmatch(line)
	if memMatch != nil {
		a := memAddr(unsafeAtoi(memMatch[1]))
		v := memVal(unsafeAtoi(memMatch[2]))

		for _, decodedA := range s.currMask.applyToAddr(a) {
			s.mem[decodedA] = v
		}
	}
}

func unsafeAtoi(in string) uint64 {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return uint64(v)
}

func (s *sys) runProgram(program string) {
	for _, line := range strings.Split(program, "\n") {
		s.run(line)
	}
}

func (s *sys) runProgramV2(program string) {
	for _, line := range strings.Split(program, "\n") {
		s.runV2(line)
	}
}

func (s *sys) sumOfMemory() uint64 {
	counter := uint64(0)
	for _, v := range s.mem {
		counter += uint64(v)
	}
	return counter
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	sys := makeSys()
	sys.runProgramV2(string(input))

	fmt.Println(sys.sumOfMemory())
}
