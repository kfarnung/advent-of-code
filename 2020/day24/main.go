package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

var directions = []string{
	"ne",
	"nw",
	"se",
	"sw",
	"e",
	"w",
}

func hexOffset(current lib.Point2D, direction string) lib.Point2D {
	switch direction {
	case "ne":
		if current.Y%2 == 0 {
			current.Add(0, -1)
		} else {
			current.Add(1, -1)
		}

	case "nw":
		if current.Y%2 == 0 {
			current.Add(-1, -1)
		} else {
			current.Add(0, -1)
		}

	case "se":
		if current.Y%2 == 0 {
			current.Add(0, 1)
		} else {
			current.Add(1, 1)
		}

	case "sw":
		if current.Y%2 == 0 {
			current.Add(-1, 1)
		} else {
			current.Add(0, 1)
		}

	case "e":
		current.Add(1, 0)

	case "w":
		current.Add(-1, 0)

	default:
		panic("Unexpected direction specified")
	}

	return current
}

func hexNeighbors(current lib.Point2D) []lib.Point2D {
	var neighbors []lib.Point2D
	for _, direction := range directions {
		neighbors = append(neighbors, hexOffset(current, direction))
	}

	return neighbors
}

type hexGrid map[lib.Point2D]bool

func (h *hexGrid) flipTiles(stepsLists [][]string) {
	for _, steps := range stepsLists {
		h.flipTile(steps)
	}
}

func (h *hexGrid) flipTile(steps []string) {
	var position lib.Point2D
	for _, step := range steps {
		position = hexOffset(position, step)
	}

	(*h)[position] = !(*h)[position]
}

func (h *hexGrid) runDay() {
	newGrid := make(hexGrid)
	var searchList []lib.Point2D
	for key := range *h {
		searchList = append(searchList, key)
		searchList = append(searchList, hexNeighbors(key)...)
	}

	visited := make(map[lib.Point2D]bool)
	for _, current := range searchList {
		if visited[current] {
			continue
		}

		visited[current] = true
		flipped := (*h)[current]
		flippedNeighborCount := h.countFlippedNeighbors(current)
		if (!flipped && flippedNeighborCount == 2) ||
			(flipped && (flippedNeighborCount != 0 && flippedNeighborCount <= 2)) {
			newGrid[current] = true
		}
	}

	*h = newGrid
}

func (h hexGrid) countFlipped() int {
	count := 0
	for _, value := range h {
		if value {
			count++
		}
	}

	return count
}

func (h hexGrid) countFlippedNeighbors(current lib.Point2D) int {
	count := 0
	for _, neighbor := range hexNeighbors(current) {
		if h[neighbor] {
			count++
		}
	}

	return count
}

func parseLine(line string) []string {
	var directionList []string
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case 'n':
			directionList = append(directionList, line[i:i+2])
			i++

		case 's':
			directionList = append(directionList, line[i:i+2])
			i++

		case 'e':
			directionList = append(directionList, line[i:i+1])

		case 'w':
			directionList = append(directionList, line[i:i+1])

		default:
			panic("Unexpected input found")
		}
	}

	return directionList
}

func parseInput(lines []string) ([][]string, error) {
	var directionLists [][]string
	for _, line := range lines {
		directionList := parseLine(line)
		directionLists = append(directionLists, directionList)
	}

	return directionLists, nil
}

func part1(lines []string) int {
	stepsLists, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	grid := make(hexGrid)
	grid.flipTiles(stepsLists)

	return grid.countFlipped()
}

func part2(lines []string) int {
	stepsLists, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	grid := make(hexGrid)
	grid.flipTiles(stepsLists)

	for i := 0; i < 100; i++ {
		grid.runDay()
	}

	return grid.countFlipped()
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
