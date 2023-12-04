/*
 * Solution for Advent of Code 2023 day 5.
 */
package com.github.kfarnung.adventofcode.aoc2023.day05;

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
        long minLocation = Long.MAX_VALUE;

        SeedMaps seedMaps = new SeedMaps(lines);
        for (Long seed : seedMaps.getSeeds()) {
            long location = seedMaps.getLocationForSeed(seed);
            if (location < minLocation) {
                minLocation = location;
            }
        }

        return Long.toString(minLocation);
    }

    public static String getPart2(List<String> lines) {
        long minLocation = Long.MAX_VALUE;

        SeedMaps seedMaps = new SeedMaps(lines);
        List<Long> seeds = seedMaps.getSeeds();
        for (int i = 0; i < seeds.size(); i += 2) {
            long start = seeds.get(i);
            long count = seeds.get(i + 1);
            Range range = new Range(start, start + count - 1);
            List<Range> ranges = seedMaps.getLocationRangesForSeed(range);
            for (Range r : ranges) {
                if (r.getStart() < minLocation) {
                    minLocation = r.getStart();
                }
            }
        }

        return Long.toString(minLocation);
    }
}
