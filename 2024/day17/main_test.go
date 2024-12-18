package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input = strings.Join([]string{
	"Register A: 729",
	"Register B: 0",
	"Register C: 0",
	"",
	"Program: 0,1,5,4,3,0",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", part1(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "7,1,5,2,4,0,7,6,1", part1(inputContent))
}

func TestPart2(t *testing.T) {
	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(37222273957364), part2(inputContent))
}
