package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type adapters struct {
	rawData []int
}

func (adaps *adapters) max() int {
	m := adaps.rawData[0]
	for _, e := range adaps.rawData {
		if e > m {
			m = e
		}
	}
	return m
}

func (adaps *adapters) chainTo(dest int) []int {
	result := make([]int, len(adaps.rawData)+2)

	result[0] = 0
	result[1] = dest
	for i, v := range adaps.rawData {
		result[i+2] = v
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func (adaps *adapters) countDifferencesInChain(dest int) map[int]int {
	chain := adaps.chainTo(dest)
	deltas := make(map[int]int)

	for i := 1; i < len(chain); i++ {
		delta := chain[i] - chain[i-1]
		deltas[delta]++
	}

	return deltas
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}
	adapters := adapters{numbers}

	diffs := adapters.countDifferencesInChain(adapters.max() + 3)

	fmt.Println(diffs[1] * diffs[3])
}
