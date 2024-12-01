package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	lines := []string{
		"35",
		"20",
		"15",
		"25",
		"47",
		"40",
		"62",
		"55",
		"65",
		"95",
		"102",
		"117",
		"150",
		"182",
		"127",
		"219",
		"299",
		"277",
		"309",
		"576",
	}

	search := part1(lines, 5)
	assert.Equal(t, int64(127), search)
	assert.Equal(t, int64(62), part2(lines, search))
}

func TestInput(t *testing.T) {
	name := lib.GetInputFilePath()
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	search := part1(lines, 25)
	assert.Equal(t, int64(2089807806), search)
	assert.Equal(t, int64(245848639), part2(lines, search))
}
