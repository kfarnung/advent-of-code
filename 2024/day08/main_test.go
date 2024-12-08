package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input = strings.Join([]string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(14), part1(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(327), part1(inputContent))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(34), part2(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(1233), part2(inputContent))
}
