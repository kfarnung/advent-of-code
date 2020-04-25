"""
Implementation for Advent of Code Day 11.

https://adventofcode.com/2018/day/11
"""

from __future__ import print_function


def get_power_level(serial_number, coord_x, coord_y):
    """Gets the power level for a given coordinate."""
    rack_id = coord_x + 10
    power_level = rack_id * coord_y
    power_level += serial_number
    power_level *= rack_id
    power_level = (power_level // 100) % 10
    power_level -= 5
    return power_level


def _init_grid(size, serial_number):
    """Initialize the grid with the sums from 0,0."""
    grid = []

    for coord_x in range(size):
        col = []
        for coord_y in range(size):
            current_power = get_power_level(
                serial_number, coord_x + 1, coord_y + 1)

            if coord_x > 0:
                current_power += grid[coord_x - 1][coord_y]

            if coord_y > 0:
                current_power += col[coord_y - 1]

            if coord_x > 0 and coord_y > 0:
                current_power -= grid[coord_x - 1][coord_y - 1]

            col.append(current_power)

        grid.append(col)

    return grid


def _find_max_power(grid, size):
    """Find the maximum power in the grid for a given size."""
    size -= 1
    max_power = grid[size][size]
    max_coords = (1, 1)

    for coord_x in range(300 - size):
        for coord_y in range(300 - size):
            power = grid[coord_x + size][coord_y + size]

            if coord_x > 0:
                power -= grid[coord_x - 1][coord_y + size]

            if coord_y > 0:
                power -= grid[coord_x + size][coord_y - 1]

            if coord_x > 0 and coord_y > 0:
                power += grid[coord_x - 1][coord_y - 1]

            if power > max_power:
                max_power = power
                max_coords = (coord_x + 1, coord_y + 1)

    return max_power, max_coords


def run_part1(serial_number):
    """Implmentation for Part 1."""
    grid = _init_grid(300, serial_number)
    return _find_max_power(grid, 3)[1]


def run_part2(serial_number):
    """Implmentation for Part 2."""
    grid_size = 300
    grid = _init_grid(grid_size, serial_number)
    max_power, max_coords = _find_max_power(grid, 1)
    max_size = 1

    for size in range(2, grid_size + 1):
        power, coords = _find_max_power(grid, size)
        if power > max_power:
            max_power = power
            max_coords = coords
            max_size = size

    return max_coords, max_size


if __name__ == "__main__":
    import sys

    def run(input_str):
        """The main function."""
        serial_number = int(input_str)

        coord_x, coord_y = run_part1(serial_number)
        print("Part 1: {},{}".format(coord_x, coord_y))

        coords, size = run_part2(serial_number)
        print("Part 2: {},{},{}".format(coords[0], coords[1], size))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
