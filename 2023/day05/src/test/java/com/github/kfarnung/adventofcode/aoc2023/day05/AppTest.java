/*
 * Test cases for Advent of Code 2023 day 5.
 */
package com.github.kfarnung.adventofcode.aoc2023.day05;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    private final List<String> input = List.of(
            "seeds: 79 14 55 13",
            "",
            "seed-to-soil map:",
            "50 98 2",
            "52 50 48",
            "",
            "soil-to-fertilizer map:",
            "0 15 37",
            "37 52 2",
            "39 0 15",
            "",
            "fertilizer-to-water map:",
            "49 53 8",
            "0 11 42",
            "42 0 7",
            "57 7 4",
            "",
            "water-to-light map:",
            "88 18 7",
            "18 25 70",
            "",
            "light-to-temperature map:",
            "45 77 23",
            "81 45 19",
            "68 64 13",
            "",
            "temperature-to-humidity map:",
            "0 69 1",
            "1 0 69",
            "",
            "humidity-to-location map:",
            "60 56 37",
            "56 93 4");

    @Test
    void testGetPart1() throws IOException {
        assertEquals("35", App.getPart1(input));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("510109797", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        assertEquals("46", App.getPart2(input));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("9622622", App.getPart2(realInput));
    }
}
