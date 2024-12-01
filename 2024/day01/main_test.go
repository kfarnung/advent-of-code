package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

var input string = strings.Join([]string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(11), part1(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(1603498), part1(inputContent))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(31), part2(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(25574739), part2(inputContent))
}
