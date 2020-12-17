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
	orig      string
	clearMask uint64
	setMask   uint64
}

func makeMask(in string) *mask {
	clearMask := uint64(0)
	setMask := uint64(0)

	for _, b := range in {
		clearMask = clearMask << 1
		setMask = setMask << 1

		switch b {
		case '1':
			setMask++
		case '0':
			clearMask++
		}
	}

	return &mask{
		in,
		clearMask,
		setMask,
	}
}

func (m *mask) apply(val uint64) uint64 {
	return (val | m.setMask) &^ m.clearMask
}

type sys struct {
	currMask *mask
	mem      map[uint64]uint64
}

func makeSys() *sys {
	return &sys{
		nil,
		make(map[uint64]uint64),
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
		a := unsafeAtoi(memMatch[1])
		v := unsafeAtoi(memMatch[2])

		s.mem[a] = s.currMask.apply(v)
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

func (s *sys) sumOfMemory() uint64 {
	counter := uint64(0)
	for _, v := range s.mem {
		counter += v
	}
	return counter
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	sys := makeSys()
	sys.runProgram(string(input))

	fmt.Println(sys.sumOfMemory())
}
