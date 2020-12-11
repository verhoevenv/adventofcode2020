package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type adapters struct {
	sortedData []int
}

func makeAdapters(data []int) *adapters {
	result := make([]int, len(data)+2)

	result[0] = 0
	result[1] = max(data) + 3
	for i, v := range data {
		result[i+2] = v
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return &adapters{result}
}

type idxRange struct {
	from int
	to   int
}

func max(data []int) int {
	m := data[0]
	for _, e := range data {
		if e > m {
			m = e
		}
	}
	return m
}

func (adaps *adapters) chainTo() []int {
	return adaps.sortedData
}

func (adaps *adapters) countDifferencesInChain() map[int]int {
	chain := adaps.chainTo()
	deltas := make(map[int]int)

	for i := 1; i < len(chain); i++ {
		delta := chain[i] - chain[i-1]
		deltas[delta]++
	}

	return deltas
}

func (adaps *adapters) countCombinations() uint64 {
	return adaps._countCombinations(
		idxRange{0, len(adaps.sortedData) - 1},
		make(map[idxRange]uint64),
	)
}

func (adaps *adapters) _countCombinations(r idxRange, mem map[idxRange]uint64) uint64 {
	if v, ok := mem[r]; ok {
		return v
	}

	if r.to-r.from < 3 {
		return 1
	}

	a := adaps.sortedData[r.from]
	b := adaps.sortedData[r.from+1]
	c := adaps.sortedData[r.from+2]
	d := adaps.sortedData[r.from+3]

	var num uint64 = 0

	switch {
	case b-a == 3:
		num = adaps._countCombinations(idxRange{r.from + 1, r.to}, mem)
	case b-a == 2 && c-a > 3:
		num = adaps._countCombinations(idxRange{r.from + 1, r.to}, mem)
	case b-a == 2 && c-a == 3:
		num = adaps._countCombinations(idxRange{r.from + 1, r.to}, mem) +
			adaps._countCombinations(idxRange{r.from + 2, r.to}, mem)
	case b-a == 1 && c-a > 3:
		num = adaps._countCombinations(idxRange{r.from + 1, r.to}, mem)
	case b-a == 1 && c-a == 3:
		num = adaps._countCombinations(idxRange{r.from + 1, r.to}, mem) +
			adaps._countCombinations(idxRange{r.from + 2, r.to}, mem)
	case b-a == 1 && c-a == 2 && d-a > 3:
		num = adaps._countCombinations(idxRange{r.from + 1, r.to}, mem) +
			adaps._countCombinations(idxRange{r.from + 2, r.to}, mem)
	case b-a == 1 && c-a == 2 && d-a == 3:
		num = adaps._countCombinations(idxRange{r.from + 1, r.to}, mem) +
			adaps._countCombinations(idxRange{r.from + 2, r.to}, mem) +
			adaps._countCombinations(idxRange{r.from + 3, r.to}, mem)
	default:
		panic(fmt.Sprintf("not sure how to handle (%v, %v, %v, %v)", a, b, c, d))
	}

	mem[r] = num
	return num
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
	adapters := makeAdapters(numbers)

	combs := adapters.countCombinations()

	fmt.Println(combs)
}
