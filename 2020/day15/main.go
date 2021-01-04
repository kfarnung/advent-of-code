package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func playGame(numbers []int32, numTurns int32) int32 {
	numberHistory := make(map[int32]int32)

	// Fill the history with the previously seen numbers.
	for i, num := range numbers[:len(numbers)-1] {
		numberHistory[num] = int32(i)
	}

	// Start with the last number from the list.
	mostRecentNumber := numbers[len(numbers)-1]

	for turn := int32(len(numbers)); turn < numTurns; turn++ {
		// Look up the most recent number and then update the history.
		history, found := numberHistory[mostRecentNumber]
		numberHistory[mostRecentNumber] = turn - 1

		if found {
			// Calculate the next number as the difference between last turn and
			// the previous time it was seen.
			mostRecentNumber = turn - history - 1
		} else {
			// Never seen before, default to zero.
			mostRecentNumber = 0
		}
	}

	return mostRecentNumber
}

func parseInput(lines []string) ([]int32, error) {
	var numbers []int32
	for _, num := range strings.Split(lines[0], ",") {
		numValue, err := lib.ParseInt32(num)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, numValue)
	}

	return numbers, nil
}

func part1(lines []string) int32 {
	numbers, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	return playGame(numbers, 2020)
}

func part2(lines []string) int32 {
	numbers, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	return playGame(numbers, 30000000)
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
