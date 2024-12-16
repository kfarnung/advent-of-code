package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type point struct {
	x, y int
}

type seenKey struct {
	position  point
	direction int
}

type reindeer struct {
	position  point
	direction int
	score     int64
	path      []point
}

func part1(input string) int64 {
	maze, start, _ := parseInput(input)
	queue := []reindeer{{position: start, direction: 90, score: 0}}
	seen := make(map[seenKey]int64)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		key := seenKey{current.position, current.direction}
		if score, ok := seen[key]; ok && current.score >= score {
			continue
		}

		seen[key] = current.score

		if maze[current.position.x][current.position.y] == 'E' {
			return current.score
		}

		straight := getMove(current.position, current.direction)
		if isValid(straight, maze) {
			queue = append(queue, reindeer{
				position:  straight,
				direction: current.direction,
				score:     current.score + 1})
		}

		left := getMove(current.position, lib.Mod(current.direction-90, 360))
		if isValid(left, maze) {
			queue = append(queue, reindeer{
				position:  left,
				direction: lib.Mod(current.direction-90, 360),
				score:     current.score + 1001})
		}

		right := getMove(current.position, lib.Mod(current.direction+90, 360))
		if isValid(right, maze) {
			queue = append(queue, reindeer{
				position:  right,
				direction: lib.Mod(current.direction+90, 360),
				score:     current.score + 1001})
		}

		slices.SortFunc(queue, func(i, j reindeer) int {
			return int(i.score - j.score)
		})
	}

	log.Fatal("No path found")
	return int64(0)
}

func part2(input string) int64 {
	maze, start, _ := parseInput(input)
	queue := []reindeer{{position: start, direction: 90, score: 0, path: []point{start}}}
	seen := make(map[seenKey]int64)
	var successPaths [][]point
	var successScore int64 = math.MaxInt64

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.score > successScore {
			// We can't get a better score than we already have
			break
		}

		key := seenKey{current.position, current.direction}
		score, ok := seen[key]
		if ok && current.score > score {
			continue
		}

		seen[key] = current.score

		if maze[current.position.x][current.position.y] == 'E' {
			successScore = current.score
			successPaths = append(successPaths, current.path)
		}

		straight := getMove(current.position, current.direction)
		if isValid(straight, maze) {
			queue = append(queue, reindeer{
				position:  straight,
				direction: current.direction,
				score:     current.score + 1,
				path:      duplicatePath(current.path, straight)})
		}

		left := getMove(current.position, lib.Mod(current.direction-90, 360))
		if isValid(left, maze) {
			queue = append(queue, reindeer{
				position:  left,
				direction: lib.Mod(current.direction-90, 360),
				score:     current.score + 1001,
				path:      duplicatePath(current.path, left)})
		}

		right := getMove(current.position, lib.Mod(current.direction+90, 360))
		if isValid(right, maze) {
			queue = append(queue, reindeer{
				position:  right,
				direction: lib.Mod(current.direction+90, 360),
				score:     current.score + 1001,
				path:      duplicatePath(current.path, right)})
		}

		slices.SortFunc(queue, func(i, j reindeer) int {
			return int(i.score - j.score)
		})
	}

	visited := make(map[point]bool)
	for _, path := range successPaths {
		for _, position := range path {
			visited[position] = true
		}
	}

	return int64(len(visited))
}

func duplicatePath(path []point, current point) []point {
	duplicate := append([]point(nil), path...)
	duplicate = append(duplicate, current)
	return duplicate
}

func getMove(position point, direction int) point {
	switch direction {
	case 0:
		return point{position.x - 1, position.y}
	case 90:
		return point{position.x, position.y + 1}
	case 180:
		return point{position.x + 1, position.y}
	case 270:
		return point{position.x, position.y - 1}
	}

	log.Fatalf("Invalid direction: %d", direction)
	return point{}
}

func isValid(position point, maze [][]rune) bool {
	return position.x >= 0 && position.x < len(maze) &&
		position.y >= 0 && position.y < len(maze[position.x]) &&
		maze[position.x][position.y] != '#'
}

func parseInput(input string) ([][]rune, point, point) {
	var data [][]rune
	var start, end point
	for i, line := range lib.SplitLines(input) {
		if len(line) == 0 {
			continue
		}

		if startIndex := strings.Index(line, "S"); startIndex != -1 {
			start = point{i, startIndex}
		}

		if endIndex := strings.Index(line, "E"); endIndex != -1 {
			end = point{i, endIndex}
		}

		data = append(data, []rune(line))
	}

	return data, start, end
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
