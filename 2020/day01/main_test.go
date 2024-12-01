package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestProductOfEntries(t *testing.T) {
	input := []int64{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	assert.Equal(t, int64(514579), productOfEntries(input, 2))
	assert.Equal(t, int64(241861950), productOfEntries(input, 3))
}

func TestInput(t *testing.T) {
	name := lib.GetInputFilePath()
	values, err := parseFile(name)
	assert.NoError(t, err)

	assert.Equal(t, int64(744475), productOfEntries(values, 2))
	assert.Equal(t, int64(70276940), productOfEntries(values, 3))
}
