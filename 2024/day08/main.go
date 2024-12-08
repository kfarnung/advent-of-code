package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type antenna struct {
	location  point
	frequency rune
}

type point struct {
	x, y int
}

func part1(input string) int64 {
	antennas, width, height := parseMap(input)
	antinodes := make(map[point]bool)
	for _, antenna := range antennas {
		for _, antenna2 := range antennas {
			if antenna.location == antenna2.location || antenna.frequency != antenna2.frequency {
				continue
			}

			delta := point{antenna2.location.x - antenna.location.x, antenna2.location.y - antenna.location.y}
			antinode1 := point{antenna.location.x + delta.x, antenna.location.y + delta.y}
			antinode2 := point{antenna.location.x - delta.x, antenna.location.y - delta.y}
			antinode3 := point{antenna2.location.x + delta.x, antenna2.location.y + delta.y}
			antinode4 := point{antenna2.location.x - delta.x, antenna2.location.y - delta.y}

			if antinode1 != antenna2.location && antinode1.x >= 0 && antinode1.x < width && antinode1.y >= 0 && antinode1.y < height {
				antinodes[antinode1] = true
			}

			if antinode2 != antenna2.location && antinode2.x >= 0 && antinode2.x < width && antinode2.y >= 0 && antinode2.y < height {
				antinodes[antinode2] = true
			}

			if antinode3 != antenna.location && antinode3.x >= 0 && antinode3.x < width && antinode3.y >= 0 && antinode3.y < height {
				antinodes[antinode3] = true
			}

			if antinode4 != antenna.location && antinode4.x >= 0 && antinode4.x < width && antinode4.y >= 0 && antinode4.y < height {
				antinodes[antinode4] = true
			}
		}
	}

	return int64(len(antinodes))
}

func part2(input string) int64 {
	antennas, rows, cols := parseMap(input)
	antinodes := make(map[point]bool)
	for _, antenna := range antennas {
		for _, antenna2 := range antennas {
			if antenna.location == antenna2.location || antenna.frequency != antenna2.frequency {
				continue
			}

			delta := point{antenna2.location.x - antenna.location.x, antenna2.location.y - antenna.location.y}
			antinodeLocation := point{antenna.location.x + delta.x, antenna.location.y + delta.y}
			for antinodeLocation.x >= 0 && antinodeLocation.x < rows && antinodeLocation.y >= 0 && antinodeLocation.y < cols {
				antinodes[antinodeLocation] = true
				antinodeLocation = point{antinodeLocation.x + delta.x, antinodeLocation.y + delta.y}
			}

			antinodeLocation = point{antenna.location.x - delta.x, antenna.location.y - delta.y}
			for antinodeLocation.x >= 0 && antinodeLocation.x < rows && antinodeLocation.y >= 0 && antinodeLocation.y < cols {
				antinodes[antinodeLocation] = true
				antinodeLocation = point{antinodeLocation.x - delta.x, antinodeLocation.y - delta.y}
			}
		}
	}

	return int64(len(antinodes))
}

func parseMap(input string) ([]antenna, int, int) {
	var antennas []antenna
	rows := 0
	lines := lib.SplitLines(input)
	for x, line := range lines {
		if len(line) == 0 {
			continue
		}

		rows++

		for y, frequency := range line {
			if frequency != '.' {
				antennas = append(antennas, antenna{point{x, y}, frequency})
			}
		}
	}
	return antennas, rows, len(lines[0])
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
