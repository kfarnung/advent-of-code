/*
 * Test cases for Advent of Code 2023 day 3.
 */
package com.github.kfarnung.adventofcode.aoc2023.day03;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    @Test
    void testGetPart1() throws IOException {
        List<String> lines = Arrays.asList(
                "467..114..",
                "...*......",
                "..35..633.",
                "......#...",
                "617*......",
                ".....+.58.",
                "..592.....",
                "......755.",
                "...$.*....",
                ".664.598..");
        assertEquals("4361", App.getPart1(lines));

        lines = readLinesFromResources(this, "input.txt");
        assertEquals("512794", App.getPart1(lines));
    }

    @Test
    void testGetPart2() throws IOException {
        List<String> lines = Arrays.asList(
                "467..114..",
                "...*......",
                "..35..633.",
                "......#...",
                "617*......",
                ".....+.58.",
                "..592.....",
                "......755.",
                "...$.*....",
                ".664.598..");
        assertEquals("467835", App.getPart2(lines));

        lines = readLinesFromResources(this, "input.txt");
        assertEquals("67779080", App.getPart2(lines));
    }
}
