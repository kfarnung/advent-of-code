package com.github.kfarnung.adventofcode.aoc2023.day02;

import java.util.ArrayList;
import java.util.List;

public class Game {
    private final int id;
    private final List<Subset> subsets;

    public Game(int id, List<Subset> subsets) {
        this.id = id;
        this.subsets = subsets;
    }

    public static Game parse(String line) {
        String[] gameSplit = line.split(":");
        int id = Integer.parseInt(gameSplit[0].substring(5).trim());
        String[] subsetSplit = gameSplit[1].split(";");
        List<Subset> subsets = new ArrayList<>();
        for (String subset : subsetSplit) {
            String[] colorSplit = subset.split(",");
            int red = 0;
            int green = 0;
            int blue = 0;
            for (String color : colorSplit) {
                String[] colorCount = color.trim().split(" ");
                int count = Integer.parseInt(colorCount[0]);
                switch (colorCount[1]) {
                    case "red":
                        red = count;
                        break;
                    case "green":
                        green = count;
                        break;
                    case "blue":
                        blue = count;
                        break;
                    default:
                        throw new IllegalArgumentException("Unknown color: " + colorCount[1]);
                }
            }
            subsets.add(new Subset(red, green, blue));
        }
        return new Game(id, subsets);
    }

    public int getId() {
        return id;
    }

    public List<Subset> getSubsets() {
        return subsets;
    }

    public boolean isValid(Subset target) {
        for (Subset subset : subsets) {
            if (!subset.isValid(target)) {
                return false;
            }
        }
        return true;
    }

    public Subset getMinimum() {
        int red = 0;
        int green = 0;
        int blue = 0;

        for (Subset subset : subsets) {
            red = Math.max(red, subset.getRed());
            green = Math.max(green, subset.getGreen());
            blue = Math.max(blue, subset.getBlue());
        }
        return new Subset(red, green, blue);
    }
}
