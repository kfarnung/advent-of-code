"""
Tests for Advent of Code Day 22.

https://adventofcode.com/2018/day/22
"""

from os import path
from .day22 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_INPUT_FILE = path.join(
    path.dirname(path.dirname(_CURRENT_FILE_DIR)),
    'private',
    'inputs',
    '2018',
    path.basename(_CURRENT_FILE_DIR) + '.txt')
_TEST_DATA = [
    'depth: 510',
    'target: 10,10',
]


def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 114

    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 10115


def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA) == 45

    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 990
