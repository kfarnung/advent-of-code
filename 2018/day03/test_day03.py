"""
Tests for Advent of Code Day 3.

https://adventofcode.com/2018/day/3
"""

from .day03 import run_part1, run_part2

def test_part1():
    """Tests for Part 1."""
    assert run_part1([
        '#1 @ 1,3: 4x4',
        '#2 @ 3,1: 4x4',
        '#3 @ 5,5: 2x2',
    ]) == 4

def test_part2():
    """Tests for Part 2."""
    assert run_part2([
        '#1 @ 1,3: 4x4',
        '#2 @ 3,1: 4x4',
        '#3 @ 5,5: 2x2',
    ]) == 3
