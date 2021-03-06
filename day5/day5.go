package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type boardingPass string
type seat struct {
	row    int
	col    int
	seatID int
}

func newSeat(row int, col int) seat {
	return seat{
		row:    row,
		col:    col,
		seatID: row*8 + col,
	}
}

func deriveSeat(pass boardingPass) seat {
	lowR := 0
	highR := 127
	for _, c := range string(pass)[0:7] {
		switch c {
		case 'F':
			highR = highR - ((highR - lowR + 1) / 2)
		case 'B':
			lowR = lowR + ((highR - lowR + 1) / 2)
		default:
			panic(fmt.Sprintf("Unexpected character %c in %v", c, pass))
		}
	}

	lowC := 0
	highC := 7
	for _, c := range string(pass)[7:10] {
		switch c {
		case 'L':
			highC = highC - ((highC - lowC + 1) / 2)
		case 'R':
			lowC = lowC + ((highC - lowC + 1) / 2)
		default:
			panic(fmt.Sprintf("Unexpected character %c in %v", c, pass))
		}
	}

	if lowR != highR && lowC != highC {
		panic(fmt.Sprintf("seats should converge, got (%v-%v), (%v-%v)",
			lowR, highR, lowC, highC))
	}
	return newSeat(lowR, lowC)
}

// mutates argument!
func findMissingSeat(seats []seat) int {
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].seatID < seats[j].seatID
	})

	previousSeatID := seats[0].seatID
	for _, seat := range seats[1:] {
		if seat.seatID != previousSeatID+1 {
			return previousSeatID + 1
		}
		previousSeatID = seat.seatID
	}
	panic("Didn't find missing seat")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	allSeats := make([]seat, 0)

	for scanner.Scan() {
		pass := boardingPass(scanner.Text())
		seat := deriveSeat(pass)
		allSeats = append(allSeats, seat)
	}

	fmt.Println(findMissingSeat(allSeats))
}
