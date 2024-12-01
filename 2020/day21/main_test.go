package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
		"trh fvjkl sbzzf mxmxvkd (contains dairy)",
		"sqjhc fvjkl (contains soy)",
		"sqjhc mxmxvkd sbzzf (contains fish)",
	}

	assert.Equal(t, 5, part1(lines))
	assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetInputFilePath()
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, 2724, part1(lines))
	assert.Equal(t, "xlxknk,cskbmx,cjdmk,bmhn,jrmr,tzxcmr,fmgxh,fxzh", part2(lines))
}
