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
			s.pos.x += int(math.Round(cos * float64(i.value)))
			s.pos.y += int(math.Round(sin * float64(i.value)))
		}
	}
}

func (s *ship) getPos() xy {
	return s.pos
}

type waypoint struct {
	shipPos        xy
	waypointRelPos xy
}

func makeWaypoint() *waypoint {
	return &waypoint{xy{0, 0}, xy{10, 1}}
}

func (w *waypoint) navigate(is []navigationInstruction) {
	for _, i := range is {
		switch i.action {
		case 'N':
			w.waypointRelPos.y += i.value
		case 'S':
			w.waypointRelPos.y -= i.value
		case 'E':
			w.waypointRelPos.x += i.value
		case 'W':
			w.waypointRelPos.x -= i.value
		case 'L':
			w.waypointRelPos = rotate(w.waypointRelPos, i.value)
		case 'R':
			w.waypointRelPos = rotate(w.waypointRelPos, -i.value)
		case 'F':
			d := scale(w.waypointRelPos, i.value)
			w.shipPos.x += d.x
			w.shipPos.y += d.y
		}
	}
}

func (w *waypoint) getPos() xy {
	return w.shipPos
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

func rotate(vec xy, degreeAngle int) xy {
	sin, cos := math.Sincos(degToRad(degreeAngle))
	return xy{
		int(math.Round(float64(vec.x)*cos - float64(vec.y)*sin)),
		int(math.Round(float64(vec.x)*sin + float64(vec.y)*cos)),
	}
}

func scale(vec xy, scalar int) xy {
	return xy{
		vec.x * scalar,
		vec.y * scalar,
	}
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

	ship := makeWaypoint()
	instructions := parse(string(input))
	navigateBy(ship, instructions)

	fmt.Println(mDist(ship.getPos()))
}
