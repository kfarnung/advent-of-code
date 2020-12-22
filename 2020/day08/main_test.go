package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	instructions, err := parseInput(lines)
	assert.NoError(t, err)

	assert.Equal(t, int32(5), part1(instructions))
	assert.Equal(t, int32(8), part2(instructions))
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	instructions, err := parseInput(lines)
	assert.NoError(t, err)

	assert.Equal(t, int32(2034), part1(instructions))
	assert.Equal(t, int32(672), part2(instructions))
}
