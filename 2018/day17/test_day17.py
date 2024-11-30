"""
Tests for Advent of Code Day 17.

https://adventofcode.com/2018/day/17
"""

import sys
from os import path
from .day17 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_INPUT_FILE = path.join(
    path.dirname(path.dirname(_CURRENT_FILE_DIR)),
    'private',
    'inputs',
    '2018',
    path.basename(_CURRENT_FILE_DIR) + '.txt')
_TEST_DATA = [
    'x=495, y=2..7',
    'y=7, x=495..501',
    'x=501, y=3..7',
    'x=498, y=2..4',
    'x=506, y=1..2',
    'x=498, y=10..13',
    'x=504, y=10..13',
    'y=13, x=498..504',
]
_TEST_DATA_2 = [
    'x=485, y=2..4',
    'x=490, y=5..16',
    'y=16, x=490..513',
    'x=513, y=4..16',
    'x=499, y=7..10',
    'y=7, x=499..508',
    'y=10, x=499..508',
    'x=508, y=7..10',
]


def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 57
    assert run_part1(_TEST_DATA_2) == 240

    sys.setrecursionlimit(10000)
    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 29063


def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA) == 29
    assert run_part2(_TEST_DATA_2) == 202

    sys.setrecursionlimit(10000)
    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 23811
