"""
Implementation for Advent of Code Day 14.

https://adventofcode.com/2018/day/14
"""

from __future__ import print_function


class RecipeGenerator:
    """Generator for the recipe combinations."""

    def __init__(self):
        self.recipes = [3, 7]
        self.index_1 = 0
        self.index_2 = 1
        self.next_index = 0

    def __iter__(self):
        return self

    def __next__(self):
        """Get the next recipe generated."""
        if self.next_index >= len(self.recipes):
            elf_1 = self.recipes[self.index_1]
            elf_2 = self.recipes[self.index_2]

            self.recipes += [int(recipe) for recipe in str(elf_1 + elf_2)]
            recipes_len = len(self.recipes)

            self.index_1 = (self.index_1 + 1 + elf_1) % recipes_len
            self.index_2 = (self.index_2 + 1 + elf_2) % recipes_len

        recipe = self.recipes[self.next_index]
        self.next_index += 1
        return recipe

    def next(self):
        return self.__next__()

    def skip(self, count):
        """Skip the specified number of recipes."""
        for _ in range(count):
            next(self)

    def take(self, count):
        """Take the specified number of recipes."""
        result = []
        for _ in range(count):
            result.append(next(self))

        return result


class DigitsMatcher:
    """Look in the incoming digits for a match."""

    def __init__(self, digits):
        self.digits = [int(digit) for digit in str(digits)]
        self.digits_next = 0
        self.current_index = 0
        self.start_index = -1

    def match(self, digit):
        """Check the next digit."""
        if self.digits[self.digits_next] == digit:
            if self.digits_next == 0:
                self.start_index = self.current_index
            self.digits_next += 1

            if self.digits_next == len(self.digits):
                return self.start_index
        else:
            self.digits_next = 0
            self.start_index = -1

        self.current_index += 1
        return None


def run_part1(recipe_count):
    """Implmentation for Part 1."""
    generator = RecipeGenerator()
    generator.skip(recipe_count)
    return ''.join(str(recipe) for recipe in generator.take(10))


def run_part2(recipe_count):
    """Implmentation for Part 2."""
    generator = RecipeGenerator()
    matcher = DigitsMatcher(recipe_count)

    result = matcher.match(next(generator))
    while not result:
        result = matcher.match(next(generator))

    return result


if __name__ == "__main__":
    import sys

    def run(argv1):
        """The main function."""
        recipe_count = int(argv1)
        print("Part 1: {}".format(run_part1(recipe_count)))
        print("Part 2: {}".format(run_part2(recipe_count)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
