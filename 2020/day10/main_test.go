package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"16",
		"10",
		"15",
		"5",
		"1",
		"11",
		"7",
		"19",
		"6",
		"12",
		"4",
	}

	nums, err := parseInput(lines)
	assert.NoError(t, err)

	assert.Equal(t, 35, part1(nums))
	assert.Equal(t, int64(8), part2(nums))
}

func TestExample2(t *testing.T) {
	lines := []string{
		"28",
		"33",
		"18",
		"42",
		"31",
		"14",
		"46",
		"20",
		"48",
		"47",
		"24",
		"23",
		"49",
		"45",
		"19",
		"38",
		"39",
		"11",
		"1",
		"32",
		"25",
		"35",
		"8",
		"17",
		"7",
		"9",
		"4",
		"2",
		"34",
		"10",
		"3",
	}

	nums, err := parseInput(lines)
	assert.NoError(t, err)

	assert.Equal(t, 220, part1(nums))
	assert.Equal(t, int64(19208), part2(nums))
}

func TestInput(t *testing.T) {
	name := lib.GetInputFilePath()
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	nums, err := parseInput(lines)
	assert.NoError(t, err)

	assert.Equal(t, 2482, part1(nums))
	assert.Equal(t, int64(96717311574016), part2(nums))
}
