package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

func expandRule(rules map[string]string, ruleID string) string {
	rule := rules[ruleID]
	var parts []string

	for _, subruleList := range strings.Split(rule, " | ") {
		repetition := 1
		var ruleParts []string

		for _, subrule := range strings.Split(subruleList, " ") {
			if strings.HasPrefix(subrule, "\"") {
				if !strings.HasSuffix(subrule, "\"") {
					panic("Couldn't find end quote")
				}

				ruleParts = append(ruleParts, subrule[1:len(subrule)-1])
			} else if subrule != ruleID {
				ruleParts = append(ruleParts, expandRule(rules, subrule))
			} else {
				repetition = 10
			}
		}

		// Hack for dealing with recursive rules. Will only match up to 10
		// instances, but that seems to be plenty for this problem. Basically
		// duplicate the rules with an explicit repetition count to ensure that
		// "aab" won't match, only "aabb" will (equal number of "a" and "b").
		for i := 0; i < repetition; i++ {
			var sb strings.Builder

			for _, val := range ruleParts {
				sb.WriteString(val)

				if i > 0 {
					sb.WriteString(fmt.Sprintf("{%d}", i+1))
				}
			}

			parts = append(parts, sb.String())
		}
	}

	if len(parts) == 1 {
		return parts[0]
	}

	return "(?:" + strings.Join(parts, "|") + ")"
}

func parseInput(lines []string) (map[string]string, []string) {
	ruleRegex := regexp.MustCompile(`^(\d+): (.+)$`)
	rules := make(map[string]string)
	var data []string
	foundData := false

	for _, line := range lines {
		if len(line) == 0 {
			foundData = true
		} else if !foundData {
			match := ruleRegex.FindStringSubmatch(line)
			rules[match[1]] = match[2]
		} else if foundData {
			data = append(data, line)
		} else {
			panic("Unexpected state")
		}
	}

	return rules, data
}

func part1(lines []string) int {
	rules, data := parseInput(lines)
	expanded := expandRule(rules, "0")
	matchRegex := regexp.MustCompile("^" + expanded + "$")

	count := 0
	for _, candidate := range data {
		if matchRegex.MatchString(candidate) {
			count++
		}
	}

	return count
}

func part2(lines []string) int {
	rules, data := parseInput(lines)
	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"

	expanded := expandRule(rules, "0")
	matchRegex := regexp.MustCompile("^" + expanded + "$")

	count := 0
	for _, candidate := range data {
		if matchRegex.MatchString(candidate) {
			count++
		}
	}

	return count
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
