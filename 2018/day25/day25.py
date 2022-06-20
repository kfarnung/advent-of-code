"""
Implementation for Advent of Code Day 25.

https://adventofcode.com/2018/day/25
"""

from __future__ import print_function

from collections import deque


class Point4D:
    """Represents a point in 4D space."""

    def __init__(self, w, x, y, z):  # pylint: disable=invalid-name
        self.coord_w = w
        self.coord_x = x
        self.coord_y = y
        self.coord_z = z

    def __str__(self):
        return '({}, {}, {}, {})'.format(
            self.coord_w,
            self.coord_x,
            self.coord_y,
            self.coord_z
        )

    def manhattan_distance(self, other):
        """Calculates the distance between two points."""
        return (
            abs(other.coord_w - self.coord_w) +
            abs(other.coord_x - self.coord_x) +
            abs(other.coord_y - self.coord_y) +
            abs(other.coord_z - self.coord_z)
        )


class FixedPointInSpacetime:
    """Represents a fixed point in spacetime as defined by the puzzle."""

    def __init__(self, coord):
        self.coord = coord
        self.connections = set()

    def try_connect(self, other):
        """Attempt to connect to points together."""
        if self is not other:
            distance = self.coord.manhattan_distance(other.coord)
            if distance <= 3:
                assert distance > 0
                self.connections.add(other)
                other.connections.add(self)
                return True

        return False

    @staticmethod
    def parse(line):
        """Parse a set of comma-separated coordinates."""
        (coord_w, coord_x, coord_y, coord_z) = [
            int(coord)
            for coord in line.strip().split(',')
        ]

        point = Point4D(coord_w, coord_x, coord_y, coord_z)
        return FixedPointInSpacetime(point)


def _parse_lines(file_content):
    return [FixedPointInSpacetime.parse(line) for line in file_content]


def _cluster_points(points):
    for point in points:
        for other in points:
            point.try_connect(other)


def _count_constellations(points):
    count = 0
    seen = set()

    for point in points:
        if point in seen:
            continue

        count += 1
        search_queue = deque()
        search_queue.append(point)

        while search_queue:
            current = search_queue.popleft()
            seen.add(current)

            for child in current.connections:
                if child not in seen:
                    search_queue.append(child)

    return count


def run_part1(file_content):
    """Implmentation for Part 1."""
    points = _parse_lines(file_content)
    _cluster_points(points)
    return _count_constellations(points)


if __name__ == "__main__":
    import sys

    def run(argv1):
        """The main function."""
        with open(argv1, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(run_part1(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
