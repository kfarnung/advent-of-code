package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func containsSum(nums []int64, index int, lookback int) bool {
	sum := nums[index]
	for i := index - lookback; i < index; i++ {
		for j := i + 1; j < index; j++ {
			if sum == nums[i]+nums[j] {
				return true
			}
		}
	}

	return false
}

func minMax(nums []int64, start int, end int) (int64, int64) {
	min := nums[start]
	max := nums[start]

	for i := start + 1; i < end; i++ {
		if nums[i] < min {
			min = nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return min, max
}

func parseInput(lines []string) ([]int64, error) {
	nums := make([]int64, len(lines))

	for i, line := range lines {
		num, err := lib.ParseInt64(line)
		if err != nil {
			return nil, err
		}

		nums[i] = num
	}

	return nums, nil
}

func part1(lines []string, lookback int) int64 {
	nums, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	for i := range nums {
		if i < lookback {
			continue
		}

		if !containsSum(nums, i, lookback) {
			return nums[i]
		}
	}

	return 0
}

func part2(lines []string, search int64) int64 {
	nums, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	var sum int64
	start := 0
	end := 0

	for sum != search {
		for sum < search {
			sum += nums[end]
			end++
		}

		for sum > search {
			sum -= nums[start]
			start++
		}
	}

	min, max := minMax(nums, start, end)
	return min + max
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	search := part1(lines, 25)
	fmt.Printf("Part 1: %d\n", search)
	fmt.Printf("Part 2: %d\n", part2(lines, search))
}
