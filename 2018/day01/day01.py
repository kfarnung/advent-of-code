"""
Implementation for Advent of Code Day 1.

https://adventofcode.com/2018/day/1
"""

def get_new_frequency(frequency, modification):
    """Calculates the frequency for a given input."""
    operation = modification[0]
    magnitude = int(modification[1:])
    if operation == '+':
        return frequency + magnitude

    return frequency - magnitude


def calculate_frequency(inputs):
    """Calculates the frequency for a given set of inputs."""
    frequency = 0

    for modification in inputs:
        frequency = get_new_frequency(frequency, modification)

    return frequency

def calculate_frequency_two_match(inputs):
    """Calculates the frequency for a given set of inputs."""
    frequency = 0
    seen = set()
    seen.add(frequency)

    while True:
        for modification in inputs:
            frequency = get_new_frequency(frequency, modification)

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
