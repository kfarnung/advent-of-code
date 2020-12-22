package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type parsedLine struct {
	min      int32
	max      int32
	char     rune
	password string
}

func parseLines(lines []string) ([]parsedLine, error) {
	lineParse := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	parsedLines := make([]parsedLine, len(lines))

	for i, line := range lines {
		parsedLine := lineParse.FindStringSubmatch(line)
		if parsedLine == nil {
			return nil, errors.New("Failed to parse line")
		}

		min, err := lib.ParseInt32(parsedLine[1])
		if err != nil {
			return nil, err
		}

		max, err := lib.ParseInt32(parsedLine[2])
		if err != nil {
			return nil, err
		}

		parsedLines[i].min = min
		parsedLines[i].max = max
		parsedLines[i].char = []rune(parsedLine[3])[0]
		parsedLines[i].password = parsedLine[4]
	}

	return parsedLines, nil
}

func countMatchingRunes(text string, char rune) int32 {
	var count int32
	for _, currentChar := range text {
		if currentChar == char {
			count++
		}
	}

	return count
}

func part1(parsedLines []parsedLine) int {
	count := 0

	for _, line := range parsedLines {
		charCount := countMatchingRunes(line.password, line.char)
		if line.min <= charCount && charCount <= line.max {
			count++
		}
	}

	return count
}

func part2(parsedLines []parsedLine) int {
	count := 0

	for _, line := range parsedLines {
		runes := []rune(line.password)
		if (runes[line.min-1] == line.char) != (runes[line.max-1] == line.char) {
			count++
		}
	}

	return count
}

func parseInput(name string) ([]parsedLine, error) {
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		return nil, err
	}

	parsedLines, err := parseLines(lines)
	if err != nil {
		return nil, err
	}

	return parsedLines, nil
}

func main() {
	name := os.Args[1]
	parsedLines, err := parseInput(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(parsedLines))
	fmt.Printf("Part 2: %d\n", part2(parsedLines))
}
