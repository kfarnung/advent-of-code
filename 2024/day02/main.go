package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

func part1(input string) int64 {
	reports, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	count := int64(0)
	for _, report := range reports {
		if isSafeReport(report) {
			count++
		}

		slices.Reverse(report)

		if isSafeReport(report) {
			count++
		}
	}

	return count
}

func part2(input string) int64 {
	reports, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	count := int64(0)
	for _, report := range reports {
		if isSafeReportAllowOneSkip(report) {
			count++
		}
	}

	return count
}

func isSafeReportAllowOneSkip(report []int64) bool {
	for i := 0; i < len(report); i++ {
		backup := report[i]
		report = slices.Delete(report, i, i+1)
		if isSafeReport(report) {
			return true
		}

		slices.Reverse(report)
		if isSafeReport(report) {
			return true
		}

		slices.Reverse(report)
		report = slices.Insert(report, i, backup)
	}

	return false
}

func isSafeReport(report []int64) bool {
	for i := 1; i < len(report); i++ {
		if report[i] <= report[i-1] || report[i]-report[i-1] > 3 {
			return false
		}
	}

	return true
}

func parseInput(input string) ([][]int64, error) {
	var reports [][]int64

	for _, line := range lib.SplitLines(input) {
		if line == "" {
			continue
		}

		var levels []int64
		numbers := strings.Split(line, " ")
		for _, number := range numbers {
			num, err := lib.ParseInt[int64](number)
			if err != nil {
				return nil, err
			}

			levels = append(levels, num)
		}

		reports = append(reports, levels)
	}

	return reports, nil
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
