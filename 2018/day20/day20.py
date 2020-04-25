"""
Implementation for Advent of Code Day 20.

https://adventofcode.com/2018/day/20
"""

from __future__ import print_function

from collections import defaultdict, deque
from operator import itemgetter
from sys import maxsize


class TreeNode:
    """Represents a single node in the Regex tree."""

    def __init__(self, direction):
        self.direction = direction
        self.children = []

    def connect(self, node):
        """Connects a given node to the current node."""
        self.children.append(node)
        return node


class Regex:
    """Represents an input 'regex' string."""

    def __init__(self, input_str):
        root_stack = []
        path_stack = []
        path_list = []
        current_index = 0
        current_root = TreeNode('^')
        current_node = current_root

        if input_str[current_index] != '^':
            raise Exception('Invalid pattern')

        current_index += 1
        while current_index < len(input_str):
            current = input_str[current_index]
            if current in ('N', 'S', 'E', 'W'):
                current_node = current_node.connect(TreeNode(current))
            elif current == '(':
                current_node = current_node.connect(TreeNode(current))
                root_stack.append(current_root)
                path_stack.append(path_list)
                path_list = []
                current_root = current_node
            elif current == '|':
                path_list.append(current_node)
                current_node = current_root
            elif current == ')':
                path_list.append(current_node)
                current_node = TreeNode(current)
                for path in path_list:
                    path.connect(current_node)
                path_list = path_stack.pop()
                current_root = root_stack.pop()
            elif current == '$':
                current_node = current_node.connect(TreeNode(current))
                break
            else:
                raise Exception('Invalid character')
            current_index += 1

        if root_stack:
            raise Exception('Malformed input')

        self.root = current_root

    def __str__(self):
        result, end_node = Regex._stringify_grouping(self.root)
        assert end_node is None

        result[-1] = '$'
        return ''.join(result)

    @staticmethod
    def _stringify_grouping(grouping_start):
        assert grouping_start.direction == '(' or grouping_start.direction == '^'
        result = [grouping_start.direction]
        end_node = None

        for child in grouping_start.children:
            current_node = child
            while current_node:
                if current_node.direction == '(':
                    grouping_result, grouping_end = Regex._stringify_grouping(
                        current_node)
                    result += grouping_result

                    assert len(grouping_end.children) == 1
                    current_node = grouping_end.children[0]
                elif current_node.direction == ')':
                    result.append('|')

                    assert end_node is None or current_node == end_node
                    end_node = current_node
                    current_node = None
                else:
                    result.append(current_node.direction)
                    assert len(current_node.children) <= 1
                    current_node = current_node.children[0] if current_node.children else None

        result[-1] = ')'
        return result, end_node


class FacilityMap:
    """Represents a map of the facility as specified by a Regex."""
    _directions = {
        'N': (-1, 0),
        'E': (0, 1),
        'S': (1, 0),
        'W': (0, -1),
    }

    def __init__(self):
        self.grid = {}

    def __str__(self):
        min_x = min(self.grid, key=itemgetter(0))[0] - 1
        max_x = max(self.grid, key=itemgetter(0))[0] + 1
        min_y = min(self.grid, key=itemgetter(1))[1] - 1
        max_y = max(self.grid, key=itemgetter(1))[1] + 1

        result = []

        for row in range(min_x, max_x + 1):
            line = []
            for col in range(min_y, max_y + 1):
                position = (row, col)
                line.append(self.grid[position]
                            if position in self.grid else '#')
            result.append(''.join(line))

        return '\n'.join(result)

    def follow_path(self, root_node):
        """Follows a path through the facility as specified by the starting node."""
        queue = deque()
        distance_map = defaultdict(lambda: maxsize)
        visitors = defaultdict(set)
        queue.append(((0, 0), 0, root_node))
        self.grid[(0, 0)] = 'X'

        while queue:
            current_item = queue.popleft()
            current_position, current_distance, current_node = current_item

            for child in current_node.children:
                if child in visitors[current_position]:
                    continue

                visitors[current_position].add(child)
                direction = FacilityMap._get_direction(child.direction)
                if direction == (0, 0):
                    queue.append((current_position, current_distance, child))
                else:
                    new_distance = current_distance + 1
                    door_position = FacilityMap._add_direction(
                        current_position, direction)
                    self.grid[door_position] = FacilityMap._get_door_type(
                        child.direction)
                    next_position = FacilityMap._add_direction(
                        door_position, direction)
                    self.grid[next_position] = '.'
                    distance_map[next_position] = min(
                        distance_map[next_position], new_distance)
                    queue.append((next_position, new_distance, child))

        return (
            max(distance_map.values()),
            sum(1 for _ in (
                distance for distance in distance_map.values() if distance >= 1000))
        )

    @staticmethod
    def _get_door_type(direction):
        if direction in ('N', 'S'):
            return '-'

        if direction in ('E', 'W'):
            return '|'

        raise Exception('Invalid direction')

    @staticmethod
    def _get_direction(direction):
        if direction in FacilityMap._directions:
            return FacilityMap._directions[direction]

        return (0, 0)

    @staticmethod
    def _add_direction(current_position, direction):
        return (current_position[0] + direction[0], current_position[1] + direction[1])

    @staticmethod
    def _get_next_node(current_position, current_distance, next_node):
        if next_node.direction == 'N':
            current_position = (current_position[0] - 2, current_position[1])
            current_distance += 1
        elif next_node.direction == 'E':
            current_position = (current_position[0], current_position[1] + 2)
            current_distance += 1
        elif next_node.direction == 'S':
            current_position = (current_position[0] + 2, current_position[1])
            current_distance += 1
        elif next_node.direction == 'W':
            current_position = (current_position[0], current_position[1] - 2)
            current_distance += 1

        return (current_position, current_distance, next_node)


def run_part1(file_content):
    """Implmentation for Part 1."""
    regex = Regex(file_content)
    facility = FacilityMap()
    distance, _ = facility.follow_path(regex.root)
    return distance


def run_part2(file_content):
    """Implmentation for Part 2."""
    regex = Regex(file_content)
    facility = FacilityMap()
    _, count = facility.follow_path(regex.root)
    return count


if __name__ == "__main__":
    import sys

    def run(argv1):
        """The main function."""
        with open(argv1, 'r') as input_file:
            file_content = input_file.read().strip()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
