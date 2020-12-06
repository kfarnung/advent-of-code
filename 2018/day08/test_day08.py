"""
Tests for Advent of Code Day 8.

https://adventofcode.com/2018/day/8
"""

from os import path
from .day08 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = '2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2'


def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 138

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part1(file_content) == 42472


def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA) == 66

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part2(file_content) == 21810
