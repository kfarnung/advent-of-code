package com.github.kfarnung.adventofcode.aoc2023.day07;

import java.util.HashMap;
import java.util.Map;

public class Hand implements Comparable<Hand> {
    public enum HandType {
        HIGH_CARD,
        PAIR,
        TWO_PAIR,
        THREE_OF_A_KIND,
        FULL_HOUSE,
        FOUR_OF_A_KIND,
        FIVE_OF_A_KIND,
    }

    public enum CardValue {
        JOKER,
        TWO,
        THREE,
        FOUR,
        FIVE,
        SIX,
        SEVEN,
        EIGHT,
        NINE,
        TEN,
        JACK,
        QUEEN,
        KING,
        ACE,
    }

    private final CardValue[] cards;
    private final HandType score;
    private final int bid;

    public Hand(String hand, boolean jokersWild) {
        String[] parts = hand.split(" ");
        this.bid = Integer.parseInt(parts[1]);

        char[] cards = parts[0].toCharArray();
        this.cards = translateCards(cards, jokersWild);
        this.score = calculateScore(cards, jokersWild);
    }

    private static CardValue[] translateCards(char[] cards, boolean jokersWild) {
        CardValue[] cardValues = new CardValue[cards.length];
        for (int i = 0; i < cards.length; i++) {
            cardValues[i] = translateCard(cards[i], jokersWild);
        }
        return cardValues;
    }

    private static CardValue translateCard(char card, boolean jokersWild) {
        switch (card) {
            case '2':
                return CardValue.TWO;
            case '3':    
                return CardValue.THREE;
            case '4':
                return CardValue.FOUR;
            case '5':
                return CardValue.FIVE;
            case '6':
                return CardValue.SIX;
            case '7':
                return CardValue.SEVEN;
            case '8':
                return CardValue.EIGHT;
            case '9':
                return CardValue.NINE;
            case 'T':
                return CardValue.TEN;
            case 'J':
                if (jokersWild) {
                    return CardValue.JOKER;
                }
                return CardValue.JACK;
            case 'Q':
                return CardValue.QUEEN;
            case 'K':
                return CardValue.KING;
            case 'A':
                return CardValue.ACE;
            default:
                throw new IllegalArgumentException("Invalid card: " + card);
        }
    }

    private static HandType calculateScore(char[] cards, boolean jokersWild) {
        Map<Character, Integer> cardCounts = new HashMap<>();
        int jokerCount = 0;
        for (char card : cards) {
            if (jokersWild && card == 'J') {
                jokerCount++;
                continue;
            }

            cardCounts.put(card, cardCounts.getOrDefault(card, 0) + 1);
        }

        if (cardCounts.size() == 1 || jokerCount == 5) {
            return HandType.FIVE_OF_A_KIND;
        } else if (cardCounts.size() == 2) {
            if (cardCounts.containsValue(4 - jokerCount)) {
                return HandType.FOUR_OF_A_KIND;
            } else {
                return HandType.FULL_HOUSE;
            }
        } else if (cardCounts.size() == 3) {
            if (cardCounts.containsValue(3 - jokerCount)) {
                return HandType.THREE_OF_A_KIND;
            } else {
                return HandType.TWO_PAIR;
            }
        } else if (cardCounts.size() == 4) {
            return HandType.PAIR;
        } else if (cardCounts.size() == 5) {
            return HandType.HIGH_CARD;
        } else {
            throw new IllegalArgumentException("Invalid hand: " + new String(cards));
        }
    }

    public int getBid() {
        return bid;
    }

    @Override
    public int compareTo(Hand o) {
        if (score != o.score) {
            return score.compareTo(o.score);
        }

        for (int i = 0; i < cards.length; i++) {
            if (cards[i] != o.cards[i]) {
                return cards[i].compareTo(o.cards[i]);
            }
        }

        return 0;
    }
}
