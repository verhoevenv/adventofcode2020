package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	id    uint64
	delay uint64
}

func findFindSubsequentDepartTimes(in string) uint64 {
	busStrs := strings.Split(in, ",")
	busses := make([]bus, 0)

	for delay, busID := range busStrs {
		if busID != "x" {
			busses = append(busses, bus{unsafeAtoi(busID), uint64(delay)})
		}
	}

	timestamp := busses[0].id
	jump := busses[0].id

	for _, b := range busses[1:] {
		for timestamp%b.id != (b.id - (b.delay % b.id)) {
			timestamp += jump
		}
		jump *= b.id
	}

	return timestamp
}

func unsafeAtoi(in string) uint64 {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return uint64(v)
}

func part2Main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	lines := strings.Split(string(input), "\n")

	result := findFindSubsequentDepartTimes(lines[1])

	fmt.Println(result)
}
