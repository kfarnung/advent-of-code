/*
 * Solution for Advent of Code 2023 day 10.
 */
package com.github.kfarnung.adventofcode.aoc2023.day10;

import java.io.IOException;
import java.util.ArrayDeque;
import java.util.HashSet;
import java.util.List;
import java.util.Queue;
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
        Graph graph = new Graph(lines);
        Coordinate start = graph.getStart();
        Queue<SearchState> queue = new ArrayDeque<>();
        Set<Coordinate> visited = new HashSet<>();
        long distance = 0;
        queue.add(new SearchState(start, 0));

        while (!queue.isEmpty()) {
            SearchState state = queue.poll();
            Coordinate coordinate = state.getCoordinate();
            GraphNode node = graph.get(coordinate);
            if (node == null) {
                continue;
            }

            visited.add(coordinate);
            distance = Math.max(distance, state.getDistance());

            for (Coordinate neighbor : node.getNeighbors()) {
                if (visited.contains(neighbor)) {
                    continue;
                }

                GraphNode neighborNode = graph.get(neighbor);
                if (neighborNode == null || !neighborNode.getNeighbors().contains(coordinate)) {
                    continue;
                }
                queue.add(new SearchState(neighbor, state.getDistance() + 1));
            }
        }
        
        return Long.toString(distance);
    }

    public static String getPart2(List<String> lines) {
        return "0";
    }
}
