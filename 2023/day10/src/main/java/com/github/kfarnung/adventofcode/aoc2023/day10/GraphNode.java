package com.github.kfarnung.adventofcode.aoc2023.day10;

import java.util.EnumSet;
import java.util.HashSet;
import java.util.Set;

public class GraphNode {
    private final Coordinate coordinate;
    private final Set<Coordinate> neighbors;

    public GraphNode(Coordinate coordinate, char c) {
        this.coordinate = coordinate;
        this.neighbors = calculateNeighbors(coordinate, c);
    }

    public Coordinate getCoordinate() {
        return coordinate;
    }

    public Set<Coordinate> getNeighbors() {
        return neighbors;
    }

    private static Set<Coordinate> calculateNeighbors(Coordinate coordinate, char c) {
        Set<Coordinate> neighbors = new HashSet<>();
        EnumSet<Direction> directions = parseDirections(c);
        for (Direction direction : directions) {
            neighbors.add(coordinate.add(direction));
        }

        return neighbors;
    }
    
    private static EnumSet<Direction> parseDirections(char c) {
        switch (c) {
            case '|':
                return EnumSet.of(Direction.NORTH, Direction.SOUTH);
            case '-':
                return EnumSet.of(Direction.EAST, Direction.WEST);
            case 'L':
                return EnumSet.of(Direction.NORTH, Direction.EAST);
            case 'J':
                return EnumSet.of(Direction.NORTH, Direction.WEST);
            case '7':
                return EnumSet.of(Direction.SOUTH, Direction.WEST);
            case 'F':
                return EnumSet.of(Direction.SOUTH, Direction.EAST);
            case '.':
                return EnumSet.noneOf(Direction.class);
            case 'S':
                return EnumSet.allOf(Direction.class);
            default:
                throw new IllegalArgumentException("Unknown character: " + c);
        }
    }
}
