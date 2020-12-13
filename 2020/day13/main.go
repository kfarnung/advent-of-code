package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

// greatestCommonDenominator calculates using the algorithm from:
// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm#Pseudocode
func greatestCommonDenominator(a, b int64) (int64, int64, int64) {
	var oldR, r int64 = a, b
	var oldS, s int64 = 1, 0
	var oldT, t int64 = 0, 1

	for r != 0 {
		quotient := oldR / r
		oldR, r = r, oldR-quotient*r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
	}

	return oldR, oldS, oldT
}

// chineseRemainderTheorem calculates using the algorithm from:
// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func chineseRemainderTheorem(remainders, moduluses []int64) (int64, error) {
	product := int64(1)
	for _, modulus := range moduluses {
		product *= modulus
	}

	var sum int64
	for i, modulus := range moduluses {
		quotient := product / modulus
		gcd, _, coefficient := greatestCommonDenominator(modulus, quotient)
		if gcd != 1 {
			return 0, fmt.Errorf("%d not coprime", modulus)
		}

		sum += remainders[i] * coefficient * quotient
	}

	return lib.ModInt64(sum, product), nil
}

func parseInput(lines []string) (int64, []int64, error) {
	if len(lines) != 2 {
		return 0, nil, errors.New("Invalid number of lines")
	}

	departureTarget, err := lib.ParseInt64(lines[0])
	if err != nil {
		return 0, nil, err
	}

	busIDs := strings.Split(lines[1], ",")
	busIDsInt := make([]int64, len(busIDs))

	for i, busID := range busIDs {
		var val int64

		if busID != "x" {
			val, err = lib.ParseInt64(busID)
			if err != nil {
				return 0, nil, err
			}
		}

		busIDsInt[i] = val
	}

	return departureTarget, busIDsInt, nil
}

func part1(lines []string) int64 {
	departureTarget, busIDs, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	var bestBus int64
	var bestBusDiff int64 = -1

	for _, busID := range busIDs {
		if busID == 0 {
			continue
		}

		mod := lib.ModInt64(departureTarget, busID)
		timeDiff := busID - mod

		if bestBusDiff == -1 || bestBusDiff > timeDiff {
			bestBus = busID
			bestBusDiff = timeDiff
		}
	}

	return bestBus * bestBusDiff
}

func part2(lines []string) int64 {
	_, busIDs, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	var remainders, moduluses []int64

	for i, busID := range busIDs {
		if busID == 0 {
			continue
		}

		remainders = append(remainders, lib.ModInt64(busID-int64(i), busID))
		moduluses = append(moduluses, busID)
	}

	value, err := chineseRemainderTheorem(remainders, moduluses)
	if err != nil {
		log.Fatal(err)
	}

	return value
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
