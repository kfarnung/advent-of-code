package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

func part1(input string) int64 {
	towels, patterns := parseInput(input)
	count := int64(0)

	for _, pattern := range patterns {
		if isValidPattern(towels, pattern) {
			count++
		}
	}

	return count
}

func part2(input string) int64 {
	towels, patterns := parseInput(input)
	memo := make(map[string]int64)
	count := int64(0)

	for _, pattern := range patterns {
		count += countValidPatterns(towels, pattern, memo)
	}

	return count
}

func isValidPattern(towels []string, pattern string) bool {
	if pattern == "" {
		return true
	}

	for _, towel := range towels {
		if len(towel) > len(pattern) {
			continue
		}

		if strings.HasPrefix(pattern, towel) {
			if isValidPattern(towels, pattern[len(towel):]) {
				return true
			}
		}
	}

	return false
}

func countValidPatterns(towels []string, pattern string, memo map[string]int64) int64 {
	if pattern == "" {
		return 1
	}

	if count, ok := memo[pattern]; ok {
		return count
	}

	count := int64(0)

	for _, towel := range towels {
		if len(towel) > len(pattern) {
			continue
		}

		if strings.HasPrefix(pattern, towel) {
			nextPattern := pattern[len(towel):]
			tempCount := countValidPatterns(towels, nextPattern, memo)
			memo[nextPattern] = tempCount
			count += tempCount
		}
	}

	return count
}

func parseInput(input string) ([]string, []string) {
	var towels []string
	var patterns []string
	for _, line := range lib.SplitLines(input) {
		if line == "" {
			continue
		}

		if towels == nil {
			towels = strings.Split(line, ", ")
		} else {
			patterns = append(patterns, line)
		}
	}

	return towels, patterns
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
