package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type position struct {
	x int
	y int
}

func part1(input string) int64 {
	mappedArea, x, y := parseMap(input)
	positionMap := getVisitedPositions(mappedArea, x, y)

	return int64(len(positionMap))
}

func part2(input string) int64 {
	mappedArea, startX, startY := parseMap(input)
	positionMap := getVisitedPositions(mappedArea, startX, startY)

	count := int64(0)
	for position := range positionMap {
		i, j := position.x, position.y
		if mappedArea[i][j] != '.' {
			continue
		}

		mappedArea[i][j] = '#'

		x, y := startX, startY
		positionMap := make(map[string]bool)
		direction := mappedArea[x][y]

		for {
			if positionMap[fmt.Sprintf("%c,%d,%d", direction, x, y)] {
				count++
				break
			}

			positionMap[fmt.Sprintf("%c,%d,%d", direction, x, y)] = true

			nextX := x
			nextY := y
			var nextDirection rune

			switch direction {
			case '>':
				nextY++
				nextDirection = 'v'
			case '<':
				nextY--
				nextDirection = '^'
			case 'v':
				nextX++
				nextDirection = '<'
			case '^':
				nextX--
				nextDirection = '>'
			default:
				log.Fatal("Invalid direction")
			}

			if nextX < 0 || nextY < 0 || nextX >= len(mappedArea) || nextY >= len(mappedArea[nextX]) {
				break
			} else if mappedArea[nextX][nextY] == '#' {
				direction = nextDirection
			} else {
				x = nextX
				y = nextY
			}
		}

		mappedArea[i][j] = '.'
	}

	return count
}

func getVisitedPositions(mappedArea [][]rune, x, y int) map[position]bool {
	direction := mappedArea[x][y]
	positionMap := make(map[position]bool)

	for {
		positionMap[position{x, y}] = true
		nextX := x
		nextY := y
		var nextDirection rune

		switch direction {
		case '>':
			nextY++
			nextDirection = 'v'
		case '<':
			nextY--
			nextDirection = '^'
		case 'v':
			nextX++
			nextDirection = '<'
		case '^':
			nextX--
			nextDirection = '>'
		default:
			log.Fatal("Invalid direction")
		}

		if nextX < 0 || nextY < 0 || nextX >= len(mappedArea) || nextY >= len(mappedArea[nextX]) {
			break
		} else if mappedArea[nextX][nextY] == '#' {
			direction = nextDirection
		} else {
			x = nextX
			y = nextY
		}
	}

	return positionMap
}

func parseMap(input string) (mappedArea [][]rune, x, y int) {
	for i, line := range lib.SplitLines(input) {
		if line == "" {
			break
		}

		startingPoint := strings.Index(line, "^")
		if startingPoint != -1 {
			x = i
			y = startingPoint
		}

		mappedArea = append(mappedArea, []rune(line))
	}

	return
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
