package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

func part1(input string) int64 {
	first, second, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	lib.SortSliceInt64(first)
	lib.SortSliceInt64(second)

	sum := int64(0)
	for i := 0; i < len(first); i++ {
		sum += lib.AbsInt64(first[i] - second[i])
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
	lines := lib.SplitLines(input)
	for _, line := range lines {
		if (len(line)) == 0 {
			continue
		}

		splitLine := strings.Split(line, " ")
		firstValue, err := lib.ParseInt64(splitLine[0])
		if err != nil {
			return nil, nil, err
		}

		secondValue, err := lib.ParseInt64(splitLine[len(splitLine)-1])
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
