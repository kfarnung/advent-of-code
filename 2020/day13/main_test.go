package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}

	assert.Equal(t, int64(295), part1(lines))
	assert.Equal(t, int64(1068781), part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, int64(171), part1(lines))
	assert.Equal(t, int64(539746751134958), part2(lines))
}
