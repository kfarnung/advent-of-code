"""
Implementation for Advent of Code Day 15.

https://adventofcode.com/2018/day/15
"""

from collections import defaultdict, deque
from operator import attrgetter

class Unit(object):
    """Represents a single unit in the battle."""
    def __init__(self, unit_type, position, attack_power):
        self.unit_type = unit_type
        self.position = position
        self.hit_points = 200
        self.attack_power = attack_power

    def is_alive(self):
        """Determines whether the unit is still alive."""
        return self.hit_points > 0

    @staticmethod
    def compare(first, second):
        """Compare two units when choosing who to attack."""
        value = cmp(first.hit_points, second.hit_points)
        if value != 0:
            return value

        return cmp(first.position, second.position)

class Battle(object):
    """Represents the current state of the battle."""
    def __init__(self, initial_state, elf_power=3):
        self.rounds_completed = 0
        self.units = []
        self.board = [[col for col in row if col != '\n'] for row in initial_state]

        for row_index, row in enumerate(self.board):
            for col_index, col in enumerate(row):
                if col == 'G':
                    self.units.append(Unit(col, (row_index, col_index), 3))
                    row[col_index] = '.'
                elif col == 'E':
                    self.units.append(Unit(col, (row_index, col_index), elf_power))
                    row[col_index] = '.'
                elif col != '#' and col != '.':
                    raise Exception('Unexpected tile type')

    def __str__(self):
        board = [[col for col in row] for row in self.board]
        for unit in self.units:
            if not unit.is_alive():
                continue

            board[unit.position[0]][unit.position[1]] = unit.unit_type

        return '\n'.join(''.join(row) for row in board)

    def get_team_sizes(self):
        """Gets the size of each team."""
        counts = defaultdict(int)
        for unit in self.units:
            counts[unit.unit_type] += 1

        return counts

    def fight(self):
        """Execute the entire battle."""
        while self._execute_round():
            pass

        return [unit for unit in self.units if unit.is_alive()]

    def _execute_round(self):
        """Execute one round of the battle."""
        for unit in sorted(self.units, key=attrgetter('position')):
            if not unit.is_alive():
                continue

            # If we can't execute a turn, the battle is over
            if not self._execute_turn(unit):
                return False

        self.rounds_completed += 1
        return True

    def _execute_turn(self, unit):
        """Execute one turn of the battle."""
        targets = self._get_targets(unit)
        if not targets:
            return False

        # Try to attack first, if we succeed that's all we'll do.
        if not self._try_attack(unit, targets):
            # Find the closest destination
            target_positions = [target.position for target in targets]
            closest_target, _ = self._find_closest_target(unit.position, target_positions)
            if closest_target:
                closest_step, _ = self._find_closest_target(
                    closest_target, self._get_open_neighbors(unit.position))
                unit.position = closest_step

                # Try to attack again after moving.
                self._try_attack(unit, targets)

        return True

    def _find_closest_target(self, start, targets):
        seen = set()
        to_visit = deque([(start, 0)])
        shortest_distance = None
        closest = []

        while to_visit:
            position, distance = to_visit.popleft()

            if shortest_distance and distance > shortest_distance:
                break

            if position in seen:
                continue

            seen.add(position)

            if position in targets:
                shortest_distance = distance
                closest.append(position)

            if position == start or self._is_open_cavern(position):
                for neighbor in self._get_neighbors(position):
                    if neighbor not in seen:
                        to_visit.append((neighbor, distance + 1))

        return min(closest) if closest else None, shortest_distance

    def _try_attack(self, unit, targets):
        in_range = [target for target in targets
                    if Battle._is_adjacent(unit.position, target.position)]
        if in_range:
            in_range.sort(Unit.compare)
            in_range[0].hit_points -= unit.attack_power
            return True

        return False

    def _is_open_cavern(self, position):
        board_cell = self.board[position[0]][position[1]]
        if board_cell != '.':
            return False

        for unit in self.units:
            if unit.is_alive() and unit.position == position:
                return False

        return True

    def _get_neighbors(self, position):
        return [
            (position[0] - 1, position[1]),
            (position[0], position[1] - 1),
            (position[0], position[1] + 1),
            (position[0] + 1, position[1]),
        ]

    def _get_open_neighbors(self, position):
        return [neighbor for neighbor in self._get_neighbors(position)
                if self._is_open_cavern(neighbor)]

    def _get_targets(self, unit):
        return [potential_target for potential_target in self.units
                if potential_target.unit_type != unit.unit_type and potential_target.is_alive()]

    @staticmethod
    def _manhattan_distance(first, second):
        """Calculates the Manhattan distance between the two points."""
        return abs(first[0] - second[0]) + abs(first[1] - second[1])

    @staticmethod
    def _is_adjacent(first, second):
        return Battle._manhattan_distance(first, second) == 1

def run_part1(file_content):
    """Implmentation for Part 1."""
    battle = Battle(file_content)
    hit_points = sum(unit.hit_points for unit in battle.fight())
    return hit_points * battle.rounds_completed

def run_part2(file_content):
    """Implmentation for Part 2."""
    elf_power = 4
    while elf_power < 1000:
        battle = Battle(file_content, elf_power)
        teams = battle.get_team_sizes()
        remaining = battle.fight()
        if len(remaining) == teams['E'] and remaining[0].unit_type == 'E':
            hit_points = sum(unit.hit_points for unit in battle.fight())
            return hit_points * battle.rounds_completed

        elf_power += 1

    return None

if __name__ == "__main__":
    import sys

    def run(argv1):
        """The main function."""
        with open(argv1, 'r') as input_file:
            file_content = input_file.readlines()
            print "Part 1: {}".format(run_part1(file_content))
            print "Part 2: {}".format(run_part2(file_content))

    if len(sys.argv) < 2:
        print "Usage: python {} <input>".format(sys.argv[0])
        sys.exit(1)

    run(sys.argv[1])
