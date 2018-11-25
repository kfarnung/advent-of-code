"""
Tests for Advent of Code Day 1.

https://adventofcode.com/2018/day/1
"""

from .day01 import calculate_frequency, calculate_frequency_two_match

def test_part1():
    """Tests for Part 1."""
    assert calculate_frequency(['+1', '-2', '+3', '+1']) == 3
    assert calculate_frequency(['+1', '+1', '+1']) == 3
    assert calculate_frequency(['+1', '+1', '-2']) == 0
    assert calculate_frequency(['-1', '-2', '-3']) == -6

def test_part2():
    """Tests for Part 2."""
    assert calculate_frequency_two_match(['+1', '-2', '+3', '+1']) == 2
    assert calculate_frequency_two_match(['+1', '-1']) == 0
    assert calculate_frequency_two_match(['+3', '+3', '+4', '-2', '-4']) == 10
    assert calculate_frequency_two_match(['-6', '+3', '+8', '+5', '-6']) == 5
    assert calculate_frequency_two_match(['+7', '+7', '-2', '-7', '-4']) == 14
