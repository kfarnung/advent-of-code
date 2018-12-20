"""
Tests for Advent of Code Day 19.

https://adventofcode.com/2018/day/19
"""

from os import path
from .day19 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    '#ip 0',
    'seti 5 0 1',
    'seti 6 0 2',
    'addi 0 1 0',
    'addr 1 2 3',
    'setr 1 0 0',
    'seti 8 0 4',
    'seti 9 0 5',
]

def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 6

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 2016

def test_part2():
    """Tests for Part 2."""
    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 22674960
