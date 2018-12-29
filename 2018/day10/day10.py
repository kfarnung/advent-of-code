"""
Implementation for Advent of Code Day 10.

https://adventofcode.com/2018/day/10
"""

import re

_ENTRY_REGEX = re.compile(r'^position=< *(-?\d+), *(-?\d+)> velocity=< *(-?\d+), *(-?\d+)>$')

class Point2D(object):
    """Represents a point in 2D space."""
    def __init__(self, coord_x, coord_y):
        self.coord_x = coord_x
        self.coord_y = coord_y

    def __eq__(self, other):
        return self.coord_x == other.coord_x and self.coord_y == other.coord_y

    def __hash__(self):
        return hash((self.coord_x, self.coord_y))

    def __ne__(self, other):
        return not self.__eq__(other)

    def add(self, other):
        """Adds the given point to the current instance."""
        self.coord_x += other.coord_x
        self.coord_y += other.coord_y

    def subtract(self, other):
        """Subtracts the given point from the current instance."""
        self.coord_x -= other.coord_x
        self.coord_y -= other.coord_y

class LightPoint(object):
    """Represents a moving point of light in the sky."""
    def __init__(self, position, velocity):
        self.current_position = position
        self.velocity = velocity

    def step_backward(self):
        """Apply the inverse velocity to step backward in time."""
        self.current_position.subtract(self.velocity)

    def step_forward(self):
        """Apply the velocity to step forward in time."""
        self.current_position.add(self.velocity)

class LightPointSystem(object):
    """Represents a system of light points in the sky."""
    def __init__(self, points):
        self.points = list(points)

    def get_area(self):
        """Gets the area of the rectange that bounds the points."""
        upper_left, lower_right = self._get_bounds()
        length_x = lower_right.coord_x - upper_left.coord_x
        length_y = lower_right.coord_y - upper_left.coord_y
        return length_x * length_y

    def plot_points(self):
        """Plot the points in coordinate space."""
        upper_left, lower_right = self._get_bounds()
        points = set(point.current_position for point in self.points)
        lines = []

        for coord_y in xrange(upper_left.coord_y, lower_right.coord_y + 1):
            line = []

            for coord_x in xrange(upper_left.coord_x, lower_right.coord_x + 1):
                point = Point2D(coord_x, coord_y)
                line.append('#' if point in points else ' ')

            lines.append(''.join(line))

        return '\n'.join(lines)

    def step_backward(self):
        """Step the system of points backward in time."""
        for point in self.points:
            point.step_backward()

    def step_forward(self):
        """Step the system of points forward in time."""
        for point in self.points:
            point.step_forward()

    def _get_bounds(self):
        """Gets the rectange that bounds the points."""
        min_x = self.points[0].current_position.coord_x
        min_y = self.points[0].current_position.coord_y
        max_x = self.points[0].current_position.coord_x
        max_y = self.points[0].current_position.coord_y

        for point in self.points:
            min_x = min(min_x, point.current_position.coord_x)
            min_y = min(min_y, point.current_position.coord_y)
            max_x = max(max_x, point.current_position.coord_x)
            max_y = max(max_y, point.current_position.coord_y)

        return Point2D(min_x, min_y), Point2D(max_x, max_y)

def _parse_entry(input_str):
    """Parse the entry that contains the starting point and velocity."""
    match = _ENTRY_REGEX.match(input_str)
    if not match:
        raise Exception('Invalid input string')

    return LightPoint(
        Point2D(int(match.group(1)), int(match.group(2))),
        Point2D(int(match.group(3)), int(match.group(4))),
    )

def find_message(file_content):
    """Implmentation for Part 1."""
    point_system = LightPointSystem(_parse_entry(entry) for entry in file_content)
    total_area = point_system.get_area()
    seconds = 0

    while True:
        point_system.step_forward()
        new_area = point_system.get_area()
        seconds += 1

        if new_area > total_area:
            point_system.step_backward()
            seconds -= 1
            break

        total_area = new_area

    return point_system.plot_points(), seconds

if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            message, seconds = find_message(input_file.readlines())
            print "Part 1:"
            print message
            print "Part 2: {}".format(seconds)

    if len(sys.argv) < 2:
        print "Usage: python {} <input>".format(sys.argv[0])
        sys.exit(1)

    run(sys.argv[1])
