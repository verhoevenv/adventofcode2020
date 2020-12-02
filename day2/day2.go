package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Policy struct {
	lowest  int
	highest int
	letter  string
}

type Password string

func ValidOld(pass Password, policy Policy) bool {
	nb := strings.Count(string(pass), policy.letter)
	return policy.lowest <= nb && nb <= policy.highest
}

func ValidNew(pass Password, policy Policy) bool {
	return (string(pass)[policy.lowest-1] == policy.letter[0]) !=
		(string(pass)[policy.highest-1] == policy.letter[0])
}

var listFormat = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

func Parse(line string) (Policy, Password) {
	matches := listFormat.FindStringSubmatch(line)
	policy := Policy{
		lowest:  unsafeAtoi(matches[1]),
		highest: unsafeAtoi(matches[2]),
		letter:  matches[3],
	}
	password := Password(matches[4])
	return policy, password
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
	numberCorrect := 0

	for scanner.Scan() {
		pol, pass := Parse(scanner.Text())
		if ValidNew(pass, pol) {
			numberCorrect++
		}
	}

	fmt.Println(numberCorrect)
}
