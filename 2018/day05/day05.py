"""
Implementation for Advent of Code Day 5.

https://adventofcode.com/2018/day/5
"""

import string

def _does_react(unit1, unit2):
    return abs(ord(unit1) - ord(unit2)) == 32

def _reduce_polymer(polymer):
    characters = list(polymer)

    # Going backwards through the list only requires one pass
    for index in reversed(xrange(len(characters) - 1)):
        if index < len(characters) and _does_react(characters[index - 1], characters[index]):
            del characters[index]
            del characters[index - 1]

    return characters

def _remove_unit_type(polymer, unit_type):
    unit_type = unit_type.lower()
    return [x for x in polymer if x.lower() != unit_type]

def run_part1(file_content):
    """Implmentation for Part 1."""
    return len(_reduce_polymer(file_content))

def run_part2(file_content):
    """Implmentation for Part 2."""
    size = len(file_content)

    for unit_type in string.ascii_lowercase:
        reduced = _remove_unit_type(file_content, unit_type)
        reduced = _reduce_polymer(reduced)
        size = min(size, len(reduced))

    return size

if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.read().strip()
            print "Part 1: {}".format(run_part1(file_content))
            print "Part 2: {}".format(run_part2(file_content))

    if len(sys.argv) < 2:
        print "Usage: python {} <input>".format(sys.argv[0])
        sys.exit(1)

    run(sys.argv[1])
