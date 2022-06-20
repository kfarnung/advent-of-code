"""
Implementation for Advent of Code Day 19.

https://adventofcode.com/2018/day/19
"""

from __future__ import print_function

import re

_IP_REGEX = re.compile(r'^#ip (\d)$')
_INSTRUCTION_REGEX = re.compile(r'^([a-z]+) (\d+) (\d+) (\d+)$')


class Device:  # pylint: disable=too-few-public-methods
    """Represents the current state of the device."""

    def __init__(self, num_registers):
        self.registers = [0] * num_registers
        self.opcodes = {
            'addr': Device._instruction_addr,
            'addi': Device._instruction_addi,
            'mulr': Device._instruction_mulr,
            'muli': Device._instruction_muli,
            'banr': Device._instruction_banr,
            'bani': Device._instruction_bani,
            'borr': Device._instruction_borr,
            'bori': Device._instruction_bori,
            'setr': Device._instruction_setr,
            'seti': Device._instruction_seti,
            'gtir': Device._instruction_gtir,
            'gtri': Device._instruction_gtri,
            'gtrr': Device._instruction_gtrr,
            'eqir': Device._instruction_eqir,
            'eqri': Device._instruction_eqri,
            'eqrr': Device._instruction_eqrr,
        }

    def execute_program(self, reg_ip, instructions):
        """Executes the entire program given an instruction-pointer register and instructions."""
        instruction_pointer = 0
        while 0 <= instruction_pointer < len(instructions):
            self.registers[reg_ip] = instruction_pointer
            instruction = instructions[instruction_pointer]
            self._execute(instruction[0], instruction[1],
                          instruction[2], instruction[3])
            instruction_pointer = self.registers[reg_ip] + 1

    def _execute(self, name, input_a, input_b, output_c):
        """Executes the given instruction."""
        self.opcodes[name](self, input_a, input_b, output_c)

    def _instruction_addr(self, reg_a, reg_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] + self.registers[reg_b]

    def _instruction_addi(self, reg_a, value_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] + value_b

    def _instruction_mulr(self, reg_a, reg_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] * self.registers[reg_b]

    def _instruction_muli(self, reg_a, value_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] * value_b

    def _instruction_banr(self, reg_a, reg_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] & self.registers[reg_b]

    def _instruction_bani(self, reg_a, value_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] & value_b

    def _instruction_borr(self, reg_a, reg_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] | self.registers[reg_b]

    def _instruction_bori(self, reg_a, value_b, reg_c):
        self.registers[reg_c] = self.registers[reg_a] | value_b

    def _instruction_setr(self, reg_a, _, reg_c):
        self.registers[reg_c] = self.registers[reg_a]

    def _instruction_seti(self, value_a, _, reg_c):
        self.registers[reg_c] = value_a

    def _instruction_gtir(self, value_a, reg_b, reg_c):
        self.registers[reg_c] = 1 if value_a > self.registers[reg_b] else 0

    def _instruction_gtri(self, reg_a, value_b, reg_c):
        self.registers[reg_c] = 1 if self.registers[reg_a] > value_b else 0

    def _instruction_gtrr(self, reg_a, reg_b, reg_c):
        self.registers[reg_c] = 1 if self.registers[reg_a] > self.registers[reg_b] else 0

    def _instruction_eqir(self, value_a, reg_b, reg_c):
        self.registers[reg_c] = 1 if value_a == self.registers[reg_b] else 0

    def _instruction_eqri(self, reg_a, value_b, reg_c):
        self.registers[reg_c] = 1 if self.registers[reg_a] == value_b else 0

    def _instruction_eqrr(self, reg_a, reg_b, reg_c):
        self.registers[reg_c] = 1 if self.registers[reg_a] == self.registers[reg_b] else 0


def _parse_program(file_content):
    reg_ip = None
    instructions = []

    for line in file_content:
        match = _IP_REGEX.match(line)
        if match:
            reg_ip = int(match.group(1))

        match = _INSTRUCTION_REGEX.match(line)
        if match:
            instruction = [match.group(1)]
            instruction += [int(group) for group in match.groups()[1:]]
            instructions.append(instruction)

    return reg_ip, instructions


def run_part1(file_content):
    """Implmentation for Part 1."""
    reg_ip, instructions = _parse_program(file_content)
    device = Device(6)
    device.execute_program(reg_ip, instructions)

    return device.registers[0]


def run_part2(_):
    """Implmentation for Part 2."""
    # Reverse-engineered the input, it simply does:
    return sum(val for val in range(1, 10551340 + 1) if 10551340 % val == 0)


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
