package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input = strings.Join([]string{
	"2333133121414131402",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(1928), part1(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(6200294120911), part1(inputContent))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(2858), part2(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(6227018762750), part2(inputContent))
}
