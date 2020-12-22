package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}
	assert.Equal(t, int64(71), part1(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, int64(20013), part1(lines))
	assert.Equal(t, int64(5977293343129), part2(lines))
}
