"""
Implementation for Advent of Code Day 7.

https://adventofcode.com/2018/day/7
"""

from __future__ import print_function

import re
from collections import defaultdict

_INSTRUCTION_REGEX = re.compile(
    r'^Step ([A-Z]) must be finished before step ([A-Z]) can begin.$')


class Graph:
    """Represents a graph of nodes."""

    def __init__(self):
        self.nodes = set()
        self.dependency_map = defaultdict(set)

    def connect_nodes(self, parent, child):
        """Connect two nodes in the graph"""
        self.nodes.add(parent)
        self.nodes.add(child)
        self.dependency_map[child].add(parent)

    def get_next_ready_node(self, done, running=None):
        """Find the next node that's ready to execute"""
        ready = []

        if not running:
            running = set()

        for node in self.nodes - done - running:
            if not self.dependency_map[node] - done:
                ready.append(node)

        return sorted(ready)[0] if ready else None

    def get_node_order(self):
        """Get the order for the nodes in the graph"""
        result = []
        done = set()
        ready_node = self.get_next_ready_node(done)
        while ready_node:
            result.append(ready_node)
            done.add(ready_node)
            ready_node = self.get_next_ready_node(done)

        return result

    def get_remaining_nodes(self, done):
        """Get the set of nodes that haven't completed"""
        return self.nodes - done


class Factory:
    """Represents a factory with time and worker constraints."""

    def __init__(self, worker_count, base_time):
        self.current_time = 0
        self.worker_available_time = [0] * worker_count
        self.base_time = base_time

    def construct(self, graph):
        """Construct the graph according to the dependencies and build times"""
        node_available_time = {}

        while True:
            # Figure out what's currently done
            done = set(
                key
                for key, value in node_available_time.items()
                if value <= self.current_time
            )

            while True:
                running = set(node_available_time.keys())
                available_work = graph.get_next_ready_node(done, running)
                available_worker = self._get_available_worker()

                if available_work and available_worker is not None and available_worker >= 0:
                    construction_time = Factory._get_execution_time(
                        available_work, self.base_time)
                    self._assign_work(available_worker, construction_time)
                    node_available_time[available_work] = self.current_time + \
                        construction_time
                else:
                    break

            if not graph.get_remaining_nodes(done):
                return max(node_available_time.values())

            self.current_time += 1

    def _get_available_worker(self):
        for index, value in enumerate(self.worker_available_time):
            if value <= self.current_time:
                return index

        return None

    def _assign_work(self, worker_index, work_length):
        self.worker_available_time[worker_index] = self.current_time + work_length

    @staticmethod
    def _get_execution_time(node, base_time=0):
        return ord(node) - ord('A') + 1 + base_time


def _parse_instruction(line):
    match = _INSTRUCTION_REGEX.match(line)
    if not match:
        raise Exception('Failed to parse input')

    return (match.group(1), match.group(2))


def run_part1(file_content):
    """Implmentation for Part 1."""
    pairs = [_parse_instruction(line) for line in file_content]

    graph = Graph()
    for pair in pairs:
        graph.connect_nodes(pair[0], pair[1])

    return ''.join(graph.get_node_order())


def run_part2(file_content, worker_count, base_time):
    """Implmentation for Part 2."""
    pairs = [_parse_instruction(line) for line in file_content]

    graph = Graph()
    for pair in pairs:
        graph.connect_nodes(pair[0], pair[1])

    factory = Factory(worker_count, base_time)
    return factory.construct(graph)


if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content, 5, 60)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
