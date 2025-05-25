"""
Implementation for Advent of Code Day 22.

https://adventofcode.com/2018/day/22
"""

import re
from collections import defaultdict
from operator import itemgetter
from sys import maxsize

_FIELD_REGEX = re.compile(r'^([a-z]+): (\d+)(?:,(\d+))?$')


class CaveSystem:
    """Represents a system of caves."""
    _NEIGHBORS = [
        (0, -1),
        (1, 0),
        (0, 1),
        (-1, 0)
    ]

    def __init__(self, depth, target):
        self.erosion_levels = {}
        self.depth = depth
        self.target = target

    def calculate_risk_level(self):
        """Calculate the total risk level of the cave area."""
        risk_level = 0

        for row in range(self.target[1] + 1):
            for col in range(self.target[0] + 1):
                risk_level += self._get_region_type((col, row))

        return risk_level

    def calculate_shortest_time(self):
        """Calculate the shortest time to reach the target position."""
        to_visit = set()
        shortest_times = defaultdict(lambda: maxsize)

        to_visit.add(((0, 0), 1, 0))

        while to_visit:
            current = min(
                to_visit,
                key=itemgetter(2)
            )
            to_visit.remove(current)

            current_position, current_item, current_time = current

            if current_position == self.target:
                if current_item == 1:
                    return current_time

                return current_time + 7

            current_x, current_y = current_position
            target_x, target_y = self.target
            if current_x > (target_x * 4) or current_y > (target_y * 4):
                continue

            st_key = (current_position, current_item)
            if shortest_times[st_key] < current_time:
                continue

            shortest_times[st_key] = current_time

            for neighbor in CaveSystem._get_neighbors(current_position):
                next_time = current_time + 1
                next_item = current_item

                if not self._is_item_valid(neighbor, next_item):
                    next_item = self._get_intersecting_item(
                        current_position, neighbor)
                    next_time += 7

                to_visit.add((neighbor, next_item, next_time))

        return None

    def _is_item_valid(self, position, item):
        region_type = self._get_region_type(position)
        if region_type == 0:
            # Torch or climbing gear
            return item in (1, 2)

        if region_type == 1:
            # Nothing or climbing gear
            return item in (0, 2)

        if region_type == 2:
            # Nothing or torch
            return item in (0, 1)

        raise ValueError('Invalid region type')

    def _get_intersecting_item(self, current_position, next_position):
        for item in range(3):
            if (self._is_item_valid(current_position, item) and
                    self._is_item_valid(next_position, item)):
                return item

        raise ValueError('Invalid combination of types')

    def _get_region_type(self, position):
        return self._get_erosion_level(position) % 3

    def _get_erosion_level(self, position):
        if position in self.erosion_levels:
            return self.erosion_levels[position]

        stack = [position]
        while stack:
            current_position = stack.pop()
            current_x, current_y = current_position
            target_x, target_y = self.target

            if ((current_x == 0 and current_y == 0) or
                    (current_x == target_x and current_y == target_y)):
                self.erosion_levels[current_position] = self._calculate_erosion_level(
                    0)
            elif current_y == 0:
                self.erosion_levels[current_position] = self._calculate_erosion_level(
                    current_x * 16807)
            elif current_x == 0:
                self.erosion_levels[current_position] = self._calculate_erosion_level(
                    current_y * 48271)
            else:
                left_one = (current_x - 1, current_y)
                up_one = (current_x, current_y - 1)

                has_left = left_one in self.erosion_levels
                has_up = up_one in self.erosion_levels
                if has_left and has_up:
                    self.erosion_levels[current_position] = self._calculate_erosion_level(
                        self.erosion_levels[(current_x - 1, current_y)] *
                        self.erosion_levels[(current_x, current_y - 1)])
                else:
                    stack.append(current_position)
                    if has_left not in self.erosion_levels:
                        stack.append(left_one)
                    if has_up not in self.erosion_levels:
                        stack.append(up_one)

        return self.erosion_levels[position]

    def _calculate_erosion_level(self, geologic_index):
        return (geologic_index + self.depth) % 20183

    @staticmethod
    def _get_neighbors(position):
        for neighbor in CaveSystem._NEIGHBORS:
            current_position = (
                position[0] + neighbor[0], position[1] + neighbor[1])
            if current_position[0] < 0 or current_position[1] < 0:
                continue

            yield current_position


def _parse_fields(lines):
    fields = {}
    for line in lines:
        match = _FIELD_REGEX.match(line)
        if match:
            value1 = int(match.group(2))
            value2 = match.group(3)
            fields[match.group(1)] = (
                value1, int(value2)) if value2 else value1

    return fields


def run_part1(file_content):
    """Implmentation for Part 1."""
    fields = _parse_fields(file_content)
    cave = CaveSystem(fields['depth'], fields['target'])
    return cave.calculate_risk_level()


def run_part2(file_content):
    """Implmentation for Part 2."""
    fields = _parse_fields(file_content)
    cave = CaveSystem(fields['depth'], fields['target'])
    return cave.calculate_shortest_time()


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
