package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type color string
type regulations map[color][]contents
type contents struct {
	num int
	bag color
}

func (reg *regulations) howManyCanHoldShinyGoldBag() int {
	canHold := make(map[color]bool)

	for len(canHold) < len(*reg) {
		for bag, inBag := range *reg {
			if _, known := canHold[bag]; known {
				continue
			}
			unknownSubBags := len(inBag)
			for _, content := range inBag {
				if content.bag == "shiny gold" {
					canHold[bag] = true
					unknownSubBags--
				}
				subBagCanHold, subBagKnown := canHold[content.bag]
				if subBagKnown {
					if subBagCanHold {
						canHold[bag] = true
					}
					unknownSubBags--
				}
			}
			if unknownSubBags == 0 && !canHold[bag] {
				canHold[bag] = false
			}
		}
	}

	count := 0
	for _, v := range canHold {
		if v {
			count++
		}
	}
	return count
}

func (reg *regulations) howManyInShinyGoldBag() int {
	howManyIn := make(map[color]int)

	for len(howManyIn) < len(*reg) {
		for bag, inBag := range *reg {
			if _, known := howManyIn[bag]; known {
				continue
			}
			unknownSubBags := len(inBag)
			bagsInThisBag := 0
			for _, content := range inBag {
				subBagBags, subBagKnown := howManyIn[content.bag]
				if subBagKnown {
					bagsInThisBag += content.num * (1 + subBagBags)
					unknownSubBags--
				}
			}
			if unknownSubBags == 0 {
				howManyIn[bag] = bagsInThisBag
			}
		}
	}

	return howManyIn["shiny gold"]
}

var lineRE = regexp.MustCompile(`(.+) bags contain (?:(no other bags)|(.+))\.`)
var contentsRE = regexp.MustCompile(`(\d+) (.+) bags?`)

func parse(in string) regulations {
	result := make(regulations)

	for _, line := range strings.Split(in, "\n") {
		matches := lineRE.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic(fmt.Sprintf("Line %v doesn't match", line))
		}

		outerColor := color(matches[1])
		if matches[2] == "no other bags" {
			result[outerColor] = nil
			continue
		}

		inside := matches[3]
		resultInner := make([]contents, 0)
		for _, innerBagStr := range strings.Split(inside, ",") {
			matches := contentsRE.FindStringSubmatch(innerBagStr)
			if len(matches) != 3 {
				panic(fmt.Sprintf("Inner contents %v don't match", innerBagStr))
			}
			num, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			resultInner = append(resultInner, contents{
				num: num,
				bag: color(matches[2]),
			})
		}
		result[outerColor] = resultInner

	}

	return result
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	regulations := parse(string(input))

	fmt.Println(regulations.howManyCanHoldShinyGoldBag())
	fmt.Println(regulations.howManyInShinyGoldBag())
}
