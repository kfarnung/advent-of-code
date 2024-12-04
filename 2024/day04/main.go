package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

var xmasString []rune = []rune("XMAS")

func part1(input string) int64 {
	parsedInput := parseInput(input)
	count := int64(0)
	for x, line := range parsedInput {
		for y, char := range line {
			if char != 'X' {
				continue
			}

			if findXmas(parsedInput, x, y, 1, 0) {
				count++
			}

			if findXmas(parsedInput, x, y, 0, 1) {
				count++
			}

			if findXmas(parsedInput, x, y, 1, 1) {
				count++
			}

			if findXmas(parsedInput, x, y, 1, -1) {
				count++
			}

			if findXmas(parsedInput, x, y, -1, 0) {
				count++
			}

			if findXmas(parsedInput, x, y, 0, -1) {
				count++
			}

			if findXmas(parsedInput, x, y, -1, -1) {
				count++
			}

			if findXmas(parsedInput, x, y, -1, 1) {
				count++
			}
		}
	}

	return count
}

func part2(input string) int64 {
	parsedInput := parseInput(input)
	count := int64(0)
	for x, line := range parsedInput {
		for y, char := range line {
			if char != 'A' {
				continue
			}

			if x == 0 || y == 0 || x == len(parsedInput)-1 || y == len(parsedInput[x])-1 {
				continue
			}

			if parsedInput[x-1][y-1] == 'M' && parsedInput[x+1][y+1] == 'S' && parsedInput[x-1][y+1] == 'M' && parsedInput[x+1][y-1] == 'S' {
				count++
			}

			if parsedInput[x-1][y-1] == 'M' && parsedInput[x+1][y+1] == 'S' && parsedInput[x-1][y+1] == 'S' && parsedInput[x+1][y-1] == 'M' {
				count++
			}

			if parsedInput[x-1][y-1] == 'S' && parsedInput[x+1][y+1] == 'M' && parsedInput[x-1][y+1] == 'S' && parsedInput[x+1][y-1] == 'M' {
				count++
			}

			if parsedInput[x-1][y-1] == 'S' && parsedInput[x+1][y+1] == 'M' && parsedInput[x-1][y+1] == 'M' && parsedInput[x+1][y-1] == 'S' {
				count++
			}
		}
	}

	return count
}

func findXmas(parsedInput [][]rune, x, y int, dx, dy int) bool {
	for i := 0; i < len(xmasString); i++ {
		if x < 0 || x >= len(parsedInput) || y < 0 || y >= len(parsedInput[x]) || parsedInput[x][y] != xmasString[i] {
			return false
		}

		x += dx
		y += dy
	}

	return true
}

func parseInput(input string) [][]rune {
	var result [][]rune
	for _, line := range lib.SplitLines(input) {
		if len(line) == 0 {
			continue
		}

		result = append(result, []rune(line))
	}

	return result
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
