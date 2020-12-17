package main

import (
	"fmt"
)

func startMemory(start []int) func() int {
	alreadySpoken := make(map[int]int)
	lastSpoken := -1
	turn := 0

	return func() int {
		turn++

		toSpeak := -1
		if len(alreadySpoken) < len(start) {
			toSpeak = start[len(alreadySpoken)]
		} else {
			prevTurn, notFirstTime := alreadySpoken[lastSpoken]
			if notFirstTime {
				toSpeak = turn - prevTurn
			} else {
				toSpeak = 0
			}
		}

		alreadySpoken[lastSpoken] = turn
		lastSpoken = toSpeak

		return toSpeak
	}
}

func nTimes(n int, f func() int) int {
	for i := 0; i < n-1; i++ {
		f()
	}

	return f()
}

func main() {
	input := []int{2, 15, 0, 9, 1, 20}

	mem := startMemory(input)
	result := nTimes(30000000, mem)

	fmt.Println(result)
}
