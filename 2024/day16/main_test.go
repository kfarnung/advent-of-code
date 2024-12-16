package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input1 = strings.Join([]string{
	"###############",
	"#.......#....E#",
	"#.#.###.#.###.#",
	"#.....#.#...#.#",
	"#.###.#####.#.#",
	"#.#.#.......#.#",
	"#.#.#####.###.#",
	"#...........#.#",
	"###.#.#####.#.#",
	"#...#.....#.#.#",
	"#.#.#.###.#.#.#",
	"#.....#...#.#.#",
	"#.###.#.#.#.#.#",
	"#S..#.....#...#",
	"###############",
	"",
}, "\n")

var input2 = strings.Join([]string{
	"#################",
	"#...#...#...#..E#",
	"#.#.#.#.#.#.#.#.#",
	"#.#.#.#...#...#.#",
	"#.#.#.#.###.#.#.#",
	"#...#.#.#.....#.#",
	"#.#.#.#.#.#####.#",
	"#.#...#.#.#.....#",
	"#.#.#####.#.###.#",
	"#.#.#.......#...#",
	"#.#.###.#####.###",
	"#.#.#...#.....#.#",
	"#.#.#.#####.###.#",
	"#.#.#.........#.#",
	"#.#.#.#########.#",
	"#S#.............#",
	"#################",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(7036), part1(input1))
	assert.Equal(t, int64(11048), part1(input2))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(102504), part1(inputContent))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(45), part2(input1))
	assert.Equal(t, int64(64), part2(input2))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(535), part2(inputContent))
}
