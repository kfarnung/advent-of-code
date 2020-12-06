"""
Tests for Advent of Code Day 5.

https://adventofcode.com/2018/day/5
"""

from os import path
from .day05 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)


def test_part1():
    """Tests for Part 1."""
    assert run_part1('aA') == 0
    assert run_part1('abBA') == 0
    assert run_part1('abAB') == 4
    assert run_part1('aabAAB') == 6
    assert run_part1('dabAcCaCBAcCcaDA') == 10

    # Some edge cases to verify the iteration boundaries
    assert run_part1('dabAcCaCBAcCcaDAa') == 9
    assert run_part1('DdabAcCaCBAcCcaDAa') == 8

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part1(file_content) == 11894


def test_part2():
    """Tests for Part 2."""
    assert run_part2('dabAcCaCBAcCcaDA') == 4

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part2(file_content) == 5310
