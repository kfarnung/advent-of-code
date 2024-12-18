"""
Tests for Advent of Code Day 2.

https://adventofcode.com/2018/day/2
"""

from os import path
from .day02 import calculate_checksum, find_correct_id

_CURRENT_FILE_DIR = path.dirname(__file__)
_INPUT_FILE = path.join(
    path.dirname(path.dirname(_CURRENT_FILE_DIR)),
    'private',
    'inputs',
    '2018',
    path.basename(_CURRENT_FILE_DIR) + '.txt')


def test_part1():
    """Tests for Part 1."""
    assert calculate_checksum([
        'abcdef',
        'bababc',
        'abbcde',
        'abcccd',
        'aabcdd',
        'abcdee',
        'ababab',
    ]) == 12

    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert calculate_checksum(file_content) == 9633


def test_part2():
    """Tests for Part 2."""
    assert find_correct_id([
        'abcde',
        'fghij',
        'klmno',
        'pqrst',
        'fguij',
        'axcye',
        'wvxyz',
    ]) == 'fgij'

    with open(_INPUT_FILE, 'r') as input_file:
        file_content = input_file.readlines()
        assert find_correct_id(file_content) == 'lujnogabetpmsydyfcovzixaw'
