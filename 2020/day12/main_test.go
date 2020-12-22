package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	assert.Equal(t, int64(25), part1(lines))
	assert.Equal(t, int64(286), part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, int64(1152), part1(lines))
	assert.Equal(t, int64(58637), part2(lines))
}
