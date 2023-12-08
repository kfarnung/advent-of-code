package com.github.kfarnung.adventofcode.aoc2023.utilities;

public class MathUtils {
    public static long leastCommonMultiple(long a, long b) {
        return a * b / greatestCommonDivisor(a, b);
    }

    public static long greatestCommonDivisor(long a, long b) {
        if (b == 0) {
            return a;
        }

        return greatestCommonDivisor(b, a % b);
    }
}
