package com.github.kfarnung.adventofcode.aoc2023.day10;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Graph {
    private final Map<Coordinate, GraphNode> graph;
    private final Coordinate start;

    public Graph(List<String> lines) {
        graph = new HashMap<>();
        Coordinate start = null;
        for (int i = 0; i < lines.size(); i++) {
            for (int j = 0; j < lines.get(i).length(); j++) {
                char c = lines.get(i).charAt(j);
                if (c == 'S') {
                    start = new Coordinate(i, j);
                }
                Coordinate coordinate = new Coordinate(i, j);
                graph.put(coordinate, new GraphNode(coordinate, c));
            }
        }

        this.start = start;
    }

    public GraphNode get(Coordinate coordinate) {
        if (graph.containsKey(coordinate)) {
            return graph.get(coordinate);
        }
        
        return null;
    }

    public Coordinate getStart() {
        return start;
    }
}
