"""
Implementation for Advent of Code Day 2.

https://adventofcode.com/2018/day/2
"""

from __future__ import print_function

from collections import Counter, defaultdict
from functools import reduce


def _count_duplicate_letters(string):
    """Returns the counts of duplicate letters in a string."""
    return set(Counter(string).values())


def _remove_differences(str1, str2):
    """Returns a new string which contains only the matching characters of the input strings."""
    assert len(str1) == len(str2)

    return reduce(
        lambda prev, pair: prev + pair[0] if pair[0] == pair[1] else prev,
        zip(str1, str2),
        ''
    )


def calculate_checksum(inputs):
    """Calculates the checksum for a given set of inputs."""
    duplicate_counts = defaultdict(int)
    for row in inputs:
        for count in _count_duplicate_letters(row):
            duplicate_counts[count] += 1

    return duplicate_counts[2] * duplicate_counts[3]


def find_correct_id(inputs):
    """Finds the pair of IDs which differ by only one character."""
    for input1 in inputs:
        for input2 in inputs:
            matching = _remove_differences(input1, input2)
            if len(matching) == len(input1) - 1:
                return matching.strip()

    return None


if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(calculate_checksum(file_content)))
            print("Part 2: {}".format(find_correct_id(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
