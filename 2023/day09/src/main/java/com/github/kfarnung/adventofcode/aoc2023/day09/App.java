/*
 * Solution for Advent of Code 2023 day 9.
 */
package com.github.kfarnung.adventofcode.aoc2023.day09;

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
        long total = 0;
        for (String line : lines) {
            List<Long> numbers = parseList(line);
            total += predictValue(numbers, false);
        }
        return Long.toString(total);
    }

    public static String getPart2(List<String> lines) {
        long total = 0;
        for (String line : lines) {
            List<Long> numbers = parseList(line);
            total += predictValue(numbers, true);
        }
        return Long.toString(total);
    }

    private static List<Long> parseList(String line) {
        String[] numbers = line.split(" ");
        List<Long> numberList = new ArrayList<>();
        for (String number : numbers) {
            numberList.add(Long.parseLong(number));
        }

        return numberList;
    }

    private static long predictValue(List<Long> numbers, boolean reverse) {
        boolean nonZero = false;
        List<Long> differences = new ArrayList<>();
        for (int i = 0; i < numbers.size() - 1; i++) {
            long difference = numbers.get(i + 1) - numbers.get(i);
            differences.add(difference);
            if (difference != 0) {
                nonZero = true;
            }
        }

        long lastDifference = 0;
        if (nonZero) {
            lastDifference = predictValue(differences, reverse);
        }

        if (reverse) {
            return numbers.get(0) - lastDifference;
        } else {
            return numbers.get(numbers.size() - 1) + lastDifference;
        }
    }
}
