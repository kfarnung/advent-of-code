package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

var pattern = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func loadPattern(pattern []string) map[lib.Point2D]rune {
	patternMap := make(map[lib.Point2D]rune)
	for i, line := range pattern {
		for j, char := range line {
			if char != ' ' {
				patternMap[lib.NewPoint2D(i, j)] = char
			}
		}
	}

	return patternMap
}

func runesToBinary(runes []rune, one rune) int {
	value := 0
	for _, char := range runes {
		value <<= 1
		if char == one {
			value |= 1
		}
	}

	return value
}

func reverseRuneSlice(runes []rune) []rune {
	output := make([]rune, len(runes))
	for i, value := range runes {
		output[len(runes)-1-i] = value
	}

	return output
}

func appendIfNotContainsInt(slice []int, value int) []int {
	for _, candidate := range slice {
		if candidate == value {
			return slice
		}
	}

	return append(slice, value)
}

type tileGrid [][]rune

func (t tileGrid) GetGridRow(index int) []rune {
	row := make([]rune, len(t[index]))
	for i, char := range t[index] {
		row[i] = char
	}

	return row
}

func (t tileGrid) GetGridColumn(index int) []rune {
	column := make([]rune, len(t))
	for i, row := range t {
		column[i] = row[index]
	}

	return column
}

func (t tileGrid) GetTopChecksum() int {
	return runesToBinary(t.GetGridRow(0), '#')
}

func (t tileGrid) GetBottomChecksum() int {
	return runesToBinary(t.GetGridRow(len(t)-1), '#')
}

func (t tileGrid) GetLeftChecksum() int {
	return runesToBinary(t.GetGridColumn(0), '#')
}

func (t tileGrid) GetRightChecksum() int {
	return runesToBinary(t.GetGridColumn(len(t)-1), '#')
}

func (t tileGrid) GetEdgeChecksums() []int {
	var rowColumns [][]rune
	rowColumns = append(rowColumns, t.GetGridRow(0))
	rowColumns = append(rowColumns, t.GetGridRow(len(t)-1))
	rowColumns = append(rowColumns, t.GetGridColumn(0))
	rowColumns = append(rowColumns, t.GetGridColumn(len(t[0])-1))

	var checksums []int
	for _, rowColumn := range rowColumns {
		checksums = append(checksums, runesToBinary(rowColumn, '#'))
		reversed := reverseRuneSlice(rowColumn)
		checksums = append(checksums, runesToBinary(reversed, '#'))
	}

	return checksums
}

func (t *tileGrid) FlipHorizontal() {
	for i, row := range *t {
		(*t)[i] = reverseRuneSlice(row)
	}
}

func (t *tileGrid) FlipVertical() {
	output := make([][]rune, len(*t))
	for i, row := range *t {
		output[len(*t)-1-i] = row
	}

	*t = output
}

func (t *tileGrid) RotateClockwise() {
	output := make([][]rune, len(*t))

	for i := 0; i < len(*t); i++ {
		output[i] = make([]rune, len(*t))

		for j := 0; j < len(*t); j++ {
			output[i][j] = (*t)[len(*t)-1-j][i]
		}
	}

	*t = output
}

func (t tileGrid) CountPattern(pattern map[lib.Point2D]rune) int {
	fmt.Print(t)
	count := 0

OUTERLOOP:
	for i := 0; i < len(t); i++ {
	INNERLOOP:
		for j := 0; j < len(t); j++ {
			found := true
			for point, value := range pattern {
				x := point.X + i
				y := point.Y + j

				if x >= len(t) {
					break OUTERLOOP
				} else if y >= len(t[x]) {
					break INNERLOOP
				} else if t[x][y] != value {
					found = false
					break
				}
			}

			if found {
				count++
			}
		}
	}

	return count
}

func (t tileGrid) CountSymbol(symbol rune) int {
	count := 0
	for _, row := range t {
		for _, column := range row {
			if column == symbol {
				count++
			}
		}
	}

	return count
}

