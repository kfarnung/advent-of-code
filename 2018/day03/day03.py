"""
Implementation for Advent of Code Day 3.

https://adventofcode.com/2018/day/3
"""

import re
from collections import defaultdict
from functools import reduce

CLAIM_REGEX = re.compile(r"^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$")


def _parse_claim(claim):
    """Parse a given claim and return a tuple containing the results."""
    match = CLAIM_REGEX.match(claim)
    assert match is not None

    return (
        int(match.group(1)),
        int(match.group(2)),
        int(match.group(3)),
        int(match.group(4)),
        int(match.group(5)),
    )


def _add_claim(claim_map, parsed_claim):
    """Claims the desired cells in the map."""
    for pos_x in range(parsed_claim[1], parsed_claim[1] + parsed_claim[3]):
        for pos_y in range(parsed_claim[2], parsed_claim[2] + parsed_claim[4]):
            key = "{},{}".format(pos_x, pos_y)
            claim_map[key] += 1


def _create_claim_map(parsed_claims):
    """Creates the map with the number of claims for each cell."""
    claim_map = defaultdict(int)
    for claim in parsed_claims:
        _add_claim(claim_map, claim)

    return claim_map


def _check_unique_claim(claim_map, parsed_claim):
    """Checks if all cells for a given claim were only claimed once."""
    for pos_x in range(parsed_claim[1], parsed_claim[1] + parsed_claim[3]):
        for pos_y in range(parsed_claim[2], parsed_claim[2] + parsed_claim[4]):
            key = "{},{}".format(pos_x, pos_y)
            if claim_map[key] > 1:
                return False

    return True


def run_part1(inputs):
    """Implmentation for Part 1."""
    parsed_claims = [_parse_claim(claim) for claim in inputs]

    return reduce(
        lambda area, position: area + 1 if position > 1 else area,
        _create_claim_map(parsed_claims).values(),
        0
    )


def run_part2(inputs):
    """Implmentation for Part 2."""
    parsed_claims = [_parse_claim(claim) for claim in inputs]
    claim_map = _create_claim_map(parsed_claims)

    for claim in parsed_claims:
        result = _check_unique_claim(claim_map, claim)
        if result:
            return claim[0]

    return None


if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
