"""
Tests for Advent of Code Day 25.

https://adventofcode.com/2018/day/25
"""

from os import path
from .day25 import run_part1

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    '0,0,0,0',
    '3,0,0,0',
    '0,3,0,0',
    '0,0,3,0',
    '0,0,0,3',
    '0,0,0,6',
    '9,0,0,0',
    '12,0,0,0',
]
_TEST_DATA2 = [
    '-1,2,2,0',
    '0,0,2,-2',
    '0,0,0,-2',
    '-1,2,0,0',
    '-2,-2,-2,2',
    '3,0,2,-1',
    '-1,3,2,2',
    '-1,0,-1,0',
    '0,2,1,-2',
    '3,0,0,0',
]
_TEST_DATA3 = [
    '1,-1,0,1',
    '2,0,-1,0',
    '3,2,-1,0',
    '0,0,3,1',
    '0,0,-1,-1',
    '2,3,-2,0',
    '-2,2,0,0',
    '2,-2,0,-1',
    '1,-1,0,-1',
    '3,2,0,2',
]
_TEST_DATA4 = [
    '1,-1,-1,-2',
    '-2,-2,0,1',
    '0,2,1,3',
    '-2,3,-2,1',
    '0,2,3,-2',
    '-1,-1,1,-2',
    '0,-2,-1,0',
    '-2,2,3,-1',
    '1,2,2,0',
    '-1,-2,0,-2',
]

def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 2
    assert run_part1(_TEST_DATA2) == 4
    assert run_part1(_TEST_DATA3) == 3
    assert run_part1(_TEST_DATA4) == 8

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 370
