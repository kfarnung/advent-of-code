/*
 * Solution for Advent of Code 2023 day 7.
 */
package com.github.kfarnung.adventofcode.aoc2023.day07;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromFile;

public class App {
    public static void main(String[] args) {
        try {
            List<String> lines = readLinesFromFile(args[0]);
            System.out.printf("Part 1: %s%n", getPart1(lines));
            System.out.printf("Part 2: %s%n", getPart2(lines));
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static String getPart1(List<String> lines) {
        return calculateScore(lines, false);
    }

    public static String getPart2(List<String> lines) {
        return calculateScore(lines, true);
    }

    public static String calculateScore(List<String> lines, boolean jokersWild) {
        
        List<Hand> hands = new ArrayList<>();
        for (String line : lines) {
            hands.add(new Hand(line, jokersWild));
        }
        hands.sort(null);
        long total = 0;
        for (int i = 0; i < hands.size(); i++) {
            total += hands.get(i).getBid() * (i + 1);
        }
        return Long.toString(total);
    }
}
