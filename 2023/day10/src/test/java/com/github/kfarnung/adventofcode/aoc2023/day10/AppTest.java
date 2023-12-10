/*
 * Test cases for Advent of Code 2023 day 10.
 */
package com.github.kfarnung.adventofcode.aoc2023.day10;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    private final List<String> input = List.of(
        "-L|F7",
        "7S-7|",
        "L|7||",
        "-L-J|",
        "L|-JF"
    );
    private final List<String> input2 = List.of(
        "7-F7-",
        ".FJ|7",
        "SJLL7",
        "|F--J",
        "LJ.LJ"
    );

    @Test
    void testGetPart1() throws IOException {
        assertEquals("4", App.getPart1(input));
        assertEquals("8", App.getPart1(input2));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("6786", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        assertEquals("0", App.getPart2(input));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("0", App.getPart2(realInput));
    }
}
