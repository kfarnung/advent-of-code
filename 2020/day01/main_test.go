package main

import (
	"testing"

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
