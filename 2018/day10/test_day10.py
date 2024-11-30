"""
Tests for Advent of Code Day 10.

https://adventofcode.com/2018/day/10
"""

from os import path
from .day10 import find_message

_CURRENT_FILE_DIR = path.dirname(__file__)
_INPUT_FILE = path.join(
    path.dirname(path.dirname(_CURRENT_FILE_DIR)),
    'private',
    'inputs',
    '2018',
    path.basename(_CURRENT_FILE_DIR) + '.txt')
_TEST_DATA = [
    'position=< 9,  1> velocity=< 0,  2>',
    'position=< 7,  0> velocity=<-1,  0>',
    'position=< 3, -2> velocity=<-1,  1>',
    'position=< 6, 10> velocity=<-2, -1>',
    'position=< 2, -4> velocity=< 2,  2>',
    'position=<-6, 10> velocity=< 2, -2>',
    'position=< 1,  8> velocity=< 1, -1>',
    'position=< 1,  7> velocity=< 1,  0>',
    'position=<-3, 11> velocity=< 1, -2>',
    'position=< 7,  6> velocity=<-1, -1>',
    'position=<-2,  3> velocity=< 1,  0>',
    'position=<-4,  3> velocity=< 2,  0>',
    'position=<10, -3> velocity=<-1,  1>',
    'position=< 5, 11> velocity=< 1, -2>',
    'position=< 4,  7> velocity=< 0, -1>',
    'position=< 8, -2> velocity=< 0,  1>',
    'position=<15,  0> velocity=<-2,  0>',
    'position=< 1,  6> velocity=< 1,  0>',
    'position=< 8,  9> velocity=< 0, -1>',
    'position=< 3,  3> velocity=<-1,  1>',
    'position=< 0,  5> velocity=< 0, -1>',
    'position=<-2,  2> velocity=< 2,  0>',
    'position=< 5, -2> velocity=< 1,  2>',
    'position=< 1,  4> velocity=< 2,  1>',
    'position=<-2,  7> velocity=< 2, -2>',
    'position=< 3,  6> velocity=<-1, -1>',
    'position=< 5,  0> velocity=< 1,  0>',
    'position=<-6,  0> velocity=< 2,  0>',
    'position=< 5,  9> velocity=< 1, -2>',
    'position=<14,  7> velocity=<-2,  0>',
    'position=<-3,  6> velocity=< 2, -1>',
]
_TEST_EXPECTED = (
    '#   #  ###\n'
    '#   #   # \n'
    '#   #   # \n'
    '#####   # \n'
    '#   #   # \n'
    '#   #   # \n'
    '#   #   # \n'
    '#   #  ###'
)
_INPUT_EXPECTED = (
    '######  ######   ####   #####    ####    ####    ####      ###\n'
    '     #       #  #    #  #    #  #    #  #    #  #    #      # \n'
    '     #       #  #       #    #  #       #       #           # \n'
    '    #       #   #       #    #  #       #       #           # \n'
    '   #       #    #       #####   #       #       #           # \n'
    '  #       #     #       #    #  #  ###  #  ###  #           # \n'
    ' #       #      #       #    #  #    #  #    #  #           # \n'
    '#       #       #       #    #  #    #  #    #  #       #   # \n'
    '#       #       #    #  #    #  #   ##  #   ##  #    #  #   # \n'
    '######  ######   ####   #####    ### #   ### #   ####    ###  '
)


def test_find_message():
    """Tests for Part 1 """
    message, seconds = find_message(_TEST_DATA)
    assert message == _TEST_EXPECTED
    assert seconds == 3

    with open(_INPUT_FILE, 'r') as input_file:
        message, seconds = find_message(input_file.readlines())
        assert message == _INPUT_EXPECTED
        assert seconds == 10886
