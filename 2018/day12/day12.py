"""
Implementation for Advent of Code Day 12.

https://adventofcode.com/2018/day/12
"""

import re
from collections import defaultdict, deque
from functools import reduce
from itertools import count

_INITIAL_STATE_REGEX = re.compile(r'^initial state: ([#.]+)$')
_RULE_REGEX = re.compile(r'^([#.]+) => ([#.])$')


def _parse_input(file_content):
    rules = []
    list_iter = iter(file_content)
    match = _INITIAL_STATE_REGEX.match(next(list_iter))
    if not match:
        raise ValueError('Invalid initial state')

    initial_state = match.group(1)

    # Skip blank line
    next(list_iter)

    for rule in list_iter:
        match = _RULE_REGEX.match(rule)
        if not match:
            raise ValueError('Invalid rule pattern')

        rules.append((match.group(1), match.group(2)))

    return initial_state, rules


def _initialize_greenhouse(initial_state):
    greenhouse = defaultdict(lambda: '.')
    for index, state in enumerate(initial_state):
        greenhouse[index] = state

    return greenhouse


def _matches_rule(greenhouse, rule, pot_index):
    offset = len(rule) // 2
    for index in range(-offset, offset + 1):
        if greenhouse[pot_index + index] != rule[index + offset]:
            return False

    return True


def _get_replacement(greenhouse, rules, pot_index):
    for rule in rules:
        if _matches_rule(greenhouse, rule[0], pot_index):
            return rule[1]

    return '.'


def _next_generation(greenhouse, rules):
    min_index = min(greenhouse) - 2
    max_index = max(greenhouse) + 2
    next_gen = defaultdict(lambda: '.')

    for index in range(min_index, max_index + 1):
        next_gen[index] = _get_replacement(greenhouse, rules, index)

    return next_gen


def _sum_plant_locations(greenhouse):
    return reduce(
        lambda prev, item: prev + item[0] if item[1] == '#' else prev,
        greenhouse.items(),
        0
    )


def _get_key(greenhouse):
    min_index = min(greenhouse)
    max_index = max(greenhouse)

    # Manually iterate since normal iteration doesn't seem quite right.
    pattern = deque(greenhouse[index]
                    for index in range(min_index, max_index + 1))

    # Strip off empty pots
    while pattern[0] == '.':
        pattern.popleft()

    while pattern[-1] == '.':
        pattern.pop()

    return ''.join(pattern)


def _predict_plant_generations(greenhouse, rules, generation_count):
    seen_patterns = {}
    seen_patterns[_get_key(greenhouse)] = _sum_plant_locations(greenhouse)

    # Work around the maximum limits of `range` and the `next` vs. `__next__`
    # differences by capturing the count iter and wrapping in a lambda.
    count_iter = count()

    for index in iter(lambda: next(count_iter), generation_count):
        greenhouse = _next_generation(greenhouse, rules)
        key = _get_key(greenhouse)

        if key not in seen_patterns:
            seen_patterns[_get_key(greenhouse)] = _sum_plant_locations(
                greenhouse)
        else:
            # We've seen this one before, assume we've hit a steady state and extrapolate
            prev_score = seen_patterns[_get_key(greenhouse)]
            current_score = _sum_plant_locations(greenhouse)
            remaining_generations = generation_count - 1 - index
            return ((current_score - prev_score) * remaining_generations) + current_score

    return _sum_plant_locations(greenhouse)


def run_part1(file_content):
    """Implmentation for Part 1."""
    initial_state, rules = _parse_input(file_content)
    greenhouse = _initialize_greenhouse(initial_state)

    return _predict_plant_generations(greenhouse, rules, 20)


def run_part2(file_content):
    """Implmentation for Part 2."""
    initial_state, rules = _parse_input(file_content)
    greenhouse = _initialize_greenhouse(initial_state)

    return _predict_plant_generations(greenhouse, rules, 50000000000)


if __name__ == "__main__":
    import sys

    def run(input_path):
        """The main function."""
        with open(input_path, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
