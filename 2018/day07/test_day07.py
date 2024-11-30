"""
Tests for Advent of Code Day 7.

https://adventofcode.com/2018/day/7
"""

from os import path
from .day07 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_INPUT_FILE = path.join(
    path.dirname(path.dirname(_CURRENT_FILE_DIR)),
    'private',
    'inputs',
    '2018',
    path.basename(_CURRENT_FILE_DIR) + '.txt')
_TEST_DATA = [
    'Step C must be finished before step A can begin.',
    'Step C must be finished before step F can begin.',
    'Step A must be finished before step B can begin.',
    'Step A must be finished before step D can begin.',
    'Step B must be finished before step E can begin.',
    'Step D must be finished before step E can begin.',
    'Step F must be finished before step E can begin.',
]


def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 'CABDFE'

    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 'GLMVWXZDKOUCEJRHFAPITSBQNY'


def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA, 2, 0) == 15

    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content, 5, 60) == 1105
