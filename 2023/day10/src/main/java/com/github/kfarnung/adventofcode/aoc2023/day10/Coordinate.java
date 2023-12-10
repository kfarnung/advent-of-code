package com.github.kfarnung.adventofcode.aoc2023.day10;

public class Coordinate {
    private final int row;
    private final int column;

    public Coordinate(int row, int column) {
        this.row = row;
        this.column = column;
    }

    public int getRow() {
        return row;
    }

    public int getColumn() {
        return column;
    }

    public Coordinate add(Direction direction) {
        return new Coordinate(row + direction.getRowOffset(), column + direction.getColumnOffset());
    }

    @Override
    public int hashCode() {
        return row ^ column;
    }

    @Override
    public boolean equals(Object obj) {
        if (!(obj instanceof Coordinate)) {
            return false;
        }

        Coordinate other = (Coordinate) obj;
        return row == other.row && column == other.column;
    }
}
