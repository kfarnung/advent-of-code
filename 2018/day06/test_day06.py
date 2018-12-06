"""
Tests for Advent of Code Day 6.

https://adventofcode.com/2018/day/6
"""

from os import path
from .day06 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    '1, 1',
    '1, 6',
    '8, 3',
    '3, 4',
    '5, 5',
    '8, 9',
]

def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 17

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 3251

def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA, 32) == 16

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content, 10000) == 47841
