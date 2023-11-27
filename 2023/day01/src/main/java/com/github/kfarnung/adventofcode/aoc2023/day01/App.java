/*
 * Solution for Advent of Code 2023 day 1.
 */
package com.github.kfarnung.adventofcode.aoc2023.day01;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromFile;

public class App {
    private static final List<String> NUMBERS = List.of(
            "one",
            "two",
            "three",
            "four",
            "five",
            "six",
            "seven",
            "eight",
            "nine");

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
        long x = 0;
        for (String string : lines) {
            List<Character> chars = string.chars()
                    .filter(Character::isDigit)
                    .mapToObj(e -> (char) e)
                    .toList();
            x += (chars.get(0) - '0') * 10;
            x += chars.get(chars.size() - 1) - '0';
        }

        return Long.toString(x);
    }

    public static String getPart2(List<String> lines) {
        long x = 0;
        for (String string : lines) {
            List<Character> chars = new ArrayList<Character>(string.chars()
                    .mapToObj(e -> (char) e)
                    .toList());
            for (int i = 0; i < NUMBERS.size(); i++) {
                int index = 0;
                while ((index = string.indexOf(NUMBERS.get(i), index)) != -1) {
                    chars.set(index, (char) (i + '1'));
                    index++;
                }
            }
            chars = chars.stream().filter(Character::isDigit).toList();

            x += (chars.get(0) - '0') * 10;
            x += chars.get(chars.size() - 1) - '0';
        }

        return Long.toString(x);
    }
}
