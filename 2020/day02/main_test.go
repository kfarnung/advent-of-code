package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	name := lib.GetInputFilePath()
	parsedLines, err := parseInput(name)
	assert.NoError(t, err)

	assert.Equal(t, 528, part1(parsedLines))
	assert.Equal(t, 497, part2(parsedLines))
}
