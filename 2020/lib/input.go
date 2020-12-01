package lib

import (
	"bufio"
	"log"
	"os"
)

// LoadFileLines loads lines from a given file
func LoadFileLines(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
