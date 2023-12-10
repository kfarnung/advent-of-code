package com.github.kfarnung.adventofcode.aoc2023.day10;

public class SearchState {
    private final Coordinate coordinate;
    private final int distance;

    public SearchState(Coordinate coordinate, int distance) {
        this.coordinate = coordinate;
        this.distance = distance;
    }

    public Coordinate getCoordinate() {
        return coordinate;
    }

    public int getDistance() {
        return distance;
    }
}
