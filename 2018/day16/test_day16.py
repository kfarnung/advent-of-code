"""
Tests for Advent of Code Day 16.

https://adventofcode.com/2018/day/16
"""

from os import path
from .day16 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_INPUT_FILE = path.join(
    path.dirname(path.dirname(_CURRENT_FILE_DIR)),
    'private',
    'inputs',
    '2018',
    path.basename(_CURRENT_FILE_DIR) + '.txt')


def test_part1():
    """Tests for Part 1."""
    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 517


def test_part2():
    """Tests for Part 2."""
    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 667
