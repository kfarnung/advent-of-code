package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input = strings.Join([]string{
	"r, wr, b, g, bwu, rb, gb, br",
	"",
	"brwrr",
	"bggr",
	"gbbr",
	"rrbgbr",
	"ubwu",
	"bwurrg",
	"brgr",
	"bbrgwb",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(6), part1(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(236), part1(inputContent))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(16), part2(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(643685981770598), part2(inputContent))
}
