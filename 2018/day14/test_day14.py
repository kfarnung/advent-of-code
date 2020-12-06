"""
Tests for Advent of Code Day 14.

https://adventofcode.com/2018/day/14
"""

from .day14 import run_part1, run_part2


def test_part1():
    """Tests for Part 1."""
    assert run_part1(9) == '5158916779'
    assert run_part1(290431) == '1776718175'


def test_part2():
    """Tests for Part 2."""
    assert run_part2(51589) == 9
    assert run_part2(290431) == 20220949
