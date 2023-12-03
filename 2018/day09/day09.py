"""
Implementation for Advent of Code Day 9.

https://adventofcode.com/2018/day/9
"""

from __future__ import print_function

import re
from collections import defaultdict

_INSTRUCTION_REGEX = re.compile(
    r'^(\d+) players; last marble is worth (\d+) points$')


class DoublyLinkedNode:
    """Represents a doubly-linked list node."""

    def __init__(self, score):
        self.score = score
        self.prev_node = None
        self.next_node = None

    def link_to_self(self):
        """Link the node to itself to form a circular list of length 1"""
        assert self.prev_node is None and self.next_node is None
        self.prev_node = self
        self.next_node = self

    def insert_after(self, other):
        """Insert the provided node after the current node and return the new node."""
        other.prev_node = self
        other.next_node = self.next_node
        self.next_node.prev_node = other
        self.next_node = other

        return other

    def walk(self, count):
        """Walk from the current node a given number of steps (either positive or negative)."""
        current = self
        if count < 0:
            for _ in range(abs(count)):
                current = current.prev_node
        else:
            for _ in range(count):
                current = current.next_node

        return current

    def remove(self):
        """Remove the current node from the list and return the next node."""
        self.prev_node.next_node = self.next_node
        self.next_node.prev_node = self.prev_node
        return self.next_node


def _parse_instruction(instruction):
    match = _INSTRUCTION_REGEX.match(instruction)
    if not match:
        raise ValueError('Invalid instruction')

    return (int(match.group(1)), int(match.group(2)))


def _play_the_game(player_count, last_move_score):
    players = defaultdict(int)
    root_node = DoublyLinkedNode(0)
    root_node.link_to_self()

    current_node = root_node
    current_player = 0

    for current_score in range(1, last_move_score + 1):
        if current_score % 23 == 0:
            current_node = current_node.walk(-7)
            players[current_player] += current_score + current_node.score
            current_node = current_node.remove()
        else:
            current_node = current_node.walk(1)
            current_node = current_node.insert_after(
                DoublyLinkedNode(current_score))

        current_player = (current_player + 1) % player_count

    return max(players.values())


def run_part1(file_content):
    """Implmentation for Part 1."""
    (player_count, last_move_score) = _parse_instruction(file_content)
    return _play_the_game(player_count, last_move_score)


def run_part2(file_content):
    """Implmentation for Part 2."""
    (player_count, last_move_score) = _parse_instruction(file_content)
    return _play_the_game(player_count, last_move_score * 100)


if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.read().strip()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
