package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type constraint struct {
	first  int64
	second int64
}

func part1(input string) int64 {
	constraints, err := parseConstraints(input)
	if err != nil {
		log.Fatal(err)
	}

	pageLists, err := parsePageList(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := int64(0)
	for _, pages := range pageLists {
		meetsAllConstraints := true
		for _, constraint := range constraints {
			firstIndex := slices.Index(pages, constraint.first)
			secondIndex := slices.Index(pages, constraint.second)

			if firstIndex == -1 || secondIndex == -1 {
				continue
			}

			if secondIndex < firstIndex {
				meetsAllConstraints = false
				break
			}
		}

		if meetsAllConstraints {
			sum += pages[len(pages)/2]
		}
	}

	return sum
}

func part2(input string) int64 {
	constraints, err := parseConstraints(input)
	if err != nil {
		log.Fatal(err)
	}

	pageLists, err := parsePageList(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := int64(0)
	for _, pages := range pageLists {
		meetsAllConstraints := true
		didSwap := true
		for didSwap {
			didSwap = false
			for _, constraint := range constraints {
				firstIndex := slices.Index(pages, constraint.first)
				secondIndex := slices.Index(pages, constraint.second)

				if firstIndex == -1 || secondIndex == -1 {
					continue
				}

				if secondIndex < firstIndex {
					meetsAllConstraints = false

					temp := pages[firstIndex]
					pages[firstIndex] = pages[secondIndex]
					pages[secondIndex] = temp

					didSwap = true
				}
			}
		}

		if !meetsAllConstraints {
			sum += pages[len(pages)/2]
		}
	}

	return sum
}

func parseConstraints(input string) ([]constraint, error) {
	var constraints []constraint
	for _, line := range lib.SplitLines(input) {
		if line == "" {
			break
		}

		split := strings.Split(line, "|")
		first, err := lib.ParseInt[int64](split[0])
		if err != nil {
			return nil, err
		}

		second, err := lib.ParseInt[int64](split[1])
		if err != nil {
			return nil, err
		}

		constraints = append(constraints, constraint{first, second})
	}

	return constraints, nil
}

func parsePageList(input string) ([][]int64, error) {
	var pages [][]int64
	foundBreak := false
	for _, line := range lib.SplitLines(input) {
		if line == "" {
			foundBreak = true
			continue
		} else if !foundBreak {
			continue
		}

		split := strings.Split(line, ",")
		var page []int64
		for _, s := range split {
			val, err := lib.ParseInt[int64](s)
			if err != nil {
				return nil, err
			}

			page = append(page, val)
		}

		pages = append(pages, page)
	}

	return pages, nil
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
