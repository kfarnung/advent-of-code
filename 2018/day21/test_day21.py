"""
Tests for Advent of Code Day 21.

https://adventofcode.com/2018/day/21
"""

from os import path
from .day21 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)

def test_part1():
    """Tests for Part 1."""
    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 6619857

def test_part2():
    """Tests for Part 2."""
    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 9547924
