package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
		"",
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	}

	assert.Equal(t, int(306), part1(lines))
	assert.Equal(t, int(291), part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, int(32824), part1(lines))
	assert.Equal(t, int(36515), part2(lines))
}
