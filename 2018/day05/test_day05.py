"""
Tests for Advent of Code Day 5.

https://adventofcode.com/2018/day/5
"""

from .day05 import run_part1, run_part2

def test_part1():
    """Tests for Part 1."""
    assert run_part1('aA') == 0
    assert run_part1('abBA') == 0
    assert run_part1('abAB') == 4
    assert run_part1('aabAAB') == 6
    assert run_part1('dabAcCaCBAcCcaDA') == 10

    # Some edge cases to verify the iteration boundaries
    assert run_part1('dabAcCaCBAcCcaDAa') == 9
    assert run_part1('DdabAcCaCBAcCcaDAa') == 8

def test_part2():
    """Tests for Part 2."""
    assert run_part2('dabAcCaCBAcCcaDA') == 4
