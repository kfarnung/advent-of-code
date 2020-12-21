package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type foodItem struct {
	ingredients []string
	allergens   []string
}

type ingredientSorter struct {
	items []string
	by    func(p1, p2 *string) bool
}

// Len is part of sort.Interface.
func (s *ingredientSorter) Len() int {
	return len(s.items)
}

// Swap is part of sort.Interface.
func (s *ingredientSorter) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *ingredientSorter) Less(i, j int) bool {
	return s.by(&s.items[i], &s.items[j])
}

func intersectSlices(slice1 []string, slice2 []string) []string {
	var intersection []string
	for _, value1 := range slice1 {
		for _, value2 := range slice2 {
			if value1 == value2 {
				intersection = append(intersection, value1)
			}
		}
	}

	return intersection
}

func unmatchedSliceItems(slice []string, matched map[string]bool) []string {
	var unmatched []string
	for _, item := range slice {
		if used := matched[item]; !used {
			unmatched = append(unmatched, item)
		}
	}

	return unmatched
}

func mapIngredientsToAllergens(foodItems []foodItem) map[string]string {
	allergenFoodMap := make(map[string][]foodItem)
	for _, foodItem := range foodItems {
		for _, allergen := range foodItem.allergens {
			allergenFoodMap[allergen] = append(allergenFoodMap[allergen], foodItem)
		}
	}

	allergenIngredientsMap := make(map[string][]string)
	for allergen, foodList := range allergenFoodMap {
		var ingredientCandidates []string
		for _, foodItem := range foodList {
			if ingredientCandidates == nil {
				ingredientCandidates = foodItem.ingredients
			} else {
				ingredientCandidates = intersectSlices(ingredientCandidates, foodItem.ingredients)
			}
		}

		allergenIngredientsMap[allergen] = ingredientCandidates
	}

	matched := make(map[string]bool)
	ingredientAllergenMap := make(map[string]string)
	for {
		madeProgress := false
		for allergen, ingredients := range allergenIngredientsMap {
			unmatched := unmatchedSliceItems(ingredients, matched)
			if len(unmatched) == 1 {
				ingredient := unmatched[0]
				matched[ingredient] = true
				ingredientAllergenMap[ingredient] = allergen
				madeProgress = true
			}
		}

		if !madeProgress {
			break
		}
	}

	if len(allergenIngredientsMap) != len(ingredientAllergenMap) {
		panic("Couldn't match all allergens and ingredients")
	}

	return ingredientAllergenMap
}

func parseLine(line string) (foodItem, error) {
	foodRegex := regexp.MustCompile(`^(.+) \(contains (.+)\)$`)
	match := foodRegex.FindStringSubmatch(line)
	if match == nil {
		return foodItem{}, errors.New("Failed to parse line")
	}

	return foodItem{
		ingredients: strings.Split(match[1], " "),
		allergens:   strings.Split(match[2], ", "),
	}, nil
}

func parseInput(lines []string) ([]foodItem, error) {
	var foodItems []foodItem
	for _, line := range lines {
		foodItem, err := parseLine(line)
		if err != nil {
			return nil, err
		}

		foodItems = append(foodItems, foodItem)
	}

	return foodItems, nil
}

func part1(lines []string) int {
	foodItems, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	ingredientAllergenMap := mapIngredientsToAllergens(foodItems)

	count := 0
	for _, foodItem := range foodItems {
		for _, ingredient := range foodItem.ingredients {
			if _, ok := ingredientAllergenMap[ingredient]; !ok {
				count++
			}
		}
	}

	return count
}

func part2(lines []string) string {
	foodItems, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	ingredientAllergenMap := mapIngredientsToAllergens(foodItems)
	var ingredients []string
	for key := range ingredientAllergenMap {
		ingredients = append(ingredients, key)
	}

	sort.Sort(&ingredientSorter{
		items: ingredients,
		by: func(p1, p2 *string) bool {
			allergen1 := ingredientAllergenMap[*p1]
			allergen2 := ingredientAllergenMap[*p2]

			return allergen1 < allergen2
		},
	})

	return strings.Join(ingredients, ",")
}

func main() {
	name := os.Args[1]
	lines, err := lib.LoadFileLines(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %s\n", part2(lines))
}
