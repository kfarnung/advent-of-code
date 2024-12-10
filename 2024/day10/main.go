package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

func part1(input string) int64 {
	topoMap, trailHeads := parseMap(input)
	sum := int64(0)
	for _, trailHead := range trailHeads {
		terminalPositions := make(map[lib.Point2D]bool)
		findTerminalPositions(topoMap, trailHead, terminalPositions)

		sum += int64(len(terminalPositions))
	}

	return sum
}

func part2(input string) int64 {
	topoMap, trailHeads := parseMap(input)
	sum := int64(0)
	for _, trailHead := range trailHeads {
		terminalPositions := make(map[lib.Point2D]bool)
		sum += findTerminalPositions(topoMap, trailHead, terminalPositions)
	}

	return sum
}

func findTerminalPositions(topoMap [][]rune, position lib.Point2D, terminalPositions map[lib.Point2D]bool) int64 {
	current := topoMap[position.X][position.Y]
	fmt.Printf("current: %v %v\n", current, position)
	if current == '9' {
		terminalPositions[position] = true
		return 1
	}

	count := int64(0)
	target := current + 1

	if position.X > 0 && topoMap[position.X-1][position.Y] == target {
		count += findTerminalPositions(topoMap, lib.Point2D{X: position.X - 1, Y: position.Y}, terminalPositions)
	}

	if position.X < int64(len(topoMap)-1) && topoMap[position.X+1][position.Y] == target {
		count += findTerminalPositions(topoMap, lib.Point2D{X: position.X + 1, Y: position.Y}, terminalPositions)
	}

	if position.Y > 0 && topoMap[position.X][position.Y-1] == target {
		count += findTerminalPositions(topoMap, lib.Point2D{X: position.X, Y: position.Y - 1}, terminalPositions)
	}

	if position.Y < int64(len(topoMap[0])-1) && topoMap[position.X][position.Y+1] == target {
		count += findTerminalPositions(topoMap, lib.Point2D{X: position.X, Y: position.Y + 1}, terminalPositions)
	}

	return count
}

func parseMap(input string) ([][]rune, []lib.Point2D) {
	var topoMap [][]rune
	var trailHeads []lib.Point2D
	for i, line := range lib.SplitLines(input) {
		if len(line) == 0 {
			continue
		}

		var row []rune
		for j, char := range line {
			row = append(row, char)

			if char == '0' {
				trailHeads = append(trailHeads, lib.Point2D{X: int64(i), Y: int64(j)})
			}
		}

		topoMap = append(topoMap, row)
	}

	return topoMap, trailHeads
}

func main() {
	name := os.Args[1]
	content, err := lib.LoadFileContent(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(content))
	fmt.Printf("Part 2: %d\n", part2(content))
}
