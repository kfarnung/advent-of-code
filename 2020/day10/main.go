package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func countPaths(nums []int, memo map[int]int64, index int) int64 {
	if value, ok := memo[index]; ok {
		return value
	}

	if index == len(nums)-1 {
		return 1
	}

	var count int64
	current := nums[index]
	for i := index + 1; i < len(nums) && nums[i]-current <= 3; i++ {
		count += countPaths(nums, memo, i)
	}

	memo[index] = count
	return count
}

func parseInput(lines []string) ([]int, error) {
	nums := make([]int, len(lines))

	for i, line := range lines {
		num, err := lib.ParseInt32(line)
		if err != nil {
			return nil, err
		}

		nums[i] = int(num)
	}

	nums = append(nums, 0)
	sort.Ints(nums)
	highest := nums[len(nums)-1]
	nums = append(nums, highest+3)

	return nums, nil
}

func part1(nums []int) int {
	var lastNum int
	diffCounts := make([]int, 4)
	for _, num := range nums {
		diff := num - lastNum
		diffCounts[diff]++
		lastNum = num
	}

	return diffCounts[1] * diffCounts[3]
}

func part2(nums []int) int64 {
	memo := make(map[int]int64)
	return countPaths(nums, memo, 0)
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	nums, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(nums))
	fmt.Printf("Part 2: %d\n", part2(nums))
}
