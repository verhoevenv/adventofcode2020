package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type hexCoord struct {
	x int
	y int
	z int
}

func parseCoordinates(in string) []hexCoord {
	result := make([]hexCoord, 0)

	for _, line := range strings.Split(in, "\n") {
		toRead := line
		hex := hexCoord{0, 0, 0}
		for toRead != "" {
			switch {
			case strings.HasPrefix(toRead, "e"):
				toRead = toRead[1:]
				hex.x++
				hex.y--
			case strings.HasPrefix(toRead, "se"):
				toRead = toRead[2:]
				hex.y--
				hex.z++
			case strings.HasPrefix(toRead, "sw"):
				toRead = toRead[2:]
				hex.x--
				hex.z++
			case strings.HasPrefix(toRead, "w"):
				toRead = toRead[1:]
				hex.x--
				hex.y++
			case strings.HasPrefix(toRead, "nw"):
				toRead = toRead[2:]
				hex.y++
				hex.z--
			case strings.HasPrefix(toRead, "ne"):
				toRead = toRead[2:]
				hex.x++
				hex.z--
			default:
				panic(fmt.Sprintf("cannot parse %v (part of %v)", toRead, line))
			}
		}
		result = append(result, hex)
	}

	return result
}

func flipAllAndCount(hexes []hexCoord) int {
	flipped := make(map[hexCoord]bool)

	for _, hex := range hexes {
		flipped[hex] = !flipped[hex]
	}

	count := 0
	for _, v := range flipped {
		if v {
			count++
		}
	}
	return count
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	tiles := parseCoordinates(string(input))

	result := flipAllAndCount(tiles)

	fmt.Println(result)
}
