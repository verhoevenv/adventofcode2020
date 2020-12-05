package main

import "testing"

func TestDeriveSeat(t *testing.T) {

	tables := []struct {
		pass         boardingPass
		expectedSeat seat
	}{
		{boardingPass("FBFBBFFRLR"), seat{row: 44, col: 5, seatID: 357}},
		{boardingPass("BFFFBBFRRR"), seat{row: 70, col: 7, seatID: 567}},
		{boardingPass("FFFBBBFRRR"), seat{row: 14, col: 7, seatID: 119}},
		{boardingPass("BBFFBBFRLL"), seat{row: 102, col: 4, seatID: 820}},
	}

	for _, table := range tables {
		result := deriveSeat(table.pass)
		if result != table.expectedSeat {
			t.Errorf("Derivation of %v was incorrect, got: %v, want: %v.",
				table.pass, result, table.expectedSeat)
		}
	}

}
