package main

import (
	"strings"
	"testing"

	"github.com/kfarnung/advent-of-code/2024/lib"
	"github.com/stretchr/testify/assert"
)

var input = strings.Join([]string{
	"47|53",
	"97|13",
	"97|61",
	"97|47",
	"75|29",
	"61|13",
	"75|53",
	"29|13",
	"97|29",
	"53|29",
	"61|53",
	"97|53",
	"61|29",
	"47|13",
	"75|47",
	"97|75",
	"47|61",
	"75|61",
	"47|29",
	"75|13",
	"53|13",
	"",
	"75,47,61,53,29",
	"97,61,53,29,13",
	"75,29,13",
	"75,97,47,61,53",
	"61,13,29",
	"97,13,75,29,47",
	"",
}, "\n")

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(143), part1(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(6041), part1(inputContent))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, int64(123), part2(input))

	inputContent, err := lib.GetInputContent()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(4884), part2(inputContent))
}
