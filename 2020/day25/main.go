package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

const publicKeySubject = 7
const valueLimit = 20201227

func doIteration(value int64, subjectValue int64) int64 {
	value *= subjectValue
	value %= valueLimit

	return value
}

func deriveEncryptionKey(subjectValue int64, loopSize int) int64 {
	value := int64(1)
	for i := 0; i < loopSize; i++ {
		value = doIteration(value, subjectValue)
	}

	return value
}

func findLoopSize(publicKey int64) int {
	value := int64(1)
	for i := 0; ; i++ {
		if value == publicKey {
			return i
		}

		value = doIteration(value, publicKeySubject)
	}
}

func parseInput(lines []string) (int64, int64, error) {
	cardPublicKey, err := lib.ParseInt64(lines[0])
	if err != nil {
		return 0, 0, err
	}

	doorPublicKey, err := lib.ParseInt64(lines[1])
	if err != nil {
		return 0, 0, err
	}

	return cardPublicKey, doorPublicKey, nil
}

func part1(lines []string) int64 {
	cardPublicKey, doorPublicKey, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	cardLoopSize := findLoopSize(cardPublicKey)
	doorLoopSize := findLoopSize(doorPublicKey)

	cardEncryptionKey := deriveEncryptionKey(doorPublicKey, cardLoopSize)
	doorEncryptionKey := deriveEncryptionKey(cardPublicKey, doorLoopSize)
	if cardEncryptionKey != doorEncryptionKey {
		panic("The derived keys must match")
	}

	return doorEncryptionKey
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
}
