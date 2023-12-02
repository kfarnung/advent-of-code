/*
 * Solution for Advent of Code 2023 day 3.
 */
package com.github.kfarnung.adventofcode.aoc2023.day03;

import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
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
        List<List<Character>> rows = lines.stream()
                .map(line -> line.chars().mapToObj(c -> (char) c).toList())
                .toList();

        long total = 0;
        StringBuilder sb = new StringBuilder();
        List<Part> symbols = new ArrayList<>();
        for (int i = 0; i < rows.size(); i++) {
            List<Character> row = rows.get(i);
            for (int j = 0; j < row.size(); j++) {
                Character c = row.get(j);
                if (Character.isDigit(c)) {
                    sb.append(c);
                    for (Part symbol : findSymbols(rows, i, j)) {
                        symbols.add(symbol);
                    }
                } else {
                    if (!symbols.isEmpty()) {
                        total += Integer.parseInt(sb.toString());
                    }

                    sb.setLength(0);
                    symbols.clear();
                }
            }
        }

        return Long.toString(total);
    }

    public static String getPart2(List<String> lines) {
        List<List<Character>> rows = lines.stream()
                .map(line -> line.chars().mapToObj(c -> (char) c).toList())
                .toList();

        Map<Part, List<Integer>> symbolMap = new HashMap<>();
        StringBuilder sb = new StringBuilder();
        Set<Part> symbols = new HashSet<>();
        for (int i = 0; i < rows.size(); i++) {
            List<Character> row = rows.get(i);
            for (int j = 0; j < row.size(); j++) {
                Character c = row.get(j);
                if (Character.isDigit(c)) {
                    sb.append(c);
                    for (Part symbol : findSymbols(rows, i, j)) {
                        if (symbol.getSymbol() == '*') {
                            symbols.add(symbol);
                        }
                    }
                } else {
                    for (Part symbol : symbols) {
                        symbolMap.computeIfAbsent(symbol, k -> new ArrayList<>()).add(Integer.parseInt(sb.toString()));
                    }

                    sb.setLength(0);
                    symbols.clear();
                }
            }
        }

        long total = 0;
        for (List<Integer> values : symbolMap.values()) {
            if (values.size() == 2) {
                long product = 1;
                for (Integer value : values) {
                    product *= value;
                }

                total += product;
            }
        }

        return Long.toString(total);
    }

    private static Set<Part> findSymbols(List<List<Character>> rows, int x, int y) {
        Set<Part> symbols = new HashSet<>();
        for (int i = -1; i <= 1; i++) {
            int rowIndex = x + i;
            if (rowIndex < 0 || rowIndex >= rows.size()) {
                continue;
            }

            List<Character> row = rows.get(rowIndex);

            for (int j = -1; j <= 1; j++) {
                if (i == 0 && j == 0) {
                    continue;
                }

                int colIndex = y + j;
                if (colIndex < 0 || colIndex >= rows.get(rowIndex).size()) {
                    continue;
                }

                Character c = row.get(colIndex);

                if (!Character.isDigit(c) && c != '.') {
                    symbols.add(new Part(c, rowIndex, colIndex));
                }
            }
        }

        return symbols;
    }
}
