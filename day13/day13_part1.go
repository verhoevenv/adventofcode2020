package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type schedule struct {
	departTime int
	busIDs     []int
}

func findMinWaitTime(s *schedule) (int, int) {
	minWaitTime := math.MaxInt32
	minBusID := -1
	for _, busID := range s.busIDs {
		waitTime := busID - (s.departTime % busID)
		if waitTime < minWaitTime {
			minWaitTime = waitTime
			minBusID = busID
		}
	}
	return minWaitTime, minBusID
}

func parsePart1(iStr string) schedule {
	lines := strings.Split(iStr, "\n")

	departTime, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	busIDs := make([]int, 0)
	for _, busIDStr := range strings.Split(lines[1], ",") {
		busID, err := strconv.Atoi(busIDStr)
		if err != nil {
			//ignore, it's the mysterious "x"
			continue
		}
		busIDs = append(busIDs, busID)
	}

	return schedule{departTime, busIDs}
}

func part1Main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	schedule := parsePart1(string(input))

	time, busID := findMinWaitTime(&schedule)

	fmt.Println(time * busID)
}
