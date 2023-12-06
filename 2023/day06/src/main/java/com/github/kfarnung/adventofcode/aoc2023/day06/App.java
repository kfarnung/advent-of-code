/*
 * Solution for Advent of Code 2023 day 6.
 */
package com.github.kfarnung.adventofcode.aoc2023.day06;

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
        String[] times = lines.get(0).split("\s+");
        String[] distances = lines.get(1).split("\s+");

        long total = 1;
        for (int i = 1; i < times.length; i++) {
            long time = Long.parseLong(times[i]);
            long distance = Long.parseLong(distances[i]);

            total *= countWinners(time, distance);
        }

        return Long.toString(total);
    }

    public static String getPart2(List<String> lines) {
        String[] times = lines.get(0).replace(" ", "").split(":");
        String[] distances = lines.get(1).replace(" ", "").split(":");
        long time = Long.parseLong(times[1]);
        long distance = Long.parseLong(distances[1]);

        return Long.toString(countWinners(time, distance));
    }

    private static long countWinners(long time, long distance) {
        long count = 0;
        for (long j = 0; j <= time; j++) {
            long distanceTraveled = j * (time - j);
            if (distanceTraveled > distance) {
                count++;
            }
        }

        return count;
    }
}
