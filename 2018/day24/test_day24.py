"""
Tests for Advent of Code Day 24.

https://adventofcode.com/2018/day/24
"""

from os import path
from .day24 import run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    'Immune System:',
    '17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does '
    '4507 fire damage at initiative 2',
    '989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an '
    'attack that does 25 slashing damage at initiative 3',
    '',
    'Infection:',
    '801 units each with 4706 hit points (weak to radiation) with an attack that does 116 '
    'bludgeoning damage at initiative 1',
    '4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack '
    'that does 12 slashing damage at initiative 4',
]

def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 5216

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 20150

def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA) == 51

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 13005
