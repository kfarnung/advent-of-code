"""
Tests for Advent of Code Day 11.

https://adventofcode.com/2018/day/11
"""

from .day11 import get_power_level, run_part1, run_part2


def test_power_level():
    """Tests for the _get_power_level function."""
    assert get_power_level(57, 122, 79) == -5
    assert get_power_level(39, 217, 196) == 0
    assert get_power_level(71, 101, 153) == 4


def test_part1():
    """Tests for Part 1."""
    assert run_part1(18) == (33, 45)
    assert run_part1(42) == (21, 61)
    assert run_part1(5535) == (19, 41)


def test_part2():
    """Tests for Part 2."""
    assert run_part2(5535) == ((237, 284), 11)
