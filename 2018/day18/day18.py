"""
Implementation for Advent of Code Day 18.

https://adventofcode.com/2018/day/18
"""

from collections import Counter
from operator import itemgetter

class LumberArea(object):
    """Represents an area for lumber collection."""
    def __init__(self, file_content):
        self.area = {}
        for row_index, row in enumerate(file_content):
            for cell_index, cell in enumerate(row):
                if cell == '\n':
                    continue

                self.area[(row_index, cell_index)] = cell

        self.max_rows = max(self.area.iterkeys(), key=itemgetter(0))[0]
        self.max_cells = max(self.area.iterkeys(), key=itemgetter(1))[1]

    def __str__(self):
        lines = []
        for row_index in xrange(self.max_rows + 1):
            line = []
            for cell_index in xrange(self.max_cells + 1):
                line.append(self.area[(row_index, cell_index)])
            lines.append(''.join(line))

        return '\n'.join(lines)

    def get_counts(self):
        """Gets the counts of the acre types."""
        return Counter(self.area.itervalues())

    def execute_minute(self):
        """Executes a single minute of growth."""
        area_copy = self.area.copy()
        for row_index in xrange(self.max_rows + 1):
            for cell_index in xrange(self.max_cells + 1):
                position = (row_index, cell_index)
                cell = self.area[position]
                if cell == '.':
                    if self._has_neighbors(position, '|', 3):
                        area_copy[position] = '|'
                elif cell == '|':
                    if self._has_neighbors(position, '#', 3):
                        area_copy[position] = '#'
                elif cell == '#':
                    if not (self._has_neighbors(position, '#', 1) and
                            self._has_neighbors(position, '|', 1)):
                        area_copy[position] = '.'
                else:
                    raise Exception('Unexpected character')

        self.area = area_copy

    def _has_neighbors(self, position, neighbor_type, at_least):
        search_directions = [
            (-1, -1),
            (-1, 0),
            (-1, 1),
            (0, 1),
            (1, 1),
            (1, 0),
            (1, -1),
            (0, -1),
        ]

        for direction in search_directions:
            cell_position = (position[0] + direction[0], position[1] + direction[1])
            if cell_position in self.area:
                cell = self.area[cell_position]
                if cell == neighbor_type:
                    at_least -= 1

        return at_least <= 0

def run_part1(file_content):
    """Implmentation for Part 1."""
    area = LumberArea(file_content)

    for _ in xrange(10):
        area.execute_minute()

    counts = area.get_counts()
    return counts['|'] * counts['#']

def run_part2(file_content):
    """Implmentation for Part 2."""
    area = LumberArea(file_content)
    seen_patterns = set()
    cycle_index = None
    cycle_patterns = []

    for minute in xrange(1000000000):
        area.execute_minute()

        pattern = str(area)
        if not cycle_patterns:
            if pattern in seen_patterns:
                cycle_patterns.append(pattern)
                cycle_index = minute

            seen_patterns.add(pattern)
        else:
            if pattern != cycle_patterns[0]:
                cycle_patterns.append(pattern)
            else:
                break

    index = (1000000000 - 1 - cycle_index) % len(cycle_patterns)
    counts = Counter(cycle_patterns[index])
    return counts['|'] * counts['#']

if __name__ == "__main__":
    import sys

    def run(argv1):
        """The main function."""
        with open(argv1, 'r') as input_file:
            file_content = input_file.readlines()
            print "Part 1: {}".format(run_part1(file_content))
            print "Part 2: {}".format(run_part2(file_content))

    if len(sys.argv) < 2:
        print "Usage: python {} <input>".format(sys.argv[0])
        sys.exit(1)

    run(sys.argv[1])
