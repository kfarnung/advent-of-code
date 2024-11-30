/*
 * Test cases for Advent of Code 2023 day 1.
 */
package com.github.kfarnung.adventofcode.aoc2023.day01;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    @Test
    void testGetPart1() throws IOException {
        List<String> testInput = List.of(
                "1abc2",
                "pqr3stu8vwx",
                "a1b2c3d4e5f",
                "treb7uchet");
        assertEquals("142", App.getPart1(testInput));

        List<String> realInput = readLinesFromResources(this, "day01.txt");
        assertEquals("54697", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        List<String> testInput = List.of(
                "two1nine",
                "eightwothree",
                "abcone2threexyz",
                "xtwone3four",
                "4nineeightseven2",
                "zoneight234",
                "7pqrstsixteen");
        assertEquals("281", App.getPart2(testInput));

        List<String> realInput = readLinesFromResources(this, "day01.txt");
        assertEquals("54885", App.getPart2(realInput));
    }
}
