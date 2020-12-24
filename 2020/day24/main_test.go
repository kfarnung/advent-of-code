package main

import (
	"testing"

	"github.com/kfarnung/advent-of-code/2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	lines := []string{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew",
	}

	assert.Equal(t, 10, part1(lines))
	assert.Equal(t, 2208, part2(lines))
}

func TestInput(t *testing.T) {
	name := lib.GetTestFilePath("input.txt")
	lines, err := lib.LoadFileLines(name)
	assert.NoError(t, err)

	assert.Equal(t, 354, part1(lines))
	assert.Equal(t, 3608, part2(lines))
}
