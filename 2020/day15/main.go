package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func playGame(numbers []int32, numTurns int32) int32 {
	var turn, mostRecentNumber int32
	numberHistory := make(map[int32][]int32)
	for i, num := range numbers {
		mostRecentNumber = num
		numberHistory[num] = append(numberHistory[num], int32(i))
		turn++
	}

	for ; turn < numTurns; turn++ {
		history := numberHistory[mostRecentNumber]
		if len(history) == 1 {
			mostRecentNumber = 0
		} else {
			mostRecentNumber = history[len(history)-1] - history[len(history)-2]
		}

		numberHistory[mostRecentNumber] = append(numberHistory[mostRecentNumber], turn)
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
