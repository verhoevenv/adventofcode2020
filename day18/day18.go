package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type precedence func(in string) (string, byte, string)

func calculate(inWithSpace string, rules precedence) int {
	in := strings.ReplaceAll(inWithSpace, " ", "")

	if len(in) == 1 {
		return unsafeAtoi(in)
	}

	if in[0] == '(' && matching(in, 0) == len(in)-1 {
		return calculate(in[1:len(in)-1], rules)
	}

	l, o, r := rules(in)

	switch o {
	case '+':
		return calculate(l, rules) + calculate(r, rules)
	case '*':
		return calculate(l, rules) * calculate(r, rules)
	default:
		panic("can't handle " + in)
	}
}

func allSame(in string) (string, byte, string) {
	opIdx := len(in) - 2
	if in[len(in)-1] == ')' {
		opIdx = matching(in, len(in)-1) - 1
	}

	return in[0:opIdx], in[opIdx], in[opIdx+1:]
}

func plusFirst(in string) (string, byte, string) {
	for idx := len(in) - 1; idx > 0; idx-- {
		if in[idx] == ')' {
			idx = matching(in, idx)
		}
		if in[idx] == '*' {
			return in[0:idx], in[idx], in[idx+1:]
		}
	}

	return allSame(in)
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
		sum += calculate(scanner.Text(), plusFirst)
	}

	fmt.Println(sum)
}
