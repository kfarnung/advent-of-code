/*
 * Test cases for Advent of Code 2023 day 6.
 */
package com.github.kfarnung.adventofcode.aoc2023.day06;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    private final List<String> input = List.of(
            "Time:      7  15   30",
            "Distance:  9  40  200");

    @Test
    void testGetPart1() throws IOException {
        assertEquals("288", App.getPart1(input));

        List<String> realInput = readLinesFromResources(this, "day06.txt");
        assertEquals("4403592", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        assertEquals("71503", App.getPart2(input));

        List<String> realInput = readLinesFromResources(this, "day06.txt");
        assertEquals("38017587", App.getPart2(realInput));
    }
}
