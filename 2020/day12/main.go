package main

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type navigationInstruction struct {
	action rune
	value  int64
}

func parseInput(lines []string) ([]navigationInstruction, error) {
	instructions := make([]navigationInstruction, len(lines))

	for i, line := range lines {
		action, width := utf8.DecodeRuneInString(line)
		value, err := lib.ParseInt64(line[width:])
		if err != nil {
			return nil, err
		}

		instructions[i] = navigationInstruction{
			action: action,
			value:  value,
		}
	}

	return instructions, nil
}

func part1(lines []string) int64 {
	instructions, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	ship := lib.Point2D{}
	slope := lib.NewPoint2D(1, 0)

	for _, instruction := range instructions {
		switch instruction.action {
		case 'N':
			ship.Add(0, instruction.value)
		case 'S':
			ship.Add(0, -instruction.value)
		case 'E':
			ship.Add(instruction.value, 0)
		case 'W':
			ship.Add(-instruction.value, 0)
		case 'L':
			slope.Rotate90DegreesCounterClockwise(int(instruction.value / 90))
		case 'R':
			slope.Rotate90DegreesClockwise(int(instruction.value / 90))
		case 'F':
			ship.Add(slope.X*instruction.value, slope.Y*instruction.value)
		}
	}

	return ship.ManhattanDistance(lib.Point2D{})
}

func part2(lines []string) int64 {
	instructions, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	ship := lib.Point2D{}
	waypoint := lib.NewPoint2D(10, 1)

	for _, instruction := range instructions {
		switch instruction.action {
		case 'N':
			waypoint.Add(0, instruction.value)
		case 'S':
			waypoint.Add(0, -instruction.value)
		case 'E':
			waypoint.Add(instruction.value, 0)
		case 'W':
			waypoint.Add(-instruction.value, 0)
		case 'L':
			waypoint.Rotate90DegreesCounterClockwise(int(instruction.value / 90))
		case 'R':
			waypoint.Rotate90DegreesClockwise(int(instruction.value / 90))
		case 'F':
			ship.Add(waypoint.X*instruction.value, waypoint.Y*instruction.value)
		}
	}

	return ship.ManhattanDistance(lib.Point2D{})
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
