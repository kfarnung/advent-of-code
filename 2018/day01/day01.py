"""
Implementation for Advent of Code Day 1.

https://adventofcode.com/2018/day/1
"""

from functools import reduce


def calculate_frequency(inputs):
    """Calculates the frequency for a given set of inputs."""
    return reduce(lambda prev, current: prev + int(current), inputs, 0)


def calculate_frequency_two_match(inputs):
    """Calculates the frequency for a given set of inputs."""
    frequency = 0
    seen = set()
    seen.add(frequency)

    while True:
        for modification in inputs:
            frequency += int(modification)

            if frequency in seen:
                return frequency

            seen.add(frequency)


if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path) as input_file:
            file_content = input_file.readlines()
            print(f"Part 1: {calculate_frequency(file_content)}")
            print(f"Part 2: {calculate_frequency_two_match(file_content)}")

    if len(sys.argv) < 2:
        print(f"Usage: python {sys.argv[0]} <input>")
        sys.exit(1)

    run(sys.argv[1])
