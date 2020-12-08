package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type instruction struct {
	operation string
	argument  int
}

type machine struct {
	instructions       []instruction
	instructionPointer int
	accumulator        int
	visited            []bool
}

func newMachine(instructions []instruction) machine {
	return machine{
		instructions:       instructions,
		instructionPointer: 0,
		accumulator:        0,
		visited:            make([]bool, len(instructions)),
	}
}

func (m *machine) step() error {
	if m.visited[m.instructionPointer] {
		return errors.New("Infinite loop detected")
	}

	m.visited[m.instructionPointer] = true
	instruction := m.instructions[m.instructionPointer]
	switch instruction.operation {
	case "acc":
		m.accumulator += instruction.argument
		m.instructionPointer++

	case "jmp":
		m.instructionPointer += instruction.argument

	case "nop":
		m.instructionPointer++

	default:
		return errors.New("Unknown operation encountered")
	}

	return nil
}

func (m *machine) run() error {
	for m.instructionPointer < len(m.instructions) && m.instructionPointer >= 0 {
		if err := m.step(); err != nil {
			return err
		}
	}

	if m.instructionPointer != len(m.instructions) {
		return errors.New("Instruction pointer out of bounds")
	}

	return nil
}

func parseInstruction(line string) (instruction, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return instruction{}, errors.New("Unable to parse instruction")
	}

	argument, err := lib.ParseInt(parts[1])
	if err != nil {
		return instruction{}, err
	}

	return instruction{
		operation: parts[0],
		argument:  argument,
	}, nil
}

func parseInput(lines []string) ([]instruction, error) {
	var instructions []instruction

	for _, line := range lines {
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func part1(instructions []instruction) int {
	machine := newMachine(instructions)

	for {
		if err := machine.step(); err != nil {
			break
		}
	}

	return machine.accumulator
}

func part2(instructions []instruction) int {
	// Make a copy of the slice to avoid tainting the incoming slice.
	modifiedInstructions := make([]instruction, len(instructions))
	copy(modifiedInstructions, instructions)

	for i := range modifiedInstructions {
		originalOperation := modifiedInstructions[i].operation
		if originalOperation == "jmp" {
			modifiedInstructions[i].operation = "nop"
		} else if originalOperation == "nop" {
			modifiedInstructions[i].operation = "jmp"
		} else {
			continue
		}

		machine := newMachine(modifiedInstructions)
		if err := machine.run(); err == nil {
			return machine.accumulator
		}

		// Restore the instruction
		modifiedInstructions[i].operation = originalOperation
	}

	return 0
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	instructions, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(instructions))
	fmt.Printf("Part 2: %d\n", part2(instructions))
}
