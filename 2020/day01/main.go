package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func findSumImpl(candidates []int64, count int, target int64, depth int, sum int64, start int) []int64 {
	for i := start; i < len(candidates); i++ {
		if depth < count-1 {
			result := findSumImpl(candidates, count, target, depth+1, sum+candidates[i], i+1)
			if result != nil {
				return append(result, candidates[i])
			}
		} else if sum+candidates[i] == target {
			return []int64{candidates[i]}
		}
	}

	return nil
}

func findSum(candidates []int64, count int, target int64) []int64 {
	return findSumImpl(candidates, count, target, 0, 0, 0)
}

func multSlice(vals []int64) int64 {
	if len(vals) == 0 {
		return 0
	}

	mult := int64(1)

	for _, val := range vals {
		mult *= val
	}

	return mult
}

func productOfEntries(values []int64, count int) int64 {
	vals := findSum(values, count, 2020)
	if vals == nil {
		log.Fatal("Failed to find sum")
	}

	return multSlice(vals)
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	values, err := lib.StringSliceToInt64(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", productOfEntries(values, 2))
	fmt.Printf("Part 2: %d\n", productOfEntries(values, 3))
}
