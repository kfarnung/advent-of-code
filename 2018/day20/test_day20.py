"""
Tests for Advent of Code Day 20.

https://adventofcode.com/2018/day/20
"""

from os import path
from .day20 import Regex, run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    ("^WNE$", 3),
    ("^ENWWW(NEEE|SSE(EE|N))$", 10),
    ("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", 18),
    ("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 23),
    ("^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$", 31),
    ("^E(NN|S)E$", 4),
    ("^(N|S)N$", 2),
    ("^EEE(NN|SSS)EEE$", 9),
    ("^E(N|SS)EEE(E|SSS)$", 9),
]


def test_part1():
    """Tests for Part 1."""
    for test in _TEST_DATA:
        assert run_part1(test[0]) == test[1]

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part1(file_content) == 4155


def test_part2():
    """Tests for Part 2."""
    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.read().strip()
        assert run_part2(file_content) == 8434


def test_regex():
    """Tests for the Regex class."""
    for test in _TEST_DATA:
        assert str(Regex(test[0])) == test[0]

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.read().strip()
        assert str(Regex(file_content)) == file_content