func (t tileGrid) String() string {
	var sb strings.Builder
	for _, row := range t {
		for _, column := range row {
			sb.WriteRune(column)
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func parseInput(lines []string) (map[int]tileGrid, error) {
	headerRegex := regexp.MustCompile(`^Tile (\d+):`)
	tileMap := make(map[int]tileGrid)
	currentID := -1
	for _, line := range lines {
		if match := headerRegex.FindStringSubmatch(line); match != nil {
			tileID, err := lib.ParseInt(match[1])
			if err != nil {
				return nil, err
			}

			currentID = tileID
		} else if len(line) > 0 {
			var row []rune
			for _, char := range line {
				row = append(row, char)
			}

			tileMap[currentID] = append(tileMap[currentID], row)
		}
	}

	return tileMap, nil
}

func calculateChecksums(tiles map[int]tileGrid) map[int][]int {
	checksumMap := make(map[int][]int)
	for tileID, grid := range tiles {
		checksums := grid.GetEdgeChecksums()
		for _, checksum := range checksums {
			checksumMap[checksum] = append(checksumMap[checksum], tileID)
		}
	}

	return checksumMap
}

func calculateAdjacency(checksumMap map[int][]int) map[int][]int {
	adjacencyMap := make(map[int][]int)
	for _, matches := range checksumMap {
		if len(matches) == 0 || len(matches) > 2 {
			panic("Something smells fishy")
		}

		if len(matches) == 2 {
			tile1 := matches[0]
			tile2 := matches[1]

			adjacencyMap[tile1] = appendIfNotContainsInt(adjacencyMap[tile1], tile2)
			adjacencyMap[tile2] = appendIfNotContainsInt(adjacencyMap[tile2], tile1)
		}
	}

	return adjacencyMap
}

func calculateTileLayout(tiles map[int]tileGrid, checksumMap map[int][]int, adjacencyMap map[int][]int) (map[int]tileGrid, [][]int) {
	cornerTileID := 0
	for key, value := range adjacencyMap {
		if len(value) == 2 {
			cornerTileID = key
			break
		}
	}

	var tileLayout [][]int

OUTERLOOP:
	for i := 0; ; i++ {
		var tileLayoutRow []int

		for j := 0; ; j++ {
			if j == 0 {
				if i == 0 {
					tileLayoutRow = append(tileLayoutRow, cornerTileID)
					tile := tiles[cornerTileID]

					for {
						if val := checksumMap[tile.GetRightChecksum()]; len(val) == 2 {
							break
						}

						tile.RotateClockwise()
					}

					if val := checksumMap[tile.GetTopChecksum()]; len(val) == 2 {
						tile.FlipVertical()
					}

					tiles[cornerTileID] = tile
				} else {
					previousTileID := tileLayout[i-1][j]
					targetCheckSum := tiles[previousTileID].GetBottomChecksum()
					checksumTargets := checksumMap[targetCheckSum]
					if len(checksumTargets) != 2 {
						break OUTERLOOP
					}

					otherTileID := -1
					for _, tileID := range checksumTargets {
						if tileID != previousTileID {
							otherTileID = tileID
							break
						}
					}

					if otherTileID == -1 {
						panic("Failed to find other tile")
					}

					tileLayoutRow = append(tileLayoutRow, otherTileID)
					tile := tiles[otherTileID]
					for i := 0; ; i++ {
						if targetCheckSum == tile.GetTopChecksum() {
							break
						} else if i%4 == 0 {
							tile.FlipVertical()
						} else {
							tile.RotateClockwise()
						}
					}

					tiles[otherTileID] = tile
				}
			} else {
				previousTileID := tileLayoutRow[j-1]
				targetCheckSum := tiles[previousTileID].GetRightChecksum()
				checksumTargets := checksumMap[targetCheckSum]
				if len(checksumTargets) != 2 {
					break
				}

				otherTileID := -1
				for _, tileID := range checksumTargets {
					if tileID != previousTileID {
						otherTileID = tileID
						break
					}
				}

				if otherTileID == -1 {
					panic("Failed to find other tile")
				}

				tileLayoutRow = append(tileLayoutRow, otherTileID)
				tile := tiles[otherTileID]
				for i := 0; ; i++ {
					if targetCheckSum == tile.GetLeftChecksum() {
						break
					} else if i%4 == 0 {
						tile.FlipVertical()
					} else {
						tile.RotateClockwise()
					}
				}

				tiles[otherTileID] = tile
			}
		}

		tileLayout = append(tileLayout, tileLayoutRow)
	}

	return tiles, tileLayout
}

func part1(lines []string) int64 {
	tiles, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	checksumMap := calculateChecksums(tiles)
	adjacencyMap := calculateAdjacency(checksumMap)

	total := int64(1)
	for key, value := range adjacencyMap {
		if len(value) == 2 {
			total *= int64(key)
		}
	}

	return total
}

func part2(lines []string) int {
	tiles, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	checksumMap := calculateChecksums(tiles)
	adjacencyMap := calculateAdjacency(checksumMap)
	updatedTiles, tileLayout := calculateTileLayout(tiles, checksumMap, adjacencyMap)

	var largeGrid tileGrid
	for _, row := range tileLayout {
		for i := 1; i < len(updatedTiles[row[0]])-1; i++ {
			var largeGridRow []rune
			for _, column := range row {
				tile := updatedTiles[column]

				for j := 1; j < len(tile[i])-1; j++ {
					largeGridRow = append(largeGridRow, tile[i][j])
				}
			}

			largeGrid = append(largeGrid, largeGridRow)
		}
	}

	patternMap := loadPattern(pattern)
	count := 0
	for i := 0; ; i++ {
		count = largeGrid.CountPattern(patternMap)
		if count > 0 {
			break
		} else if i%4 == 0 {
			largeGrid.FlipVertical()
		} else {
			largeGrid.RotateClockwise()
		}
	}

	// This assumes that the pattern never overlaps, apparently that worked for
	// my input.
	return largeGrid.CountSymbol('#') - (count * len(patternMap))
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
