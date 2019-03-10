"""
Tests for Advent of Code Day 23.

https://adventofcode.com/2018/day/23
"""

from os import path
from .day23 import Point3D, Rect3D, run_part1, run_part2

_CURRENT_FILE_DIR = path.dirname(__file__)
_TEST_DATA = [
    'pos=<0,0,0>, r=4',
    'pos=<1,0,0>, r=1',
    'pos=<4,0,0>, r=3',
    'pos=<0,2,0>, r=1',
    'pos=<0,5,0>, r=3',
    'pos=<0,0,3>, r=1',
    'pos=<1,1,1>, r=1',
    'pos=<1,1,2>, r=1',
    'pos=<1,3,1>, r=1',
]
_TEST_DATA_2 = [
    'pos=<10,12,12>, r=2',
    'pos=<12,14,12>, r=2',
    'pos=<16,12,12>, r=4',
    'pos=<14,14,14>, r=6',
    'pos=<50,50,50>, r=200',
    'pos=<10,10,10>, r=5',
]

def test_closest_within_rect():
    """Tests for closest point."""
    rect = Rect3D(Point3D(3, 4, 5), Point3D(7, 9, 14))

    # Test the "corner" cases
    assert rect.closest_within_rect(Point3D(0, 0, 0)) == Point3D(3, 4, 5)
    assert rect.closest_within_rect(Point3D(0, 0, 16)) == Point3D(3, 4, 14)
    assert rect.closest_within_rect(Point3D(8, 10, 15)) == Point3D(7, 9, 14)

    # Test the "edge" cases
    assert rect.closest_within_rect(Point3D(5, 4, 5)) == Point3D(5, 4, 5)

    # Test contained cases
    assert rect.closest_within_rect(Point3D(5, 6, 8)) == Point3D(5, 6, 8)

def test_part1():
    """Tests for Part 1."""
    assert run_part1(_TEST_DATA) == 7

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part1(file_content) == 780

def test_part2():
    """Tests for Part 2."""
    assert run_part2(_TEST_DATA_2) == 36

    with open(path.join(_CURRENT_FILE_DIR, 'input'), 'r') as input_file:
        file_content = input_file.readlines()
        assert run_part2(file_content) == 110841112
