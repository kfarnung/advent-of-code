/*
 * Test cases for Advent of Code 2023 day 8.
 */
package com.github.kfarnung.adventofcode.aoc2023.day08;

import static com.github.kfarnung.adventofcode.aoc2023.utilities.InputUtils.readLinesFromResources;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.io.IOException;
import java.util.List;

import org.junit.jupiter.api.Test;

class AppTest {
    private final List<String> input = List.of(
            "RL",
            "",
            "AAA = (BBB, CCC)",
            "BBB = (DDD, EEE)",
            "CCC = (ZZZ, GGG)",
            "DDD = (DDD, DDD)",
            "EEE = (EEE, EEE)",
            "GGG = (GGG, GGG)",
            "ZZZ = (ZZZ, ZZZ)");
    private final List<String> input2 = List.of(
            "LLR",
            "",
            "AAA = (BBB, BBB)",
            "BBB = (AAA, ZZZ)",
            "ZZZ = (ZZZ, ZZZ)");
    private final List<String> input3 = List.of(
            "LR",
            "",
            "11A = (11B, XXX)",
            "11B = (XXX, 11Z)",
            "11Z = (11B, XXX)",
            "22A = (22B, XXX)",
            "22B = (22C, 22C)",
            "22C = (22Z, 22Z)",
            "22Z = (22B, 22B)",
            "XXX = (XXX, XXX)");

    @Test
    void testGetPart1() throws IOException {
        assertEquals("2", App.getPart1(input));
        assertEquals("6", App.getPart1(input2));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("13207", App.getPart1(realInput));
    }

    @Test
    void testGetPart2() throws IOException {
        assertEquals("6", App.getPart2(input3));

        List<String> realInput = readLinesFromResources(this, "input.txt");
        assertEquals("12324145107121", App.getPart2(realInput));
    }
}
