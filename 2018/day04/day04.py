"""
Implementation for Advent of Code Day 4.

https://adventofcode.com/2018/day/4
"""

import re
from collections import defaultdict

_INSTRUCTION_REGEX = re.compile(
    r"^\[(\d{4}-\d{2}-\d{2} \d{2}:(\d{2}))] (Guard #(\d+) begins shift|.+)$")

class _GuardAction(object):
    """Represents a single action by a single guard"""
    def __init__(self, instruction):
        match = _INSTRUCTION_REGEX.match(instruction)
        assert match is not None

        self.time = match.group(1)
        self.minute = int(match.group(2))
        self.event = match.group(3)
        self.guard_id = int(match.group(4)) if match.group(4) else None

    def __lt__(self, other):
        return self.time < other.time

class _Guard(object):
    """Represents the aggregate actions of a single guard"""
    def __init__(self, guard_id):
        self.guard_id = guard_id
        self.sleep_minutes = defaultdict(int)
        self.total_sleep = 0

    def add_sleep_event(self, sleep, wake):
        """Logs a sleep event for the guard"""
        self.total_sleep += wake - sleep

        for minute in range(sleep, wake):
            self.sleep_minutes[minute] += 1

    def get_sleepiest_minute(self):
        """Finds the sleepiest minute for the guard"""
        if self.sleep_minutes:
            return max(self.sleep_minutes, key=self.sleep_minutes.get)

        return None

    def get_sleepiest_count(self):
        """Gets the frequency of the sleepiest minute for the guard"""
        sleepiest_minute = self.get_sleepiest_minute()
        if sleepiest_minute != None:
            return self.sleep_minutes[sleepiest_minute]

        return 0

def _parse_input(inputs):
    guard_map = {}
    guard = None
    sleep_minute = None

    for action in sorted([_GuardAction(action) for action in inputs]):
        if action.guard_id != None:
            if action.guard_id in guard_map:
                guard = guard_map[action.guard_id]
            else:
                guard = _Guard(action.guard_id)
                guard_map[action.guard_id] = guard
        elif 'sleep' in action.event:
            assert sleep_minute is None
            sleep_minute = action.minute
        elif 'wake' in action.event:
            guard.add_sleep_event(sleep_minute, action.minute)
            sleep_minute = None
        else:
            assert False

    return guard_map

def run_part1(inputs):
    """Implmentation for Part 1."""
    sleepiest_guard = max(
        _parse_input(inputs).itervalues(),
        key=lambda item: item.total_sleep,
    )

    return sleepiest_guard.guard_id * sleepiest_guard.get_sleepiest_minute()

def run_part2(inputs):
    """Implmentation for Part 2."""
    sleepiest_guard = max(
        _parse_input(inputs).itervalues(),
        key=lambda item: item.get_sleepiest_count(),
    )

    return sleepiest_guard.guard_id * sleepiest_guard.get_sleepiest_minute()

if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print "Part 1: {}".format(run_part1(file_content))
            print "Part 2: {}".format(run_part2(file_content))

    if len(sys.argv) < 2:
        print "Usage: python {} <input>".format(sys.argv[0])
        sys.exit(1)

    run(sys.argv[1])
