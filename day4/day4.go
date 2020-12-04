package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
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

func validOld(pass *passport) bool {
	for _, f := range requiredFields {
		if _, ok := pass.fields[f]; !ok {
			return false
		}
	}
	return true
}

type validationRule = func(string) bool

var validationRules = map[string]validationRule{
	"byr": numInRange(1920, 2002),
	"iyr": numInRange(2010, 2020),
	"eyr": numInRange(2020, 2030),
	"hgt": validHgt,
	"hcl": matches(`^#[0-9a-f]{6}$`),
	"ecl": matches(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": matches(`^[0-9]{9}$`),
}

func numInRange(low int, high int) func(string) bool {
	regex := regexp.MustCompile(`^\d{4}$`)
	return func(val string) bool {
		if !regex.MatchString(val) {
			return false
		}
		valNum, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		if valNum < low {
			return false
		}
		if valNum > high {
			return false
		}
		return true
	}
}

func matches(regexS string) func(string) bool {
	regex := regexp.MustCompile(regexS)
	return func(val string) bool {
		if !regex.MatchString(val) {
			return false
		}
		return true
	}
}

func validHgt(hgt string) bool {
	regex := regexp.MustCompile(`^(\d+)(in|cm)$`)
	matches := regex.FindStringSubmatch(hgt)
	if len(matches) != 3 {
		return false
	}
	num, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	if matches[2] == "cm" {
		return num >= 150 && num <= 193
	}
	return num >= 59 && num <= 76
}

func validStrict(pass *passport) bool {
	for key, rule := range validationRules {
		val, present := pass.fields[key]
		if !present {
			return false
		}
		if !rule(val) {
			return false
		}
	}

	return true
}

func count(ps []passport, validation func(*passport) bool) int {
	nbValid := 0
	for _, p := range ps {
		if validation(&p) {
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
				panic(fmt.Sprintf(
					"%v does not contain k/v pair! During parsing of passport %v",
					field, ppString))
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

	fmt.Println(count(ps, validStrict))
}
