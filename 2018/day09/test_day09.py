"""
Tests for Advent of Code Day 9.

https://adventofcode.com/2018/day/9
"""

from os import path
from .day09 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_INPUT_FILE = path.join(
    path.dirname(path.dirname(_CURRENT_FILE_DIR)),
    'private',
    'inputs',
    '2018',
    path.basename(_CURRENT_FILE_DIR) + '.txt')
_TEST_DATA = [
    ('9 players; last marble is worth 25 points', 32),
    ('10 players; last marble is worth 1618 points', 8317),
    ('13 players; last marble is worth 7999 points', 146373),
    ('17 players; last marble is worth 1104 points', 2764),
    ('21 players; last marble is worth 6111 points', 54718),
    ('30 players; last marble is worth 5807 points', 37305),
]


def test_part1():
    """Tests for Part 1."""
    for (instruction, high_score) in _TEST_DATA:
        assert run_part1(instruction) == high_score

    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part1(file_content) == 399745


def test_part2():
    """Tests for Part 2."""
    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part2(file_content) == 3349098263
