package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input = strings.Join([]string{
	"5,4",
	"4,2",
	"4,5",
	"3,0",
	"2,1",
	"6,3",
	"2,4",
	"1,5",
	"0,6",
	"3,3",
	"2,6",
	"5,1",
	"1,2",
	"5,5",
	"2,5",
	"6,5",
	"1,4",
	"0,4",
	"6,4",
	"1,1",
	"6,1",
	"1,0",
	"0,5",
	"1,6",
	"2,0",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(22), part1(input, 6, 12))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(330), part1(inputContent, 70, 1024))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "6,1", part2(input, 6))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "10,38", part2(inputContent, 70))
}
