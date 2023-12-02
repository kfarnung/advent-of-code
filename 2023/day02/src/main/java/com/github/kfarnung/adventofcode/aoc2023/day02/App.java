/*
 * Solution for Advent of Code 2023 day 2.
 */
package com.github.kfarnung.adventofcode.aoc2023.day02;

import java.io.IOException;
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
        long total = 0;
        Subset target = new Subset(12, 13, 14);

        for (String line : lines) {
            Game game = Game.parse(line);
            if (game.isValid(target)) {
                total += game.getId();
            }
        }

        return Long.toString(total);
    }

    public static String getPart2(List<String> lines) {
        long total = 0;
        for (String line : lines) {
            Game game = Game.parse(line);
            total += game.getMinimum().getPower();
        }

        return Long.toString(total);
    }
}
