package com.github.kfarnung.adventofcode.aoc2023.day02;

public class Subset {
    private final int red;
    private final int green;
    private final int blue;

    public Subset(int red, int green, int blue) {
        this.red = red;
        this.green = green;
        this.blue = blue;
    }

    public int getRed() {
        return red;
    }

    public int getGreen() {
        return green;
    }

    public int getBlue() {
        return blue;
    }

    public boolean isValid(Subset target) {
        return red <= target.getRed() && green <= target.getGreen() && blue <= target.getBlue();
    }

    public int getPower() {
        return red * green * blue;
    }
}
