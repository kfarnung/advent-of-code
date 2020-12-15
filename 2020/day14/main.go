package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type instruction struct {
	opcode  string
	address int64
	value   string
}

func (i instruction) ValueAsInt64() (int64, error) {
	value, err := lib.ParseInt64(i.value)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (i instruction) ValueAsMask() (mask, error) {
	if i.opcode != "mask" {
		return mask{}, errors.New("Current opcode doesn't contain a mask")
	}

	return newMask(i.value), nil
}

type mask struct {
	zeros       int64
	ones        int64
	fluctuating []int
}

func newMask(text string) mask {
	var zeros, ones int64
	var fluctuating []int
	bitCount := len(text) - 1
	for i, value := range text {
		zeros <<= 1
		ones <<= 1

		if value == '1' {
			ones |= 1
		} else if value == '0' {
			zeros |= 1
		} else if value == 'X' {
			fluctuating = append(fluctuating, bitCount-i)
		} else {
			panic("Unexpected mask value")
		}
	}

	return mask{
		zeros:       zeros,
		ones:        ones,
		fluctuating: fluctuating,
	}
}

func (m mask) Apply(value int64) int64 {
	value |= m.ones
	value &= ^m.zeros

	return value
}

func (m mask) GetAddresses(address int64) []int64 {
	address |= m.ones
	return getAddressesImpl(address, m.fluctuating)
}

func getAddressesImpl(address int64, fluctuating []int) []int64 {
	if len(fluctuating) == 0 {
		return []int64{address}
	}

	var bitMask int64 = 1 << fluctuating[0]
	addressZero := address & ^bitMask
	addressOne := address | bitMask
	newFluctuating := fluctuating[1:]

	return append(getAddressesImpl(addressZero, newFluctuating), getAddressesImpl(addressOne, newFluctuating)...)
}

func parseInput(lines []string) ([]instruction, error) {
	inputRegex := regexp.MustCompile(`^(\w+)(?:\[(\d+)\])? = (.+)$`)
	var instructions []instruction

	for _, line := range lines {
		match := inputRegex.FindStringSubmatch(line)
		if match == nil {
			return nil, errors.New("Could not match line")
		}

		current := instruction{
			opcode: match[1],
			value:  match[3],
		}

		addressString := match[2]
		if len(addressString) > 0 {
			address, err := lib.ParseInt64(addressString)
			if err != nil {
				return nil, err
			}

			current.address = address
		}

		instructions = append(instructions, current)
	}

	return instructions, nil
}

func part1(lines []string) int64 {
	instructions, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	var m mask
	memory := make(map[int64]int64)
	for _, instruction := range instructions {
		switch instruction.opcode {
		case "mask":
			m, err = instruction.ValueAsMask()
			if err != nil {
				log.Fatal(err)
			}

		case "mem":
			value, err := instruction.ValueAsInt64()
			if err != nil {
				log.Fatal(err)
			}
			memory[instruction.address] = m.Apply(value)
		}
	}

	var sum int64
	for _, value := range memory {
		sum += value
	}

	return sum
}

func part2(lines []string) int64 {
	instructions, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	var m mask
	memory := make(map[int64]int64)
	for _, instruction := range instructions {
		switch instruction.opcode {
		case "mask":
			m, err = instruction.ValueAsMask()
			if err != nil {
				log.Fatal(err)
			}

		case "mem":
			value, err := instruction.ValueAsInt64()
			if err != nil {
				log.Fatal(err)
			}

			for _, address := range m.GetAddresses(instruction.address) {
				memory[address] = value
			}
		}
	}

	var sum int64
	for _, value := range memory {
		sum += value
	}

	return sum
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
