package main

import (
	"fmt"
)

func startMemory(start []int) func() int {
	alreadySpoken := make([]int, 0)

	return func() int {
		if len(alreadySpoken) < len(start) {
			toSpeak := start[len(alreadySpoken)]
			alreadySpoken = append(alreadySpoken, toSpeak)
			return toSpeak
		}

		lastSpoken := alreadySpoken[len(alreadySpoken)-1]

		spokenBefore := -1
		toSpeak := 0

		for i := len(alreadySpoken) - 1; i >= 0; i-- {
			if alreadySpoken[i] == lastSpoken {
				if spokenBefore == -1 {
					spokenBefore = i
				} else {
					toSpeak = spokenBefore - i
					break
				}
			}
		}

		alreadySpoken = append(alreadySpoken, toSpeak)
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
	result := nTimes(2020, mem)

	fmt.Println(result)
}
