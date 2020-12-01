package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PairThatSumsTo2020(report []int) (a, b int) {
	for i1, v1 := range report {
		for _, v2 := range report[i1:] {
			if v1+v2 == 2020 {
				return v1, v2
			}
		}
	}
	panic("No summing pair found")
}

func TripletThatSumsTo2020(report []int) (a, b, c int) {
	for i1, v1 := range report {
		for i2, v2 := range report[i1:] {
			for _, v3 := range report[i2:] {
				if v1+v2+v3 == 2020 {
					return v1, v2, v3
				}
			}
		}
	}
	panic("No summing triplet found")
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

	first, second, third := TripletThatSumsTo2020(numbers)
	result := first * second * third
	fmt.Println(result)
}
