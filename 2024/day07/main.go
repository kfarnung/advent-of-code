package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type equation struct {
	testValue int64
	operands  []int64
}

func part1(input string) int64 {
	equations, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := int64(0)
	for _, equation := range equations {
		if isValidEquation(equation.operands[0], equation.operands[1:], equation.testValue, false) {
			sum += equation.testValue
		}
	}

	return sum
}

func part2(input string) int64 {
	equations, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := int64(0)
	for _, equation := range equations {
		if isValidEquation(equation.operands[0], equation.operands[1:], equation.testValue, true) {
			sum += equation.testValue
		}
	}

	return sum
}

func isValidEquation(current int64, operands []int64, target int64, allowConcat bool) bool {
	if len(operands) == 0 {
		return current == target
	}

	operand := operands[0]
	operands = operands[1:]

	if isValidEquation(current+operand, operands, target, allowConcat) {
		return true
	}

	if isValidEquation(current*operand, operands, target, allowConcat) {
		return true
	}

	if allowConcat {
		if isValidEquation(current*getMultiplier(operand)+operand, operands, target, allowConcat) {
			return true
		}
	}

	return false
}

func getMultiplier(num int64) int64 {
	if num == 0 {
		return 1
	}

	multiplier := int64(1)
	for num > 0 {
		num /= 10
		multiplier *= 10
	}

	return multiplier
}

func parseInput(input string) ([]equation, error) {
	var equations []equation
	for _, line := range lib.SplitLines(input) {
		before, after, found := strings.Cut(line, ": ")
		if !found {
			continue
		}

		testValue, err := lib.ParseInt[int64](before)
		if err != nil {
			return nil, err
		}

		var operands []int64
		for _, operandValue := range strings.Split(after, " ") {
			operand, err := lib.ParseInt[int64](operandValue)
			if err != nil {
				return nil, err
			}

			operands = append(operands, operand)
		}

		equations = append(equations, equation{testValue: testValue, operands: operands})
	}

	return equations, nil
}

func main() {
	name := os.Args[1]
	content, err := lib.LoadFileContent(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(content))
	fmt.Printf("Part 2: %d\n", part2(content))
}
