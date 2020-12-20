package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(inWithSpace string) int {
	in := strings.ReplaceAll(inWithSpace, " ", "")

	if len(in) == 1 {
		return unsafeAtoi(in)
	}

	if in[0] == '(' && matching(in, 0) == len(in)-1 {
		return calculate(in[1 : len(in)-1])
	}

	l, o, r := split(in)

	switch o {
	case '+':
		return l + r
	case '*':
		return l * r
	default:
		panic("can't handle " + in)
	}
}

func split(in string) (int, byte, int) {
	opIdx := len(in) - 2
	if in[len(in)-1] == ')' {
		opIdx = matching(in, len(in)-1) - 1
	}

	return calculate(in[0:opIdx]), in[opIdx], calculate(in[opIdx+1:])
}

func matching(in string, bracketIdx int) int {
	var dir int
	switch in[bracketIdx] {
	case '(':
		dir = 1
	case ')':
		dir = -1
	}
	level := 1
	idx := bracketIdx
	for level > 0 {
		idx += dir
		switch in[idx] {
		case '(':
			level += dir
		case ')':
			level -= dir
		}
	}
	return idx
}

func unsafeAtoi(in string) int {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return v
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		sum += calculate(scanner.Text())
	}

	fmt.Println(sum)
}
