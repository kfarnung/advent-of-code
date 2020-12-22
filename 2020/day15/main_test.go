package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"0,3,6",
	}

	assert.Equal(t, int32(436), part1(lines))
	assert.Equal(t, int32(175594), part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, int32(517), part1(lines))
	assert.Equal(t, int32(1047739), part2(lines))
}
