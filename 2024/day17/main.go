package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2024/lib"
)

type computer struct {
	registers registers
	program   []int32
	output    func(value int32) bool
}

type registers struct {
	a, b, c int64
	ip      int
}

func (c computer) LiteralOperand() int64 {
	return int64(c.program[c.registers.ip+1])
}

func (c computer) ComboOperand() int64 {
	operand := c.LiteralOperand()
	switch operand {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return operand
	case 4:
		return c.registers.a
	case 5:
		return c.registers.b
	case 6:
		return c.registers.c
	}

	log.Fatalf("Invalid operand: %d", operand)
	return 0
}

func (c computer) Divide() int64 {
	return int64(int64(c.registers.a) / intPow(2, int64(c.ComboOperand())))
}

func (c *computer) Run() error {
	seenStates := make(map[registers]bool)

	for c.registers.ip < len(c.program) {
		if seenStates[c.registers] {
			return fmt.Errorf("infinite loop detected")
		}

		seenStates[c.registers] = true

		switch c.program[c.registers.ip] {
		case 0:
			// adv
			c.registers.a = c.Divide()
			c.registers.ip += 2

		case 1:
			// bxl
			c.registers.b = c.registers.b ^ c.LiteralOperand()
			c.registers.ip += 2

		case 2:
			// bst
			c.registers.b = c.ComboOperand() % 8
			c.registers.ip += 2

		case 3:
			// jnz
			if c.registers.a != 0 {
				c.registers.ip = int(c.LiteralOperand())
			} else {
				c.registers.ip += 2
			}

		case 4:
			// bxc
			c.registers.b = c.registers.b ^ c.registers.c
			c.registers.ip += 2

		case 5:
			// out
			if !c.output(int32(c.ComboOperand() % 8)) {
				return fmt.Errorf("program aborted")
			}

			c.registers.ip += 2

		case 6:
			// bdv
			c.registers.b = c.Divide()
			c.registers.ip += 2

		case 7:
			// cdv
			c.registers.c = c.Divide()
			c.registers.ip += 2
		}
	}

	return nil
}

func (c *computer) Reset() {
	c.registers.a = 0
	c.registers.b = 0
	c.registers.c = 0
	c.registers.ip = 0
}

func part1(input string) string {
	computer := parseInput(input)
	var output []int32
	computer.output = func(value int32) bool {
		output = append(output, value)
		return true
	}

	err := computer.Run()
	if err != nil {
		log.Fatal(err)
	}

	outputStrings := make([]string, len(output))
	for i, val := range output {
		outputStrings[i] = fmt.Sprintf("%d", val)
	}

	return strings.Join(outputStrings, ",")
}

func part2(input string) int64 {
	computer := parseInput(input)
	value, err := solve(0, computer.program)
	if err != nil {
		log.Fatal(err)
	}

	return value
}

func solve(current int64, program []int32) (int64, error) {
	target := program[len(program)-1]
	for i := int64(0); i < 8; i++ {
		value := doIteration(current + i)
		if value == int64(target) {
			if len(program) == 1 {
				return current + i, nil
			}

			next, err := solve((current+i)*8, program[:len(program)-1])
			if err == nil {
				return next, nil
			}
		}
	}

	return 0, fmt.Errorf("no solution found")
}

func doIteration(a int64) int64 {
	b, c := int64(0), int64(0)
	b = a % 8
	b = b ^ 2
	c = a / intPow(2, b)
	b = b ^ 3
	b = b ^ c
	return b % 8
}

func intPow(base, exp int64) int64 {
	if exp == 0 {
		return 1
	}

	value := base
	for i := int64(1); i < exp; i++ {
		value *= base
	}

	return value
}

func parseInput(input string) computer {
	lines := lib.SplitLines(input)

	a, err := lib.ParseInt[int64](strings.Split(lines[0], ": ")[1])
	if err != nil {
		log.Fatal(err)
	}

	b, err := lib.ParseInt[int64](strings.Split(lines[1], ": ")[1])
	if err != nil {
		log.Fatal(err)
	}

	c, err := lib.ParseInt[int64](strings.Split(lines[2], ": ")[1])
	if err != nil {
		log.Fatal(err)
	}

	programString := strings.Split(lines[4], ": ")[1]
	program, err := lib.StringSliceToInt32(strings.Split(programString, ","))
	if err != nil {
		log.Fatal(err)
	}

	return computer{registers: registers{a: a, b: b, c: c}, program: program}
}

func main() {
	name := os.Args[1]
	content, err := lib.LoadFileContent(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %s\n", part1(content))
	fmt.Printf("Part 2: %d\n", part2(content))
}
