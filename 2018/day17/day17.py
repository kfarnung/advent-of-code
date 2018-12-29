"""
Implementation for Advent of Code Day 17.

https://adventofcode.com/2018/day/17
"""

import re
from collections import defaultdict
from functools import reduce
from itertools import repeat
from operator import itemgetter

_SPAN_REGEX = re.compile(r'^([xy])=(\d+), [xy]=(\d+)..(\d+)$')

class Grid:
    """Represents a grid of water, clay, and sand."""
    def __init__(self, clay_positions):
        self.ground = defaultdict(repeat('.').__next__)
        for clay in clay_positions:
            self.ground[clay] = '#'

        self.min_y = min(self.ground.keys(), key=itemgetter(1))[1]
        self.max_y = max(self.ground.keys(), key=itemgetter(1))[1]

    def __str__(self):
        min_x = min(self.ground.keys(), key=itemgetter(0))[0]
        max_x = max(self.ground.keys(), key=itemgetter(0))[0]

        return '\n'.join(''.join(self.ground[(coord_x, coord_y)]
                                 for coord_x in range(min_x, max_x + 1))
                         for coord_y in range(self.min_y, self.max_y + 1))

    def fill(self):
        """Fill the grid with water."""
        self._flow((500, self.min_y))

    def get_water_count(self):
        """Get the total number of water tiles."""
        return reduce(
            lambda prev, current: prev + 1 if current in ('~', '|') else prev,
            self.ground.values(),
            0
        )

    def get_settled_count(self):
        """Get the total number of settled water tiles."""
        return reduce(
            lambda prev, current: prev + 1 if current == '~' else prev,
            self.ground.values(),
            0
        )

    def _flow(self, position):
        self.ground[position] = '|'

        if position[1] == self.max_y:
            return

        down = (position[0], position[1] + 1)
        down_tile = self.ground[down]
        if down_tile == '|':
            # Combining with other flowing water.
            return

        if down_tile == '.':
            self._flow(down)

        # Try to fill left and right
        left = position
        while self._can_flow_horizontal(left):
            self.ground[left] = '|'
            left = (left[0] - 1, left[1])

        right = position
        while self._can_flow_horizontal(right):
            self.ground[right] = '|'
            right = (right[0] + 1, right[1])

        if self.ground[left] == '#' and self.ground[right] == '#':
            self._mark_settled(left, right)
            return

        if self.ground[left] == '.':
            self._flow(left)

        if self.ground[right] == '.':
            self._flow(right)

    def _mark_settled(self, left, right):
        assert left[1] == right[1]

        for coord_x in range(left[0] + 1, right[0]):
            self.ground[(coord_x, left[1])] = '~'

    def _can_flow_horizontal(self, position):
        down = (position[0], position[1] + 1)
        current_tile = self.ground[position]
        if current_tile not in ('.', '|'):
            return False

        down_tile = self.ground[down]
        if down_tile not in ('#', '~'):
            return False

        return True

def _parse_clay(file_content):
    clay = []
    for line in file_content:
        match = _SPAN_REGEX.match(line)
        if not match:
            raise Exception('Invalid line')

        flipped = match.group(1) == 'y'
        first = int(match.group(2))
        clay += [(first, second) if not flipped else (second, first)
                 for second in range(int(match.group(3)), int(match.group(4)) + 1)]

    return clay

def run_part1(file_content):
    """Implmentation for Part 1."""
    grid = Grid(_parse_clay(file_content))
    grid.fill()
    return grid.get_water_count()

def run_part2(file_content):
    """Implmentation for Part 2."""
    grid = Grid(_parse_clay(file_content))
    grid.fill()
    return grid.get_settled_count()

if __name__ == "__main__":
    import sys

    def run(argv1):
        """The main function."""
        with open(argv1, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    # It would be better to use a separate stack given the depth required.
    sys.setrecursionlimit(10000)
    run(sys.argv[1])
