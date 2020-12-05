package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func partition(start, end int, direction rune) (int, int) {
	if end <= start {
		panic("End must be greater than start")
	}

	mid := start + ((end - start) / 2)
	if direction == 'F' || direction == 'L' {
		return start, mid
	} else if direction == 'B' || direction == 'R' {
		return mid + 1, end
	}

	panic("Unexpected direction provided")
}

func followPath(start, end int, path string) int {
	for _, char := range path {
		start, end = partition(start, end, char)
	}

	if start != end {
		panic("Start and end must converge on a single value")
	}

	return start
}

func calculateSeatID(path string) int {
	row := followPath(0, 127, path[:7])
	col := followPath(0, 7, path[7:])
	return row*8 + col
}

func part1(lines []string) int {
	maxSeatID := 0
	for _, line := range lines {
		seatID := calculateSeatID(line)

		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID
}

func part2(lines []string) int {
	occupiedSeats := make([]bool, 128*8)
	for _, line := range lines {
		seatID := calculateSeatID(line)
		occupiedSeats[seatID] = true
	}

	seatID := 0
	for ; !occupiedSeats[seatID]; seatID++ {
		// Find the first occupied seat
	}

	for ; occupiedSeats[seatID]; seatID++ {
		// Find the first unoccupied seat
	}

	return seatID
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
