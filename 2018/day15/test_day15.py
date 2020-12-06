"""
Tests for Advent of Code Day 15.

https://adventofcode.com/2018/day/15
"""

from os import path
from .day15 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    '#######',
    '#.G...#',
    '#...EG#',
    '#.#.#G#',
    '#..G#E#',
    '#.....#',
    '#######',
]
_TEST_DATA_2 = [
    '#######',
    '#E..EG#',
    '#.#G.E#',
    '#E.##E#',
    '#G..#.#',
    '#..E#.#',
    '#######',
]


def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 27730
    assert run_part1(_TEST_DATA_2) == 39514

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 179968


def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA) == 4988
    assert run_part2(_TEST_DATA_2) == 31284

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 42098
