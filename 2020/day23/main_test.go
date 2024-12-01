package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"389125467",
	}

	assert.Equal(t, "67384529", part1(lines))
	assert.Equal(t, uint64(149245887792), part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetInputFilePath()
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, "28946753", part1(lines))
	assert.Equal(t, uint64(519044017360), part2(lines))
}
