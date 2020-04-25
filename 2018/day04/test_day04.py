"""
Tests for Advent of Code Day 4.

https://adventofcode.com/2018/day/4
"""

from os import path
from .day04 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    '[1518-11-01 00:05] falls asleep',
    '[1518-11-03 00:29] wakes up',
    '[1518-11-04 00:36] falls asleep',
    '[1518-11-01 00:00] Guard #10 begins shift',
    '[1518-11-05 00:03] Guard #99 begins shift',
    '[1518-11-01 23:58] Guard #99 begins shift',
    '[1518-11-04 00:02] Guard #99 begins shift',
    '[1518-11-01 00:55] wakes up',
    '[1518-11-03 00:05] Guard #10 begins shift',
    '[1518-11-01 00:30] falls asleep',
    '[1518-11-02 00:50] wakes up',
    '[1518-11-05 00:55] wakes up',
    '[1518-11-01 00:25] wakes up',
    '[1518-11-04 00:46] wakes up',
    '[1518-11-02 00:40] falls asleep',
    '[1518-11-05 00:45] falls asleep',
    '[1518-11-03 00:24] falls asleep',
]


def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 240

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 73646


def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA) == 4455

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 4727
