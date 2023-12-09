/*
 * Test cases for Advent of Code 2023 day 9.
 */
package com.github.kfarnung.adventofcode.aoc2023.day09;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    private final List<String> input = List.of(
        "0 3 6 9 12 15",
        "1 3 6 10 15 21",
        "10 13 16 21 30 45"
    );

    @Test
    void testGetPart1() throws IOException {
        assertEquals("114", App.getPart1(input));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("1684566095", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        assertEquals("2", App.getPart2(input));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("1136", App.getPart2(realInput));
    }
}
