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

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) parseLine(line string) error {
	tokens := strings.Split(line, " ")
	for _, token := range tokens {
		keyValue := strings.Split(token, ":")
		if len(keyValue) != 2 {
			return errors.New("Failed to parse token, exactly two components required")
		}

		switch keyValue[0] {
		case "byr":
			p.byr = keyValue[1]
		case "iyr":
			p.iyr = keyValue[1]
		case "eyr":
			p.eyr = keyValue[1]
		case "hgt":
			p.hgt = keyValue[1]
		case "hcl":
			p.hcl = keyValue[1]
		case "ecl":
			p.ecl = keyValue[1]
		case "pid":
			p.pid = keyValue[1]
		case "cid":
			p.cid = keyValue[1]
		default:
			return errors.New("Invalid key specified")
		}
	}

	return nil
}

func (p passport) isValidBirthYear() bool {
	year, err := lib.ParseInt32(p.byr)
	if err != nil {
		return false
	}

	return year >= 1920 && year <= 2002
}

func (p passport) isValidIssueYear() bool {
	year, err := lib.ParseInt32(p.iyr)
	if err != nil {
		return false
	}

	return year >= 2010 && year <= 2020
}

func (p passport) isValidExpirationYear() bool {
	year, err := lib.ParseInt32(p.eyr)
	if err != nil {
		return false
	}

	return year >= 2020 && year <= 2030
}

func (p passport) isValidHeight() bool {
	heightRegex := regexp.MustCompile(`^(\d+)(in|cm)$`)
	match := heightRegex.FindStringSubmatch(p.hgt)
	if match == nil {
		return false
	}

	height, err := lib.ParseInt32(match[1])
	if err != nil {
		return false
	}

	switch match[2] {
	case "cm":
		return height >= 150 && height <= 193
	case "in":
		return height >= 59 && height <= 76
	}

	return false
}

func (p passport) isValidHairColor() bool {
	hairRegex := regexp.MustCompile(`^#[a-f0-9]+$`)
	return hairRegex.MatchString(p.hcl)
}

func (p passport) isValidEyeColor() bool {
	return p.ecl == "amb" ||
		p.ecl == "blu" ||
		p.ecl == "brn" ||
		p.ecl == "gry" ||
		p.ecl == "grn" ||
		p.ecl == "hzl" ||
		p.ecl == "oth"
}

func (p passport) isValidPassportID() bool {
	idRegex := regexp.MustCompile(`^\d{9}$`)
	return idRegex.MatchString(p.pid)
}

func (p *passport) isValidPart1() bool {
	return len(p.byr) > 0 &&
		len(p.ecl) > 0 &&
		len(p.eyr) > 0 &&
		len(p.hcl) > 0 &&
		len(p.hgt) > 0 &&
		len(p.iyr) > 0 &&
		len(p.pid) > 0
}

func (p *passport) isValidPart2() bool {
	return p.isValidBirthYear() &&
		p.isValidExpirationYear() &&
		p.isValidEyeColor() &&
		p.isValidHairColor() &&
		p.isValidHeight() &&
		p.isValidIssueYear() &&
		p.isValidPassportID()
}

func parsePassports(lines []string) ([]passport, error) {
	var passports []passport
	current := passport{}
	for _, line := range lines {
		if len(line) > 0 {
			if err := current.parseLine(line); err != nil {
				return nil, err
			}
		} else {
			passports = append(passports, current)
			current = passport{}
		}
	}

	passports = append(passports, current)

	return passports, nil
}

func parseInput(name string) ([]passport, error) {
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		return nil, err
	}

	return parsePassports(lines)
}

func part1(passports []passport) int {
	count := 0
	for _, passport := range passports {
		if passport.isValidPart1() {
			count++
		}
	}

	return count
}

func part2(passports []passport) int {
	count := 0
	for _, passport := range passports {
		if passport.isValidPart2() {
			count++
		}
	}

	return count
}

func main() {
	name := os.Args[1]
	passports, err := parseInput(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(passports))
	fmt.Printf("Part 2: %d\n", part2(passports))
}
