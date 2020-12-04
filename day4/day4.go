package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type passport struct {
	fields map[string]string
}

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func valid(pass *passport) bool {
	for _, f := range requiredFields {
		if _, ok := pass.fields[f]; !ok {
			return false
		}
	}
	return true
}

func countValid(ps []passport) int {
	nbValid := 0
	for _, p := range ps {
		if valid(&p) {
			nbValid++
		}
	}
	return nbValid
}

var whiteSpace = regexp.MustCompile("\\s")

func parse(in string) []passport {
	result := make([]passport, 0)

	ppStrings := strings.Split(in, "\n\n")
	for _, ppString := range ppStrings {
		fieldStrings := whiteSpace.Split(ppString, -1)

		fields := make(map[string]string)

		for _, field := range fieldStrings {
			kv := strings.Split(field, ":")
			if len(kv) != 2 {
				panic(fmt.Sprintf("%v does not contain k/v pair!", field))
			}
			fields[kv[0]] = kv[1]
		}

		pp := passport{fields}
		result = append(result, pp)
	}

	return result
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	ps := parse(string(input))

	fmt.Println(countValid(ps))
}
