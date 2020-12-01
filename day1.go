package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ElemsThatSumTo2020(report []int) (a, b int) {
	for i1, v1 := range report {
		for _, v2 := range report[i1:] {
			if v1+v2 == 2020 {
				return v1, v2
			}
		}
	}
	panic("No summing pair found")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 5)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}

	first, second := ElemsThatSumTo2020(numbers)
	result := first * second
	fmt.Println(result)
}
