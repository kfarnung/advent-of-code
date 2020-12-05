package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

type inputOutput struct {
	input  string
	output int
}

func TestPart1(t *testing.T) {
	testCases := []inputOutput{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.output, calculateSeatID(testCase.input))
	}
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, 864, part1(lines))
	assert.Equal(t, 739, part2(lines))
}
