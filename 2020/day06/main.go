package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type customsDeclaration struct {
	count          int
	questionCounts map[rune]int
}

func newCustomsDeclaration() customsDeclaration {
	return customsDeclaration{
		questionCounts: make(map[rune]int),
	}
}

func (c *customsDeclaration) processLine(line string) {
	c.count++
	for _, char := range line {
		c.questionCounts[char]++
	}
}

func (c *customsDeclaration) countUnion() int {
	return len(c.questionCounts)
}

func (c *customsDeclaration) countIntersection() int {
	count := 0
	for _, questionCount := range c.questionCounts {
		if questionCount == c.count {
			count++
		}
	}

	return count
}

func part1(lines []string) int {
	count := 0
	current := newCustomsDeclaration()
	for _, line := range lines {
		if len(line) > 0 {
			current.processLine(line)
		} else {
			count += current.countUnion()
			current = newCustomsDeclaration()
		}
	}

	count += current.countUnion()
	return count
}

func part2(lines []string) int {
	count := 0
	current := newCustomsDeclaration()
	for _, line := range lines {
		if len(line) > 0 {
			current.processLine(line)
		} else {
			count += current.countIntersection()
			current = newCustomsDeclaration()
		}
	}

	count += current.countIntersection()
	return count
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
