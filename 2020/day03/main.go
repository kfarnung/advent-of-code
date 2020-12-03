package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func createGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))

	for row, line := range lines {
		grid[row] = make([]rune, len(line))

		for col, char := range line {
			grid[row][col] = char
		}
	}

	return grid
}

func followSlope(grid [][]rune, right int, down int) int {
	rows := len(grid)
	treeCount := 0

	for row, col := 0, 0; row < rows; row, col = row+down, col+right {
		gridRow := grid[row]
		char := gridRow[col%len(gridRow)]
		if char == '#' {
			treeCount++
		}
	}

	return treeCount
}

func part1(grid [][]rune) int {
	return followSlope(grid, 3, 1)
}

func part2(grid [][]rune) int {
	return followSlope(grid, 1, 1) *
		followSlope(grid, 3, 1) *
		followSlope(grid, 5, 1) *
		followSlope(grid, 7, 1) *
		followSlope(grid, 1, 2)
}

func parseInput(name string) ([][]rune, error) {
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		return nil, err
	}

	return createGrid(lines), nil
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	grid := createGrid(lines)
	fmt.Printf("Part 1: %d\n", part1(grid))
	fmt.Printf("Part 2: %d\n", part2(grid))
}
