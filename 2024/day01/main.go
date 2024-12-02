package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

var lineParseRegex = regexp.MustCompile(`^(\d+) +(\d+)$`)

func part1(input string) int64 {
	first, second, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	lib.SortSliceAscending(first)
	lib.SortSliceAscending(second)

	sum := int64(0)
	for i := 0; i < len(first); i++ {
		sum += lib.Abs(first[i] - second[i])
	}

	return sum
}

func part2(input string) int64 {
	first, second, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := int64(0)
	for _, value := range first {
		count := int64(0)
		for _, otherValue := range second {
			if value == otherValue {
				count++
			}
		}

		sum += value * count
	}

	return sum
}

func parseInput(input string) ([]int64, []int64, error) {
	var first []int64
	var second []int64

	for _, line := range lib.SplitLines(input) {
		matches := lineParseRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		firstValue, err := lib.ParseInt[int64](matches[1])
		if err != nil {
			return nil, nil, err
		}

		secondValue, err := lib.ParseInt[int64](matches[2])
		if err != nil {
			return nil, nil, err
		}

		first = append(first, firstValue)
		second = append(second, secondValue)
	}

	return first, second, nil
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
