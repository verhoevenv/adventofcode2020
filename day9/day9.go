package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type xmasStream struct {
	data              []uint64
	preambleBeginning int
	preambleLength    int
	next              int
}

func makeXMAS(data []uint64, preambleLength int) xmasStream {
	return xmasStream{data, 0, preambleLength, preambleLength}
}

func (xmas *xmasStream) existsAsSum(next uint64) bool {
	//naive solution seems fast enough
	for i := xmas.preambleBeginning; i < xmas.next; i++ {
		for j := i + 1; j < xmas.next; j++ {
			if xmas.data[i]+xmas.data[j] == next {
				return true
			}
		}
	}
	return false
}

func (xmas *xmasStream) shift() {
	xmas.preambleBeginning++
	xmas.next++
}

func (xmas *xmasStream) findNonSumNumber() uint64 {
	for { // <-- will naturally oob-crash if not found
		if !xmas.existsAsSum(xmas.data[xmas.next]) {
			return xmas.data[xmas.next]
		}
		xmas.shift()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]uint64, 0)
	for scanner.Scan() {
		i, err := strconv.ParseUint(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}
	xmas := makeXMAS(numbers, 25)

	fmt.Println(xmas.findNonSumNumber())
}
