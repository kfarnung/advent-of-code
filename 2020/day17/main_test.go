package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		".#.",
		"..#",
		"###",
	}
	assert.Equal(t, 112, part1(lines))
	assert.Equal(t, 848, part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetInputFilePath()
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, 375, part1(lines))
	assert.Equal(t, 2192, part2(lines))
}
