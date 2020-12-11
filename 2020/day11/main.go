package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

const seatOccupied = '#'
const seatEmpty = 'L'
const floor = '.'

type slope struct {
	drow    int
	dcolumn int
}

var slopes = []slope{
	{drow: -1, dcolumn: -1},
	{drow: -1, dcolumn: 0},
	{drow: -1, dcolumn: 1},
	{drow: 0, dcolumn: -1},
	{drow: 0, dcolumn: 1},
	{drow: 1, dcolumn: -1},
	{drow: 1, dcolumn: 0},
	{drow: 1, dcolumn: 1},
}

func countAdjacentOccupied(grid [][]rune, row int, column int) int {
	count := 0
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i == row && j == column {
				continue
			}

			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == '#' {
				count++
			}
		}
	}

	return count
}

func countOccupied(grid [][]rune) int {
	count := 0
	for _, row := range grid {
		for _, current := range row {
			if current == seatOccupied {
				count++
			}
		}
	}

	return count
}

func isDirectionOccupied(grid [][]rune, row int, column int, drow int, dcolumn int) bool {
	for i, j := row+drow, column+dcolumn; i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]); i, j = i+drow, j+dcolumn {
		current := grid[i][j]
		if current == seatOccupied {
			return true
		} else if current == seatEmpty {
			return false
		}
	}

	return false
}

func countDirectionOccupied(grid [][]rune, row int, column int) int {
	count := 0

	for _, slope := range slopes {
		if isDirectionOccupied(grid, row, column, slope.drow, slope.dcolumn) {
			count++
		}
	}

	return count
}

func parseInput(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
		}
	}

	return grid
}

func findStability(lines []string, occupiedCounter func([][]rune, int, int) int, occupiedLimit int) int {
	grid := parseInput(lines)

	for {
		changes := 0
		newGrid := make([][]rune, len(grid))

		for i, row := range grid {
			newGrid[i] = make([]rune, len(row))

			for j, current := range row {
				if current != floor {
					occupied := occupiedCounter(grid, i, j)

					if current == seatEmpty && occupied == 0 {
						current = seatOccupied
						changes++
					} else if current == seatOccupied && occupied >= occupiedLimit {
						current = seatEmpty
						changes++
					}
				}

				newGrid[i][j] = current
			}
		}

		if changes == 0 {
			break
		}

		grid = newGrid
	}

	return countOccupied(grid)
}

func part1(lines []string) int {
	return findStability(lines, countAdjacentOccupied, 4)
}

func part2(lines []string) int {
	return findStability(lines, countDirectionOccupied, 5)
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
