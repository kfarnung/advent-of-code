package lib

import (
	"os"
	"strings"
)

func LoadFileContent(name string) (string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}
