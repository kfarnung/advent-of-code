"""
Implementation for Advent of Code Day 6.

https://adventofcode.com/2018/day/6
"""

from __future__ import print_function

import sys
from functools import reduce

_MIN_INT = -sys.maxsize - 1

class Point:
    """Represents a point in 2D space."""
    @staticmethod
    def from_string(input_str):
        """Creates a point from a string of the form 'x, y'"""
        [coord_x, coord_y] = [int(part) for part in input_str.split(', ')]
        return Point(coord_x, coord_y)

    def __init__(self, coord_x, coord_y):
        self.coord_x = coord_x
        self.coord_y = coord_y

    def __str__(self):
        return '({}, {})'.format(self.coord_x, self.coord_y)

    def manhattan_distance(self, other):
        """Calculates the Manhattan distance between the two points"""
        return abs(other.coord_x - self.coord_x) + abs(other.coord_y - self.coord_y)

class Rect:
    """Represents a rectangle in 2D space."""
    @staticmethod
    def bounding_box(points):
        """Returns a rectangle which forms the boundary of all points"""
        min_x = sys.maxsize
        min_y = sys.maxsize
        max_x = _MIN_INT
        max_y = _MIN_INT

        for point in points:
            min_x = min(min_x, point.coord_x)
            min_y = min(min_y, point.coord_y)
            max_x = max(max_x, point.coord_x)
            max_y = max(max_y, point.coord_y)

        return Rect(Point(min_x, min_y), Point(max_x, max_y))

    def __init__(self, upper_left, lower_right):
        self.upper_left = upper_left
        self.lower_right = lower_right

    def __str__(self):
        return '[{}, {}]'.format(self.upper_left, self.lower_right)

    def get_range_x(self):
        """Gets the range of x-coordinates for the rectangle"""
        return range(self.upper_left.coord_x, self.lower_right.coord_x + 1)

    def get_range_y(self):
        """Gets the range of y-coordinates for the rectangle"""
        return range(self.upper_left.coord_y, self.lower_right.coord_y + 1)

    def point_on_boundary(self, point):
        """Checks if the point is on the boundary of the rectangle"""
        return (point.coord_x == self.upper_left.coord_x or
                point.coord_x == self.lower_right.coord_x or
                point.coord_y == self.upper_left.coord_y or
                point.coord_y == self.lower_right.coord_y)

def _distance_to_all(points, point_to_check):
    return reduce(
        lambda prev, current: prev + current.manhattan_distance(point_to_check),
        points,
        0
    )

def _find_closest_point(points, point_to_check):
    distance_map = {}

    for point in points:
        distance = point.manhattan_distance(point_to_check)
        if distance in distance_map:
            distance_map[distance] = None
        else:
            distance_map[distance] = point

    return distance_map[min(distance_map)]

def _map_closest_points(points, boundary):
    area_map = {}

    for coord_x in boundary.get_range_x():
        for coord_y in boundary.get_range_y():
            point_to_check = Point(coord_x, coord_y)
            area_map[point_to_check] = _find_closest_point(points, point_to_check)

    return area_map

def _calculate_point_area(area_map, boundary, point):
    total_area = 0

    for key, value in area_map.items():
        if value == point:
            if boundary.point_on_boundary(key):
                # Assume boundary points would have led to infinity
                return 0

            total_area += 1

    return total_area

def _calculate_max_area(points):
    points = list(points)
    boundary = Rect.bounding_box(points)
    area_map = _map_closest_points(points, boundary)

    return max(_calculate_point_area(area_map, boundary, point) for point in points)

def _find_points_within_distance(points, distance):
    points = list(points)
    area_set = set()
    boundary = Rect.bounding_box(points)

    for coord_x in boundary.get_range_x():
        for coord_y in boundary.get_range_y():
            point_to_check = Point(coord_x, coord_y)
            if _distance_to_all(points, point_to_check) < distance:
                area_set.add(point_to_check)

    return area_set

def run_part1(file_content):
    """Implmentation for Part 1."""
    return _calculate_max_area(Point.from_string(input_str) for input_str in file_content)

def run_part2(file_content, distance):
    """Implmentation for Part 2."""
    return len(_find_points_within_distance(
        (Point.from_string(input_str) for input_str in file_content),
        distance))

if __name__ == "__main__":
    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content, 10000)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
