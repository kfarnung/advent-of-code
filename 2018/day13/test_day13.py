"""
Tests for Advent of Code Day 13.

https://adventofcode.com/2018/day/13
"""
from os import path
from .day13 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    '/->-\\        \n',
    '|   |  /----\\\n',
    '| /-+--+-\\  |\n',
    '| | |  | v  |\n',
    '\\-+-/  \\-+--/\n',
    '  \\------/   \n',
]

_TEST_DATA_2 = [
    '/>-<\\  ',
    '|   |  ',
    '| /<+-\\',
    '| | | v',
    '\\>+</ |',
    '  |   ^',
    '  \\<->/',
]

def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == (7, 3)

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == (111, 13)

def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA_2) == (6, 4)

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == (16, 73)
