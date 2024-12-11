package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type mapKey struct {
	integer int64
	steps   int
}

func part1(input string) int64 {
	integers, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	count := int64(0)
	seen := make(map[mapKey]int64)

	for j := 0; j < len(integers); j++ {
		count += recurse(integers[j], 25, seen)
	}

	return count
}

func part2(input string) int64 {
	integers, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	count := int64(0)
	seen := make(map[mapKey]int64)

	for j := 0; j < len(integers); j++ {
		count += recurse(integers[j], 75, seen)
	}

	return count
}

func recurse(current int64, steps int, seen map[mapKey]int64) int64 {
	if steps == 0 {
		return 1
	}

	if count, ok := seen[mapKey{current, steps}]; ok {
		return count
	}

	if current == 0 {
		count := recurse(1, steps-1, seen)
		seen[mapKey{current, steps}] = count
		return count
	}

	digitCount := countDigits(current)
	if digitCount%2 == 1 {
		count := recurse(current*2024, steps-1, seen)
		seen[mapKey{current, steps}] = count
		return count
	}

	multiplier := getMultiplier(digitCount / 2)
	count := recurse(current/multiplier, steps-1, seen) +
		recurse(current%multiplier, steps-1, seen)
	seen[mapKey{current, steps}] = count
	return count
}

func countDigits(number int64) int {
	count := 0
	for number != 0 {
		number /= 10
		count++
	}

	return count
}

func getMultiplier(digits int) int64 {
	multiplier := int64(1)
	for i := 0; i < digits; i++ {
		multiplier *= 10
	}

	return multiplier
}

func parseInput(input string) ([]int64, error) {
	var data []int64
	for _, line := range lib.SplitLines(input) {
		if len(line) == 0 {
			continue
		}

		integers, err := lib.StringSliceToInt64(strings.Split(line, " "))
		if err != nil {
			return nil, err
		}
		data = append(data, integers...)
	}

	return data, nil
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
