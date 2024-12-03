package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

var mulRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

func part1(input string) int64 {
	matches := mulRegex.FindAllStringSubmatch(input, -1)

	sum := int64(0)
	for _, match := range matches {
		if match[0] == "do()" || match[0] == "don't()" {
			continue
		}

		first, err := lib.ParseInt[int64](match[1])
		if err != nil {
			log.Fatal(err)
		}

		second, err := lib.ParseInt[int64](match[2])
		if err != nil {
			log.Fatal(err)
		}

		sum += first * second
	}

	return sum
}

func part2(input string) int64 {
	matches := mulRegex.FindAllStringSubmatch(input, -1)

	enabled := true
	sum := int64(0)
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {

			first, err := lib.ParseInt[int64](match[1])
			if err != nil {
				log.Fatal(err)
			}

			second, err := lib.ParseInt[int64](match[2])
			if err != nil {
				log.Fatal(err)
			}

			sum += first * second
		}
	}

	return sum
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
