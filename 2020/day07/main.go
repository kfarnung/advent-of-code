package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type bagInfo struct {
	count int32
	color string
}

func countChildren(bags map[string][]bagInfo, memo map[string]int32, color string) int32 {
	if val, ok := memo[color]; ok {
		return val
	}

	var count int32
	children := bags[color]
	for _, child := range children {
		count += child.count + (child.count * countChildren(bags, memo, child.color))
	}

	memo[color] = count
	return count
}

func parseBag(line string) bagInfo {
	bagSplitter := regexp.MustCompile(`^(?:(\d+) )?(.+) bags?$`)
	match := bagSplitter.FindStringSubmatch(line)
	if match == nil {
		panic("Failed to parse bag")
	}

	var count int32
	if len(match[1]) > 0 {
		parsedCount, err := lib.ParseInt32(match[1])
		if err != nil {
			panic("Couldn't parse integer")
		}

		count = parsedCount
	}

	return bagInfo{
		count: count,
		color: match[2],
	}
}

func parseLine(line string) (string, []bagInfo) {
	lineSplitter := regexp.MustCompile(`^(.+) contain (.+)\.$`)
	matches := lineSplitter.FindStringSubmatch(line)
	if matches == nil {
		panic("Failed to parse line")
	}

	parentBag := parseBag(matches[1])
	childBagsText := strings.Split(matches[2], ", ")
	childBags := make([]bagInfo, len(childBagsText))
	for i, val := range childBagsText {
		childBags[i] = parseBag(val)
	}

	return parentBag.color, childBags
}

func parseInput(lines []string) map[string][]bagInfo {
	bags := make(map[string][]bagInfo)

	for _, line := range lines {
		color, children := parseLine(line)
		bags[color] = children
	}

	return bags
}

func part1(lines []string) int {
	bags := parseInput(lines)
	containedBy := make(map[string][]string)

	for color, childBags := range bags {
		for _, childBag := range childBags {
			containedBy[childBag.color] = append(containedBy[childBag.color], color)
		}
	}

	visited := make(map[string]bool)
	queue := list.New()
	queue.PushBack("shiny gold")

	for queue.Len() > 0 {
		item := queue.Front()
		queue.Remove(item)
		color := item.Value.(string)
		visited[color] = true

		for _, bag := range containedBy[color] {
			if !visited[bag] {
				queue.PushBack(bag)
			}
		}
	}

	// Because we also visited the "shiny gold" bag.
	return len(visited) - 1
}

func part2(lines []string) int32 {
	bags := parseInput(lines)
	memo := make(map[string]int32)
	return countChildren(bags, memo, "shiny gold")
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
