package main

import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type cupsGame struct {
	cups     *list.List
	current  *list.Element
	elements map[uint32]*list.Element
	minValue uint32
	maxValue uint32
}

func newCupsGame(line string, totalElements int) *cupsGame {
	cups := list.New()
	elements := make(map[uint32]*list.Element)
	minValue := ^uint32(0)
	maxValue := uint32(0)

	for _, char := range line {
		value := uint32(char) - uint32('0')
		cups.PushBack(value)
		elements[value] = cups.Back()

		if value > maxValue {
			maxValue = value
		}

		if value < minValue {
			minValue = value
		}
	}

	for i := cups.Len(); i < totalElements; i++ {
		maxValue++
		cups.PushBack(maxValue)
		elements[maxValue] = cups.Back()
	}

	return &cupsGame{
		cups:     cups,
		current:  cups.Front(),
		elements: elements,
		minValue: minValue,
		maxValue: maxValue,
	}
}

func (c *cupsGame) play(moveCount int) {
	for i := 0; i < moveCount; i++ {
		c.move()
	}
}

func (c *cupsGame) move() {
	currentValue := c.current.Value.(uint32)
	cupsToMove := c.nextCups(c.current, 3)

	nextValue := c.findNextValid(currentValue, cupsToMove)
	nextElement := c.elements[nextValue]
	for i := len(cupsToMove) - 1; i >= 0; i-- {
		c.cups.MoveAfter(cupsToMove[i], nextElement)
	}

	c.current = c.nextCup(c.current)
}

func (c cupsGame) nextCup(current *list.Element) *list.Element {
	if next := current.Next(); next != nil {
		return next
	}

	return c.cups.Front()
}

func (c cupsGame) nextCups(current *list.Element, count int) []*list.Element {
	cups := make([]*list.Element, count)
	for i := 0; i < count; i++ {
		current = c.nextCup(current)
		cups[i] = current
	}

	return cups
}

func (c cupsGame) decrement(current uint32) uint32 {
	if current == c.minValue {
		return c.maxValue
	}

	return current - 1
}

func (c cupsGame) findNextValid(current uint32, pickedUp []*list.Element) uint32 {
OUTERLOOP:
	for {
		current = c.decrement(current)

		for _, item := range pickedUp {
			if item.Value.(uint32) == current {
				continue OUTERLOOP
			}
		}

		break
	}

	return current
}

func (c cupsGame) getProductOfAdjacentCups(cup uint32) uint64 {
	current := c.elements[cup]
	product := uint64(1)

	for i := 0; i < 2; i++ {
		current = c.nextCup(current)
		product *= uint64(current.Value.(uint32))
	}

	return product
}

func (c cupsGame) String() string {
	var sb strings.Builder

	current := c.elements[1]
	for i := 0; i < c.cups.Len()-1; i++ {
		current = c.nextCup(current)
		fmt.Fprintf(&sb, "%d", current.Value.(uint32))
	}

	return sb.String()
}

func parseInput(lines []string, totalElements int) (*cupsGame, error) {
	if len(lines) != 1 {
		return nil, errors.New("Expected exactly one line")
	}

	return newCupsGame(lines[0], totalElements), nil
}

func part1(lines []string) string {
	game, err := parseInput(lines, 0)
	if err != nil {
		log.Fatal(err)
	}

	game.play(100)
	return game.String()
}

func part2(lines []string) uint64 {
	game, err := parseInput(lines, 1000000)
	if err != nil {
		log.Fatal(err)
	}

	game.play(10000000)
	return game.getProductOfAdjacentCups(1)
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %s\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
