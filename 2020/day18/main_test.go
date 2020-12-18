package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input         string
	expectedPart1 int64
	expectedPart2 int64
}

func TestExample(t *testing.T) {
	testCases := []testCase{
		{"1 + 2 * 3 + 4 * 5 + 6", 71, 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51, 51},
		{"2 * 3 + (4 * 5)", 26, 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437, 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240, 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632, 23340},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expectedPart1, doMath(testCase.input, isHigherPrecedencePart1))
		assert.Equal(t, testCase.expectedPart2, doMath(testCase.input, isHigherPrecedencePart2))
	}
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, int64(53660285675207), part1(lines))
	assert.Equal(t, int64(141993988282687), part2(lines))
}
