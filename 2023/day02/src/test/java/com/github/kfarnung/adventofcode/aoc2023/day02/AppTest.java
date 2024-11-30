/*
 * Test cases for Advent of Code 2023 day 2.
 */
package com.github.kfarnung.adventofcode.aoc2023.day02;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    private final List<String> testInput = List.of(
            "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
            "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
            "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
            "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
            "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green");
    @Test
    void testGetPart1() throws IOException {
        assertEquals("8", App.getPart1(testInput));

        List<String> realInput = readLinesFromResources(this, "day02.txt");
        assertEquals("2913", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        assertEquals("2286", App.getPart2(testInput));

        List<String> realInput = readLinesFromResources(this, "day02.txt");
        assertEquals("55593", App.getPart2(realInput));
    }
}
