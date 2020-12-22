package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type spaceCardDeck []int32

func (s spaceCardDeck) clone(cardCount int) spaceCardDeck {
	newDeck := make(spaceCardDeck, cardCount)
	copy(newDeck, s[:cardCount])
	return newDeck
}

func (s *spaceCardDeck) draw() int32 {
	topCard := (*s)[0]
	*s = (*s)[1:]

	return topCard
}

func (s *spaceCardDeck) discard(card ...int32) {
	*s = append(*s, card...)
}

func (s spaceCardDeck) score() int64 {
	var total int64
	for i, card := range s {
		total += int64(card) * int64(len(s)-i)
	}

	return total
}

func (s *spaceCardDeck) String() string {
	var sb strings.Builder

	// Give the string builder enough room for the string without having to grow.
	sb.Grow(len(*s) * 4)

	for i, card := range *s {
		if i > 0 {
			sb.WriteString(", ")
		}

		fmt.Fprintf(&sb, "%d", card)
	}

	return sb.String()
}

func playGame(playerDeck1, playerDeck2 *spaceCardDeck, recursive bool) int {
	deckStates := make(map[string]bool)
	for len(*playerDeck1) > 0 && len(*playerDeck2) > 0 {
		deckState := playerDeck1.String() + ":" + playerDeck2.String()
		if val := deckStates[deckState]; val {
			return 1
		}

		deckStates[deckState] = true
		playerCard1 := playerDeck1.draw()
		playerCard2 := playerDeck2.draw()

		winner := 0
		if recursive && int(playerCard1) <= len(*playerDeck1) && int(playerCard2) <= len(*playerDeck2) {
			playerDeckClone1 := playerDeck1.clone(int(playerCard1))
			playerDeckClone2 := playerDeck2.clone(int(playerCard2))

			winner = playGame(&playerDeckClone1, &playerDeckClone2, recursive)
		} else if playerCard1 > playerCard2 {
			winner = 1
		} else if playerCard2 > playerCard1 {
			winner = 2
		} else {
			panic("There should be no ties")
		}

		switch winner {
		case 1:
			playerDeck1.discard(playerCard1, playerCard2)
		case 2:
			playerDeck2.discard(playerCard2, playerCard1)
		default:
			panic("Invalid winner specified")
		}
	}

	if len(*playerDeck1) > 0 {
		return 1
	} else if len(*playerDeck2) > 0 {
		return 2
	} else {
		panic("There should be a winner...")
	}
}

func parseInput(lines []string) ([]spaceCardDeck, error) {
	var decks []spaceCardDeck
	var current spaceCardDeck
	for _, line := range lines {
		if strings.HasPrefix(line, "Player ") {
			if len(current) > 0 {
				decks = append(decks, current)
				current = spaceCardDeck{}
			}
		} else if len(line) > 0 {
			value, err := lib.ParseInt32(line)
			if err != nil {
				return nil, err
			}

			current.discard(value)
		}
	}

	if len(current) > 0 {
		decks = append(decks, current)
		current = spaceCardDeck{}
	}

	return decks, nil
}

func run(lines []string, recursive bool) int64 {
	spaceCardDecks, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	playerDeck1 := spaceCardDecks[0]
	playerDeck2 := spaceCardDecks[1]
	winner := playGame(&playerDeck1, &playerDeck2, recursive)

	switch winner {
	case 1:
		return playerDeck1.score()
	case 2:
		return playerDeck2.score()
	default:
		panic("There should be a winner...")
	}
}

func part1(lines []string) int64 {
	return run(lines, false)
}

func part2(lines []string) int64 {
	return run(lines, true)
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
