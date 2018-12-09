"""
Implementation for Advent of Code Day 9.

https://adventofcode.com/2018/day/9
"""

import re
from collections import defaultdict

_INSTRUCTION_REGEX = re.compile(r'^(\d+) players; last marble is worth (\d+) points$')

class _DoublyLinkedNode(object):
    def __init__(self, score):
        self.score = score
        self.prev = None
        self.next = None

    def link_to_self(self):
        """Link the node to itself to form a circular list of length 1"""
        assert self.prev is None and self.next is None
        self.prev = self
        self.next = self

    def insert_after(self, other):
        """Insert the provided node after the current node and return the new node."""
        other.prev = self
        other.next = self.next
        self.next.prev = other
        self.next = other

        return other

    def walk(self, count):
        """Walk from the current node a given number of steps (either positive or negative)."""
        current = self
        if count < 0:
            for _ in xrange(abs(count)):
                current = current.prev
        else:
            for _ in xrange(count):
                current = current.next

        return current

    def remove(self):
        """Remove the current node from the list and return the next node."""
        self.prev.next = self.next
        self.next.prev = self.prev
        return self.next

def _parse_instruction(instruction):
    match = _INSTRUCTION_REGEX.match(instruction)
    if not match:
        raise Exception('Invalid instruction')

    return (int(match.group(1)), int(match.group(2)))

def _play_the_game(player_count, last_move_score):
    players = defaultdict(int)
    root_node = _DoublyLinkedNode(0)
    root_node.link_to_self()

    current_node = root_node
    current_player = 0

    for current_score in xrange(1, last_move_score + 1):
        if current_score % 23 == 0:
            current_node = current_node.walk(-7)
            players[current_player] += current_score + current_node.score
            current_node = current_node.remove()
        else:
            current_node = current_node.walk(1)
            current_node = current_node.insert_after(_DoublyLinkedNode(current_score))

        current_player = (current_player + 1) % player_count

    return max(players.itervalues())

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
            print "Part 1: {}".format(run_part1(file_content))
            print "Part 2: {}".format(run_part2(file_content))

    if len(sys.argv) < 2:
        print "Usage: python {} <input>".format(sys.argv[0])
        sys.exit(1)

    run(sys.argv[1])
