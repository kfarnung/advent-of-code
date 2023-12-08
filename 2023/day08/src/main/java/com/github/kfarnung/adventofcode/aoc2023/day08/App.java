/*
 * Solution for Advent of Code 2023 day 8.
 */
package com.github.kfarnung.adventofcode.aoc2023.day08;


import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromFile;
import static com.github.kfarnung.adventofcode.aoc2023.utilities.MathUtils.leastCommonMultiple;

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
        char[] directions = lines.get(0).toCharArray();
        Map<String, String[]> map = new HashMap<>();
        for (int i = 2; i < lines.size(); i++) {
            String[] parts = lines.get(i).split(" = ");
            String[] values = parts[1].substring(1, parts[1].length() - 1).split(", ");
            map.put(parts[0], values);
        }

        long count = 0;
        int directionsIndex = 0;
        String current = "AAA";
        while (!current.equals("ZZZ")) {
            String[] values = map.get(current);
            if (directions[directionsIndex] == 'L') {
                current = values[0];
            } else {
                current = values[1];
            }
            directionsIndex = (directionsIndex + 1) % directions.length;
            count++;
        }

        return Long.toString(count);
    }

    public static String getPart2(List<String> lines) {
        char[] directions = lines.get(0).toCharArray();
        Map<String, String[]> map = new HashMap<>();
        List<String> currentNodes = new ArrayList<>();
        for (int i = 2; i < lines.size(); i++) {
            String[] parts = lines.get(i).split(" = ");
            String[] values = parts[1].substring(1, parts[1].length() - 1).split(", ");
            map.put(parts[0], values);
            if (parts[0].endsWith("A")) {
                currentNodes.add(parts[0]);
            }
        }

        long count = 1;
        for (String current : currentNodes) {
            count = leastCommonMultiple(count, findCycle(map, current, directions));
        }

        return Long.toString(count);
    }

    private static long findCycle(Map<String, String[]> map, String current, char[] directions) {
        Map<String, Long> stateMap = new HashMap<>();
        long count = 0;
        int directionsIndex = 0;
        while (!stateMap.containsKey(current + ":" + directionsIndex)) {
            stateMap.put(current + ":" + directionsIndex, count);
            String[] values = map.get(current);
            if (directions[directionsIndex] == 'L') {
                current = values[0];
            } else {
                current = values[1];
            }
            directionsIndex = (directionsIndex + 1) % directions.length;
            count++;
        }

        return count - stateMap.get(current + ":" + directionsIndex);
    }
}
