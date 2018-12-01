"""
Implementation for Advent of Code Day 1.

https://adventofcode.com/2018/day/1
"""

def calculate_frequency(inputs):
    """Calculates the frequency for a given set of inputs."""
    return reduce(
        lambda prev, current: prev + int(current),
        inputs,
        0)

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
            else:
                seen.add(frequency)

if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print "Part 1: {}".format(calculate_frequency(file_content))
            print "Part 2: {}".format(calculate_frequency_two_match(file_content))

    if len(sys.argv) < 2:
        raise RuntimeError("Usage: day01.py <input>")

    run(sys.argv[1])
