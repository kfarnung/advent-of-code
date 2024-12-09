package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type file struct {
	id   int64
	size int
	free int
}

func part1(input string) int64 {
	layout := parseInput(input)

	left := 0
	right := len(layout) - 1
	for {
		if left >= right {
			break
		} else if layout[left] != -1 {
			left++
		} else if layout[right] == -1 {
			right--
		} else {
			layout[left] = layout[right]
			layout[right] = -1
		}
	}

	sum := int64(0)
	for i, id := range layout {
		if id == -1 {
			break
		}

		sum += int64(id * i)
	}

	return sum
}

func part2(input string) int64 {
	layout := parseInput2(input)

	for i := len(layout) - 1; i >= 0; {
		didMove := false
		for j := 0; j < i; j++ {
			if layout[j].free >= layout[i].size {
				toMove := layout[i]
				layout[i-1].free += toMove.size + toMove.free
				layout = slices.Delete(layout, i, i+1)

				toMove.free = layout[j].free - toMove.size
				layout[j].free = 0
				layout = slices.Insert(layout, j+1, toMove)
				didMove = true
				break
			}
		}

		if !didMove {
			i--
		}
	}

	block := int64(0)
	sum := int64(0)
	for _, file := range layout {
		for i := 0; i < file.size; i++ {
			sum += int64(block * file.id)
			block++
		}

		block += int64(file.free)
	}

	return sum
}

func parseInput(input string) []int {
	var layout []int
	for _, line := range lib.SplitLines(input) {
		if len(line) == 0 {
			continue
		}

		for i := 0; i < len(line); i += 2 {
			id := i / 2
			size := int(line[i] - '0')

			for j := 0; j < size; j++ {
				layout = append(layout, id)
			}

			if i+1 < len(line) {
				free := int(line[i+1] - '0')
				for j := 0; j < free; j++ {
					layout = append(layout, -1)
				}
			}
		}
	}

	return layout
}

func parseInput2(input string) []file {
	var files []file
	for _, line := range lib.SplitLines(input) {
		if len(line) == 0 {
			continue
		}

		for i := 0; i < len(line); i += 2 {
			file := file{
				id:   int64(i / 2),
				size: int(line[i] - '0'),
			}

			if i+1 < len(line) {
				file.free = int(line[i+1] - '0')
			}

			files = append(files, file)
		}
	}

	return files
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
