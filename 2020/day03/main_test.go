package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	grid, err := parseInput(name)
	assert.NoError(t, err)

	assert.Equal(t, 214, part1(grid))
	assert.Equal(t, 8336352024, part2(grid))
}
