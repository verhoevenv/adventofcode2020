package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type navigable interface {
	navigate(is []navigationInstruction)
	getPos() xy
}

func navigateBy(nav navigable, is []navigationInstruction) {
	nav.navigate(is)
}

type xy struct {
	x int
	y int
}

type navigationInstruction struct {
	action rune
	value  int
}

type ship struct {
	pos  xy
	head int
}

func makeShip() *ship {
	return &ship{xy{0, 0}, 0}
}

func (s *ship) navigate(is []navigationInstruction) {
	for _, i := range is {
		switch i.action {
		case 'N':
			s.pos.y += i.value
		case 'S':
			s.pos.y -= i.value
		case 'E':
			s.pos.x += i.value
		case 'W':
			s.pos.x -= i.value
		case 'L':
			s.head += i.value
		case 'R':
			s.head -= i.value
		case 'F':
			sin, cos := math.Sincos(degToRad(s.head))
			s.pos.x += int(cos * float64(i.value))
			s.pos.y += int(sin * float64(i.value))
		}
	}
}

func (s *ship) getPos() xy {
	return s.pos
}

func parse(iStr string) []navigationInstruction {
	result := make([]navigationInstruction, 0)

	for _, line := range strings.Split(iStr, "\n") {
		action := rune(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		result = append(result, navigationInstruction{action, num})
	}

	return result
}

func degToRad(deg int) float64 {
	return float64(deg) / 180 * math.Pi
}

func mDist(p xy) int {
	return abs(p.x) + abs(p.y)
}

// this is why I don't like Go
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	ship := makeShip()
	instructions := parse(string(input))
	navigateBy(ship, instructions)

	fmt.Println(mDist(ship.pos))
}
