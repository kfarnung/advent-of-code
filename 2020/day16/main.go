package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type ticketRange struct {
	lower int
	upper int
}

type ticketRule struct {
	name   string
	ranges []ticketRange
}

func (t ticketRule) InRange(value int) bool {
	for _, r := range t.ranges {
		if value >= r.lower && value <= r.upper {
			return true
		}
	}

	return false
}

func aggregateRules(rules []ticketRule) map[int]bool {
	aggregatedRules := make(map[int]bool)
	for _, rule := range rules {
		for _, r := range rule.ranges {
			for i := r.lower; i <= r.upper; i++ {
				aggregatedRules[i] = true
			}
		}
	}

	return aggregatedRules
}

func filterInvalid(rules []ticketRule, tickets [][]int) [][]int {
	aggregatedRules := aggregateRules(rules)

	var validTickets [][]int
	for _, ticket := range tickets {
		valid := true
		for _, field := range ticket {
			if !aggregatedRules[field] {
				valid = false
				break
			}
		}

		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func filterEntries(rules []ticketRule, names map[string]bool) []ticketRule {
	var filteredEntries []ticketRule
	for _, rule := range rules {
		if !names[rule.name] {
			filteredEntries = append(filteredEntries, rule)
		}
	}

	return filteredEntries
}

func reduceRules(possibleRules [][]ticketRule) []ticketRule {
	reducedRules := make([]ticketRule, len(possibleRules))
	reducedRuleNames := make(map[string]bool)

	for len(reducedRuleNames) < len(possibleRules) {
		found := false
		for i, ruleList := range possibleRules {
			filteredEntries := filterEntries(ruleList, reducedRuleNames)
			if len(filteredEntries) == 1 {
				rule := filteredEntries[0]
				reducedRules[i] = rule
				reducedRuleNames[rule.name] = true
				found = true
			}
		}

		if !found {
			panic("Didn't find a reduction")
		}
	}

	return reducedRules
}

func parseRules(lines []string) ([]ticketRule, error) {
	var rules []ticketRule
	ruleRegex := regexp.MustCompile(`^([^:]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

	for _, line := range lines {
		if len(line) == 0 {
			return rules, nil
		}

		match := ruleRegex.FindStringSubmatch(line)
		if match == nil {
			return nil, errors.New("Couldn't match line")
		}

		rule := ticketRule{name: match[1]}

		lower, err := lib.ParseInt(match[2])
		if err != nil {
			return nil, err
		}
		upper, err := lib.ParseInt(match[3])
		if err != nil {
			return nil, err
		}
		rule.ranges = append(rule.ranges, ticketRange{
			lower: lower,
			upper: upper,
		})

		lower, err = lib.ParseInt(match[4])
		if err != nil {
			return nil, err
		}
		upper, err = lib.ParseInt(match[5])
		if err != nil {
			return nil, err
		}
		rule.ranges = append(rule.ranges, ticketRange{
			lower: lower,
			upper: upper,
		})

		rules = append(rules, rule)
	}

	return nil, errors.New("Couldn't find the end of the rules")
}

func parseTicket(line string) ([]int, error) {
	numberStrings := strings.Split(line, ",")
	return lib.StringSliceToInt(numberStrings)
}

func parseYourTicket(lines []string) ([]int, error) {
	found := false
	for _, line := range lines {
		if line == "your ticket:" {
			found = true
		} else if found {
			return parseTicket(line)
		}
	}

	return nil, errors.New("Unable to locate your ticket information")
}

func parseNearbyTickets(lines []string) ([][]int, error) {
	var tickets [][]int
	found := false
	for _, line := range lines {
		if line == "nearby tickets:" {
			found = true
		} else if found {
			ticket, err := parseTicket(line)
			if err != nil {
				return nil, err
			}

			tickets = append(tickets, ticket)
		}
	}

	if !found {
		return nil, errors.New("Unable to locate nearby ticket information")
	}

	return tickets, nil
}

func parseInput(lines []string) ([]ticketRule, []int, [][]int, error) {
	rules, err := parseRules(lines)
	if err != nil {
		return nil, nil, nil, err
	}

	yourTicket, err := parseYourTicket(lines)
	if err != nil {
		return nil, nil, nil, err
	}

	nearbyTickets, err := parseNearbyTickets(lines)
	if err != nil {
		return nil, nil, nil, err
	}

	return rules, yourTicket, nearbyTickets, nil
}

func part1(lines []string) int {
	rules, _, nearbyTickets, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	aggregatedRules := aggregateRules(rules)

	sumOfInvalid := 0
	for _, ticket := range nearbyTickets {
		for _, field := range ticket {
			if !aggregatedRules[field] {
				sumOfInvalid += field
			}
		}
	}

	return sumOfInvalid
}

func part2(lines []string) int {
	rules, yourTicket, nearbyTickets, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	validNearby := filterInvalid(rules, nearbyTickets)

	ticketsToCheck := [][]int{yourTicket}
	ticketsToCheck = append(ticketsToCheck, validNearby...)

	possibleRules := make([][]ticketRule, len(yourTicket))
	for _, rule := range rules {
		for i := 0; i < len(yourTicket); i++ {
			valid := true
			for _, ticket := range ticketsToCheck {
				if !rule.InRange(ticket[i]) {
					valid = false
					break
				}
			}

			if valid {
				possibleRules[i] = append(possibleRules[i], rule)
			}
		}
	}

	reducedRules := reduceRules(possibleRules)
	total := 1
	for i, rule := range reducedRules {
		if strings.HasPrefix(rule.name, "departure") {
			total *= yourTicket[i]
		}
	}

	return total
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
