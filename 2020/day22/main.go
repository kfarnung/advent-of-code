package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kfarnung/advent-of-code/2020/lib"
)

type spaceCard int

func (s spaceCard) String() string {
	return fmt.Sprintf("%d", s)
}

type spaceCardDeck []spaceCard

func (s spaceCardDeck) Clone(cardCount int) spaceCardDeck {
	newDeck := make(spaceCardDeck, cardCount)
	for i, card := range s {
		if i == cardCount {
			break
		}

		newDeck[i] = card
	}

	return newDeck
}

func (s *spaceCardDeck) DrawTopCard() spaceCard {
	topCard := (*s)[0]
	*s = (*s)[1:]

	return topCard
}

func (s *spaceCardDeck) DiscardToBottom(card spaceCard) {
	*s = append(*s, card)
}

func (s spaceCardDeck) Len() int {
	return len(s)
}

func (s spaceCardDeck) Score() int {
	var total int
	for i, card := range s {
		total += int(card) * (len(s) - i)
	}

	return total
}

func (s spaceCardDeck) String() string {
	var sb strings.Builder
	for i, card := range s {
		if i > 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(card.String())
	}

	return sb.String()
}

func playGame(gameOutcomes map[string]int, playerDeck1, playerDeck2 *spaceCardDeck, recursive bool) int {
	initialDeckState := fmt.Sprintf("%s:%s", *playerDeck1, *playerDeck2)
	if winner, ok := gameOutcomes[initialDeckState]; ok {
		return winner
	}

	deckStates := make(map[string]bool)
	roundNumber := 1
	for playerDeck1.Len() > 0 && playerDeck2.Len() > 0 {
		deckState := fmt.Sprintf("%s:%s", *playerDeck1, *playerDeck2)
		if val := deckStates[deckState]; val {
			gameOutcomes[initialDeckState] = 1
			return 1
		}

		deckStates[deckState] = true

		playerCard1 := playerDeck1.DrawTopCard()

		playerCard2 := playerDeck2.DrawTopCard()

		winner := 0
		if recursive && int(playerCard1) <= playerDeck1.Len() && int(playerCard2) <= playerDeck2.Len() {
			playerDeckClone1 := playerDeck1.Clone(int(playerCard1))
			playerDeckClone2 := playerDeck2.Clone(int(playerCard2))

			winner = playGame(gameOutcomes, &playerDeckClone1, &playerDeckClone2, recursive)
		} else if playerCard1 > playerCard2 {
			winner = 1
		} else if playerCard2 > playerCard1 {
			winner = 2
		} else {
			panic("There should be no ties")
		}

		switch winner {
		case 1:
			playerDeck1.DiscardToBottom(playerCard1)
			playerDeck1.DiscardToBottom(playerCard2)

		case 2:
			playerDeck2.DiscardToBottom(playerCard2)
			playerDeck2.DiscardToBottom(playerCard1)
		default:
			panic("Invalid winner specified")
		}

		roundNumber++
	}

	if playerDeck1.Len() > 0 {
		gameOutcomes[initialDeckState] = 1
		return 1
	} else if playerDeck2.Len() > 0 {
		gameOutcomes[initialDeckState] = 2
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
			if current.Len() > 0 {
				decks = append(decks, current)
				current = spaceCardDeck{}
			}
		} else if len(line) > 0 {
			value, err := lib.ParseInt(line)
			if err != nil {
				return nil, err
			}

			card := spaceCard(value)
			current.DiscardToBottom(card)
		}
	}

	if current.Len() > 0 {
		decks = append(decks, current)
		current = spaceCardDeck{}
	}

	return decks, nil
}

func run(lines []string, recursive bool) int {
	spaceCardDecks, err := parseInput(lines)
	if err != nil {
		log.Fatal(err)
	}

	gameOutcomes := make(map[string]int)
	playerDeck1 := spaceCardDecks[0]
	playerDeck2 := spaceCardDecks[1]
	playGame(gameOutcomes, &playerDeck1, &playerDeck2, recursive)

	if playerDeck1.Len() > 0 {
		return playerDeck1.Score()
	} else if playerDeck2.Len() > 0 {
		return playerDeck2.Score()
	} else {
		panic("There should be a winner...")
	}
}

func part1(lines []string) int {
	return run(lines, false)
}

func part2(lines []string) int {
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
