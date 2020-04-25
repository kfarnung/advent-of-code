"""
Implementation for Advent of Code Day 23.

https://adventofcode.com/2018/day/23
"""

from __future__ import print_function

import heapq
import re
from operator import attrgetter

_POSITIONS_REGEX = re.compile(r'^pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(\d+)$')

class Point3D(object):
    """Represents a point in 3D space."""
    def __init__(self, coord_x, coord_y, coord_z):
        self.coord_x = coord_x
        self.coord_y = coord_y
        self.coord_z = coord_z

    def __eq__(self, other):
        return (
            self.coord_x == other.coord_x and
            self.coord_y == other.coord_y and
            self.coord_z == other.coord_z
        )

    def __le__(self, other):
        return (
            self.coord_x <= other.coord_x and
            self.coord_y <= other.coord_y and
            self.coord_z <= other.coord_z
        )

    def manhattan_distance(self, other):
        """Calculates the distance between two points."""
        return (
            abs(other.coord_x - self.coord_x) +
            abs(other.coord_y - self.coord_y) +
            abs(other.coord_z - self.coord_z)
        )

    def min_coords(self, other):
        """Returns the minimum coordinates of both points."""
        return Point3D(
            min(self.coord_x, other.coord_x),
            min(self.coord_y, other.coord_y),
            min(self.coord_z, other.coord_z)
        )

    def max_coords(self, other):
        """Returns the maximum coordinates of both points."""
        return Point3D(
            max(self.coord_x, other.coord_x),
            max(self.coord_y, other.coord_y),
            max(self.coord_z, other.coord_z)
        )

    def clamp(self, min_coord, max_coord):
        """Returns a new point clamped to the range of the given points."""
        return Point3D(
            min(max(self.coord_x, min_coord.coord_x), max_coord.coord_x),
            min(max(self.coord_y, min_coord.coord_y), max_coord.coord_y),
            min(max(self.coord_z, min_coord.coord_z), max_coord.coord_z)
        )

class Rect3D(object):
    """Represents a rectangle in 3D space."""
    def __init__(self, negative_corner, positive_corner):
        self.negative_corner = negative_corner
        self.positive_corner = positive_corner

    def __eq__(self, other):
        return (
            self.negative_corner == other.negative_corner and
            self.positive_corner == other.positive_corner
        )

    def split(self):
        """Splits the rectangle into smaller pieces."""
        neg = self.negative_corner
        pos = self.positive_corner
        step_x = (pos.coord_x - neg.coord_x) // 2
        step_y = (pos.coord_y - neg.coord_y) // 2
        step_z = (pos.coord_z - neg.coord_z) // 2

        for coord_x in xrange(neg.coord_x, pos.coord_x + 1, step_x + 1):
            for coord_y in xrange(neg.coord_y, pos.coord_y + 1, step_y + 1):
                for coord_z in xrange(neg.coord_z, pos.coord_z + 1, step_z + 1):
                    max_corner = Point3D(coord_x + step_x, coord_y + step_y, coord_z + step_z)
                    yield Rect3D(
                        Point3D(coord_x, coord_y, coord_z),
                        max_corner.min_coords(pos)
                    )

    def volume(self):
        """Calculates the volume enclosed by the rectangle."""
        neg = self.negative_corner
        pos = self.positive_corner
        return (
            (pos.coord_x - neg.coord_x + 1) *
            (pos.coord_y - neg.coord_y + 1) *
            (pos.coord_z - neg.coord_z + 1)
        )

    def closest_within_rect(self, point):
        """Returns the closest point within the rect to the given point."""
        return point.clamp(self.negative_corner, self.positive_corner)

    @staticmethod
    def bounding_box(rects):
        """Creates a bounding box for the input rectangles."""
        negative_corner = None
        positive_corner = None

        for rect in rects:
            negative_corner = (
                rect.negative_corner.min_coords(negative_corner)
                if negative_corner
                else rect.negative_corner
            )
            positive_corner = (
                rect.positive_corner.max_coords(positive_corner)
                if positive_corner
                else rect.positive_corner
            )

        return Rect3D(negative_corner, positive_corner)

class NanoBot(object):
    """Represents a single nanobot."""
    def __init__(self, position, signal_radius):
        self.position = position
        self.signal_radius = signal_radius

    def bounding_box(self):
        """Gets the bounding box that contains the bot's range."""
        return Rect3D(
            Point3D(
                self.position.coord_x - self.signal_radius,
                self.position.coord_y - self.signal_radius,
                self.position.coord_z - self.signal_radius
            ),
            Point3D(
                self.position.coord_x + self.signal_radius,
                self.position.coord_y + self.signal_radius,
                self.position.coord_z + self.signal_radius
            )
        )

    def within_range(self, other):
        """Check if another bot is within range."""
        return self.point_within_range(other.position)

    def point_within_range(self, point):
        """Check if a point is within range."""
        return self.position.manhattan_distance(point) <= self.signal_radius

    def intersects(self, rect):
        """Check if the bot intersects the rectangle."""
        # Find the closest point within the rectangle and check if it's within range.
        rect_point = rect.closest_within_rect(self.position)
        return self.point_within_range(rect_point)

    @staticmethod
    def from_string(data):
        """Create a nanobot from input data."""
        match = _POSITIONS_REGEX.match(data)
        if not match:
            raise Exception('Invalid nanobot data')

        position = Point3D(
            int(match.group(1)),
            int(match.group(2)),
            int(match.group(3))
        )

        return NanoBot(position, int(match.group(4)))

def run_part1(file_content):
    """Implmentation for Part 1."""
    bot_list = [NanoBot.from_string(line) for line in file_content if line]
    largest_bot = max(bot_list, key=attrgetter('signal_radius'))
    within_range = [bot for bot in bot_list if largest_bot.within_range(bot)]

    return len(within_range)

def run_part2(file_content):
    """Implmentation for Part 2."""
    bot_list = [NanoBot.from_string(line) for line in file_content if line]

    # Create the box that contains all bot coverage areas
    rect = Rect3D.bounding_box(bot.bounding_box() for bot in bot_list)

    origin = Point3D(0, 0, 0)
    original_len = len(bot_list)

    priority_queue = []
    heapq.heappush(
        priority_queue,
        (
            original_len - len(bot_list),
            origin.manhattan_distance(rect.closest_within_rect(origin)),
            bot_list,
            rect
        )
    )

    while priority_queue:
        current = heapq.heappop(priority_queue)

        if current[3].volume() == 1:
            # We've found the smallest possible box and it's already the best option because of the
            # priority queue.
            return current[1]

        for split in current[3].split():
            new_list = [bot for bot in current[2] if bot.intersects(split)]
            if new_list:
                heapq.heappush(
                    priority_queue,
                    (
                        original_len - len(new_list),
                        origin.manhattan_distance(split.closest_within_rect(origin)),
                        new_list,
                        split
                    )
                )

    return None

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

    run(sys.argv[1])
