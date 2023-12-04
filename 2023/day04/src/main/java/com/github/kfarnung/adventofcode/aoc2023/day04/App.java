/*
 * Solution for Advent of Code 2023 day 4.
 */
package com.github.kfarnung.adventofcode.aoc2023.day04;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

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
            int score = scoreCard(line);
            if (score > 0) {
                total += Math.pow(2, score - 1);
            }
        }
        return Long.toString(total);
    }

    public static String getPart2(List<String> lines) {
        List<Integer> scores = new ArrayList<>();
        for (String line : lines) {
            scores.add(scoreCard(line));
        }

        long finalScore = scores.size();
        Long[] finalScores = new Long[scores.size()];
        
        for (int i = scores.size() - 1; i >= 0; i--) {
            int score = scores.get(i);
            long total = score;
            for (int j = i + 1, end = Math.min(i + score + 1, scores.size() - 1); j < end; j++) {
                total += finalScores[j];
            }
            finalScores[i] = total;
            finalScore += total;
        }

        return Long.toString(finalScore);
    }

    private static Integer scoreCard(String line) {
        int colonIndex = line.indexOf(':');
        int pipeIndex = line.indexOf('|', colonIndex);
        String winningString = line.substring(colonIndex + 1, pipeIndex).trim();
        String numbersString = line.substring(pipeIndex + 1).trim();
        Set<Integer> winningNumbers = parseNumbers(winningString);
        Set<Integer> numbers = parseNumbers(numbersString);
        numbers.retainAll(winningNumbers);
        return numbers.size();
    }

    private static Set<Integer> parseNumbers(String numbersString) {
        Set<Integer> result = new HashSet<>();
        String[] numbers = numbersString.split("\\s+");
        for (String number : numbers) {
            result.add(Integer.parseInt(number));
        }
        return result;
    }
}
