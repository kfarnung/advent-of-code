package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type point struct {
	x, y int
}

type mazeState struct {
	position  point
	stepCount int
}

func part1(input string, dimension, count int) int64 {
	points, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	grid := initializeGrid(dimension)
	corruptGrid(grid, points, count)

	return int64(findShortestPath(grid, point{0, 0}, point{dimension, dimension}))
}

func part2(input string, dimension int) string {
	points, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	grid := initializeGrid(dimension)

	for _, p := range points {
		grid[p.y][p.x] = '#'

		if findShortestPath(grid, point{0, 0}, point{dimension, dimension}) == 0 {
			return fmt.Sprintf("%d,%d", p.x, p.y)
		}
	}

	log.Fatal("No solution found")
	return ""
}

func findShortestPath(grid [][]rune, start, end point) int {
	queue := []mazeState{{start, 0}}
	steps := make(map[point]int)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if stepCount, ok := steps[current.position]; ok && stepCount <= current.stepCount {
			continue
		}

		steps[current.position] = current.stepCount

		if current.position == end {
			break
		}

		for _, direction := range []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := point{current.position.x + direction.x, current.position.y + direction.y}
			if next.y >= 0 && next.y < len(grid) && next.x >= 0 && next.x < len(grid[next.y]) && grid[next.y][next.x] == '.' {
				queue = append(queue, mazeState{next, current.stepCount + 1})
			}
		}
	}

	return steps[end]
}

func corruptGrid(grid [][]rune, points []point, count int) {
	for i := 0; i < count; i++ {
		grid[points[i].y][points[i].x] = '#'
	}
}

func initializeGrid(dimension int) [][]rune {
	grid := make([][]rune, dimension+1)
	for i := range grid {
		grid[i] = make([]rune, dimension+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	return grid
}

func parseInput(input string) ([]point, error) {
	var points []point
	for _, line := range lib.SplitLines(input) {
		if line == "" {
			continue
		}

		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			return nil, err
		}
		points = append(points, point{x, y})
	}

	return points, nil
}

func main() {
	name := os.Args[1]
	content, err := lib.LoadFileContent(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(content, 70, 1024))
	fmt.Printf("Part 2: %s\n", part2(content, 70))
}
