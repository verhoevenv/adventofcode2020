package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const dividor = 20201227

func findEncryptionKey(cardPubKey int, doorPubKey int) int {
	loopSize := 1

	subject := 7
	value := 1

	for {
		value = value * subject
		value = value % dividor

		if value == cardPubKey {
			return transform(loopSize, doorPubKey)
		} else if value == doorPubKey {
			return transform(loopSize, cardPubKey)
		}

		loopSize++
	}
}

func transform(loopSize int, subject int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = value * subject
		value = value % dividor
	}
	return value
}

func unsafeAtoi(in string) int {
	v, e := strconv.Atoi(in)
	if e != nil {
		panic(e)
	}
	return v
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	pubKeys := strings.Split(string(input), "\n")

	result := findEncryptionKey(unsafeAtoi(pubKeys[0]), unsafeAtoi(pubKeys[1]))

	fmt.Println(result)
}
