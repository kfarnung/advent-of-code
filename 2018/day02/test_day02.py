"""
Tests for Advent of Code Day 2.

https://adventofcode.com/2018/day/2
"""

from .day02 import calculate_checksum, find_correct_ids

def test_part1():
    """Tests for Part 1."""
    assert calculate_checksum([
        'abcdef',
        'bababc',
        'abbcde',
        'abcccd',
        'aabcdd',
        'abcdee',
        'ababab',
    ]) == 12

def test_part2():
    """Tests for Part 2."""
    assert find_correct_ids([
        'abcde',
        'fghij',
        'klmno',
        'pqrst',
        'fguij',
        'axcye',
        'wvxyz',
    ]) == 'fgij'
