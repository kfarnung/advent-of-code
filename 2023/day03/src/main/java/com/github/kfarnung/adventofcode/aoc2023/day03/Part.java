package com.github.kfarnung.adventofcode.aoc2023.day03;

public class Part {
    private final char symbol;
    private final int x;
    private final int y;

    public Part(char symbol, int x, int y) {
        this.symbol = symbol;
        this.x = x;
        this.y = y;
    }

    public char getSymbol() {
        return symbol;
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }

    @Override
    public int hashCode() {
        return x ^ y;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj instanceof Part) {
            Part other = (Part) obj;
            return x == other.x && y == other.y;
        }

        return false;
    }
}
