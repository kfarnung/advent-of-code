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

func part1(input string) int64 {
	garden := parseInput(input)

	sum := int64(0)
	seen := make(map[point]bool)

	for i, row := range garden {
		for j := range row {
			location := point{i, j}
			if seen[location] {
				continue
			}

			area, perimeter := findAreaAndPerimeter(garden, i, j, seen)
			sum += area * perimeter
		}
	}

	return sum
}

func part2(_ string) int64 {
	return int64(0)
}

func findAreaAndPerimeter(garden [][]rune, i, j int, seen map[point]bool) (int64, int64) {
	area := int64(0)
	perimeter := int64(0)
	region := garden[i][j]

	var queue []point
	queue = append(queue, point{x: i, y: j})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if seen[current] {
			continue
		}

		seen[current] = true
		area++

		if current.x > 0 && garden[current.x-1][current.y] == region {
			queue = append(queue, point{x: current.x - 1, y: current.y})
		} else {
			perimeter++
		}

		if current.x < len(garden)-1 && garden[current.x+1][current.y] == region {
			queue = append(queue, point{x: current.x + 1, y: current.y})
		} else {
			perimeter++
		}

		if current.y > 0 && garden[current.x][current.y-1] == region {
			queue = append(queue, point{x: current.x, y: current.y - 1})
		} else {
			perimeter++
		}

		if current.y < len(garden[0])-1 && garden[current.x][current.y+1] == region {
			queue = append(queue, point{x: current.x, y: current.y + 1})
		} else {
			perimeter++
		}
	}

	return area, perimeter
}

func parseInput(input string) [][]rune {
	var data [][]rune
	for _, line := range lib.SplitLines(input) {
		if len(line) == 0 {
			continue
		}

		data = append(data, []rune(line))
	}

	return data
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
