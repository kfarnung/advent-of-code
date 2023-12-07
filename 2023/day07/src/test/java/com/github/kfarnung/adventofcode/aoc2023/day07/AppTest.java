/*
 * Test cases for Advent of Code 2023 day 7.
 */
package com.github.kfarnung.adventofcode.aoc2023.day07;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

class AppTest {
    private final List<String> input = List.of(
        "32T3K 765",
        "T55J5 684",
        "KK677 28",
        "KTJJT 220",
        "QQQJA 483"
    );

    @Test
    void testGetPart1() throws IOException {
        assertEquals("6440", App.getPart1(input));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("253910319", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        assertEquals("5905", App.getPart2(input));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("254083736", App.getPart2(realInput));
    }
}
