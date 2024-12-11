package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input = strings.Join([]string{
	"125 17",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(55312), part1(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(183620), part1(inputContent))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(65601038650482), part2(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(220377651399268), part2(inputContent))
}
