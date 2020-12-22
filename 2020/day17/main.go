package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func runCycle(grid map[lib.Point3D]bool) map[lib.Point3D]bool {
	newGrid := make(map[lib.Point3D]bool)

	var searchSpace []lib.Point3D
	for point, active := range grid {
		if active {
			searchSpace = append(searchSpace, point)
			searchSpace = append(searchSpace, point.GetAdjacent()...)
		}
	}

	for _, searchPoint := range searchSpace {
		if _, ok := newGrid[searchPoint]; ok {
			// The point was already evaluated, skip it
			continue
		}

		activeAdjacent := 0
		for _, adjacent := range searchPoint.GetAdjacent() {
			if grid[adjacent] {
				activeAdjacent++
			}
		}

		newGrid[searchPoint] = activeAdjacent == 3 || (grid[searchPoint] && activeAdjacent == 2)
	}

	return newGrid
}

func runCycle4D(grid map[lib.Point4D]bool) map[lib.Point4D]bool {
	newGrid := make(map[lib.Point4D]bool)

	var searchSpace []lib.Point4D
	for point, active := range grid {
		if active {
			searchSpace = append(searchSpace, point)
			searchSpace = append(searchSpace, point.GetAdjacent()...)
		}
	}

	for _, searchPoint := range searchSpace {
		if _, ok := newGrid[searchPoint]; ok {
			// The point was already evaluated, skip it
			continue
		}

		activeAdjacent := 0
		for _, adjacent := range searchPoint.GetAdjacent() {
			if grid[adjacent] {
				activeAdjacent++
			}
		}

		newGrid[searchPoint] = activeAdjacent == 3 || (grid[searchPoint] && activeAdjacent == 2)
	}

	return newGrid
}

func convertToPoint4D(grid map[lib.Point3D]bool) map[lib.Point4D]bool {
	newGrid := make(map[lib.Point4D]bool)
	for point, active := range grid {
		newPoint := lib.NewPoint4D(point.X, point.Y, point.Z, 0)
		newGrid[newPoint] = active
	}

	return newGrid
}

func parseInput(lines []string) map[lib.Point3D]bool {
	grid := make(map[lib.Point3D]bool)

	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				point := lib.NewPoint3D(int64(i), int64(j), 0)
				grid[point] = true
			}
		}
	}

	return grid
}

func part1(lines []string) int {
	grid := parseInput(lines)

	for i := 0; i < 6; i++ {
		grid = runCycle(grid)
	}

	count := 0
	for _, value := range grid {
		if value {
			count++
		}
	}

	return count
}

func part2(lines []string) int {
	tempGrid := parseInput(lines)
	grid := convertToPoint4D(tempGrid)

	for i := 0; i < 6; i++ {
		grid = runCycle4D(grid)
	}

	count := 0
	for _, value := range grid {
		if value {
			count++
		}
	}

	return count
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
