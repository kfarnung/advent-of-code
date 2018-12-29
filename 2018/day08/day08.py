"""
Implementation for Advent of Code Day 8.

https://adventofcode.com/2018/day/8
"""

class TreeNode:
    """Represents a single node in the tree."""
    def __init__(self):
        self.children = []
        self.metadata = []

def _build_tree(inputs):
    child_count = next(inputs)
    metadata_count = next(inputs)
    node = TreeNode()

    for _ in range(child_count):
        node.children.append(_build_tree(inputs))

    for _ in range(metadata_count):
        node.metadata.append(next(inputs))

    return node

def _sum_metadata(node):
    metadata_sum = 0

    for child in node.children:
        metadata_sum += _sum_metadata(child)

    metadata_sum += sum(node.metadata)
    return metadata_sum

def _node_value(node):
    value = 0

    if node.children:
        for index in node.metadata:
            index -= 1
            if index < len(node.children):
                value += _node_value(node.children[index])
    else:
        value += sum(node.metadata)

    return value

def run_part1(file_content):
    """Implmentation for Part 1."""
    numbers = (int(number) for number in file_content.split())
    root = _build_tree(numbers)
    return _sum_metadata(root)

def run_part2(file_content):
    """Implmentation for Part 2."""
    numbers = (int(number) for number in file_content.split())
    root = _build_tree(numbers)
    return _node_value(root)

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
