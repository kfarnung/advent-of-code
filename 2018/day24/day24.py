"""
Implementation for Advent of Code Day 24.

https://adventofcode.com/2018/day/24
"""

import re
from collections import defaultdict
from functools import reduce

_TEAM_REGEX = re.compile(r'^([A-Za-z ]+):$')
_GROUP_REGEX = re.compile(
    r'^(\d+) units each with (\d+) hit points (?:\(([a-z,; ]+)\) )?with an attack that does (\d+) '
    r'([a-z]+) damage at initiative (\d+)$'
)
_MODIFIER_REGEX = re.compile(r'([a-z]+) to ([a-z, ]+)')


class Group:  # pylint: disable=too-many-instance-attributes
    """Represents a single group of attackers."""

    def __init__(self, team_name, unit_count, hit_points, attack_damage, attack_type, initiative):  # pylint: disable=too-many-arguments,too-many-positional-arguments
        self.team_name = team_name
        self._unit_count = unit_count
        self._hit_points = hit_points
        self._attack_damage = attack_damage
        self.attack_type = attack_type
        self.initiative = initiative

        self._attack_boost = 0
        self._units_remaining = self._unit_count
        self._weaknesses = set()
        self._immunities = set()

    def __str__(self):
        return '{} contains {} units'.format(self.team_name, self.remaining_units())

    def boost(self, attack_boost):
        """Apply a boost to attacks"""
        self._attack_boost = attack_boost

    def effective_damage(self, target):
        """Calculates the effective damage of the group on a given target."""
        modifier = 1
        if target.immune_to_attack(self.attack_type):
            modifier = 0
        elif target.weak_to_attack(self.attack_type):
            modifier = 2

        return modifier * self.effective_power()

    def effective_power(self):
        """Calculates the effective power of the group."""
        return self._units_remaining * (self._attack_damage + self._attack_boost)

    def is_enemy(self, other):
        """Determine whether two players are enemies."""
        return self.team_name != other.team_name

    def get_target_selection_key(self):
        """Gets the key which determines the order of target selection."""
        return lambda other: (
            self.effective_damage(other),
            other.effective_power(),
            other.initiative
        )

    def get_targeting_key(self):
        """Gets the key which determines the order of attacks."""
        return (-self.effective_power(), -self.initiative)

    def has_units(self):
        """Indicates whether the group has any units remaining."""
        return self._units_remaining > 0

    def immune_to_attack(self, attack_type):
        """Determines whether the group is immune to the specified attack."""
        return attack_type in self._immunities

    def parse_modifiers(self, line):
        """Parse the modifiers from the provided text."""
        if not line:
            return

        for modifier in line.split('; '):
            match = _MODIFIER_REGEX.match(modifier)
            if not match:
                raise ValueError('Could not locate modifiers')

            modifier_set = None
            modifier_category = match.group(1)
            if modifier_category == 'weak':
                modifier_set = self._weaknesses
            elif modifier_category == 'immune':
                modifier_set = self._immunities
            else:
                raise ValueError('Unknown modifier category encountered')

            for modifier_type in match.group(2).split(', '):
                modifier_set.add(modifier_type)

    def remaining_units(self):
        """Count the number of units remaining."""
        return self._units_remaining if self.has_units() else 0

    def respawn(self):
        """Respawn the units."""
        self._units_remaining = self._unit_count

    def take_damage(self, attacker):
        """Apply damage to the group."""
        self._units_remaining -= (attacker.effective_damage(self) //
                                  self._hit_points)

    def weak_to_attack(self, attack_type):
        """Determines whether the group is weak to the specified attack."""
        return attack_type in self._weaknesses

    @staticmethod
    def parse(team_name, line):
        """Parse the group from the provided text."""
        group_match = _GROUP_REGEX.match(line)
        if not group_match:
            raise ValueError('Could not locate group')

        group = Group(
            team_name,
            int(group_match.group(1)),
            int(group_match.group(2)),
            int(group_match.group(4)),
            group_match.group(5),
            int(group_match.group(6)),
        )

        group.parse_modifiers(group_match.group(3))

        return group


class Battle:
    """Represents a battle between the groups."""

    def __init__(self, groups):
        self._groups = groups
        self._teams = defaultdict(list)
        for group in groups:
            self._teams[group.team_name].append(group)

    def boost_team(self, team_name, attack_boost):
        """Boost the attack for a named team."""
        for group in self._groups:
            if group.team_name == team_name:
                group.boost(attack_boost)

    def fight(self):
        """Simulate the fight between the groups."""
        for group in self._groups:
            group.respawn()

        while self.winner() is None:
            targets_list = []
            possible_groups = [
                group for group in self._groups if group.has_units()]
            possible_targets = set(possible_groups)

            possible_groups.sort(key=Group.get_targeting_key)
            for attacker in possible_groups:
                targets = [
                    group for group in possible_targets
                    if attacker.is_enemy(group) and attacker.effective_damage(group) > 0
                ]

                if not targets:
                    continue

                target = max(
                    targets,
                    key=attacker.get_target_selection_key()
                )

                possible_targets.remove(target)
                targets_list.append((attacker, target))

            if not targets_list:
                # We've reached a stalemate.
                break

            targets_list.sort(key=Battle._attacking_key)
            for attacker, target in targets_list:
                if attacker.has_units():
                    target.take_damage(attacker)

    def remaining_units(self):
        """Count the number of units remaining in the battle."""
        return reduce(
            lambda prev, item: prev + item.remaining_units(),
            self._groups,
            0
        )

    def winner(self):
        """Determine the winner of the battle."""
        teams_with_units = [
            team for team in self._teams.values() if Battle._team_has_units(team)
        ]

        if len(teams_with_units) == 1:
            return teams_with_units[0][0].team_name

        return None

    @staticmethod
    def _attacking_key(pair):
        return -pair[0].initiative

    @staticmethod
    def _team_has_units(team):
        for group in team:
            if group.has_units():
                return True

        return False


def _parse_team(group_list, lines):
    """Parse the team from the provided lines of text."""
    team_match = _TEAM_REGEX.match(next(lines))
    if not team_match:
        raise ValueError('Could not locate team')

    team_name = team_match.group(1)

    for group_line in lines:
        group_line = group_line.strip()
        if not group_line:
            break

        group_list.append(Group.parse(team_name, group_line))


def _load_groups(file_content):
    line_iterator = iter(file_content)
    groups = []
    team_count = 0

    try:
        while True:
            _parse_team(groups, line_iterator)
            team_count += 1
    except StopIteration:
        pass

    if team_count != 2:
        raise ValueError('Failed to find two teams')

    return groups


def run_part1(file_content):
    """Implmentation for Part 1."""
    groups = _load_groups(file_content)
    battle = Battle(groups)
    battle.fight()

    return battle.remaining_units()


def run_part2(file_content):
    """Implmentation for Part 2."""
    groups = _load_groups(file_content)
    battle = Battle(groups)

    boost = 0

    while True:
        battle.boost_team("Immune System", boost)
        battle.fight()

        if battle.winner() == "Immune System":
            return battle.remaining_units()

        boost += 1

    return None


if __name__ == "__main__":
    import sys

    def run(argv1):
        """The main function."""
        with open(argv1, 'r') as input_file:
            file_content = input_file.readlines()
            print("Part 1: {}".format(run_part1(file_content)))
            print("Part 2: {}".format(run_part2(file_content)))

    if len(sys.argv) < 2:
        print("Usage: python {} <input>".format(sys.argv[0]))
        sys.exit(1)

    run(sys.argv[1])
