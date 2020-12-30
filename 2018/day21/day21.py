"""
Implementation for Advent of Code Day 21.

https://adventofcode.com/2018/day/21
"""

from __future__ import print_function

import re

_IP_REGEX = re.compile(r'^#ip (\d)$')
_INSTRUCTION_REGEX = re.compile(r'^([a-z]+) (\d+) (\d+) (\d+)$')


def _addr(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] + registers[in_b]


def _addi(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] + in_b


def _mulr(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] * registers[in_b]


def _muli(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] * in_b


def _banr(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] & registers[in_b]


def _bani(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] & in_b


def _borr(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] | registers[in_b]


def _bori(registers, in_a, in_b, out_c):
    registers[out_c] = registers[in_a] | in_b


def _setr(registers, in_a, _in_b, out_c):
    registers[out_c] = registers[in_a]


def _seti(registers, in_a, _in_b, out_c):
    registers[out_c] = in_a


def _gtir(registers, in_a, in_b, out_c):
    registers[out_c] = 1 if in_a > registers[in_b] else 0


def _gtri(registers, in_a, in_b, out_c):
    registers[out_c] = 1 if registers[in_a] > in_b else 0


def _gtrr(registers, in_a, in_b, out_c):
    registers[out_c] = 1 if registers[in_a] > registers[in_b] else 0


def _eqir(registers, in_a, in_b, out_c):
    registers[out_c] = 1 if in_a == registers[in_b] else 0


def _eqri(registers, in_a, in_b, out_c):
    registers[out_c] = 1 if registers[in_a] == in_b else 0


def _eqrr(registers, in_a, in_b, out_c):
    registers[out_c] = 1 if registers[in_a] == registers[in_b] else 0


_INSTRUCTION_HANDLERS = {
    'addr': _addr,
    'addi': _addi,
    'mulr': _mulr,
    'muli': _muli,
    'banr': _banr,
    'bani': _bani,
    'borr': _borr,
    'bori': _bori,
    'setr': _setr,
    'seti': _seti,
    'gtir': _gtir,
    'gtri': _gtri,
    'gtrr': _gtrr,
    'eqir': _eqir,
    'eqri': _eqri,
    'eqrr': _eqrr,
}


class Device:
    """Represents the current state of the device."""

    def __init__(self, num_registers):
        self.registers = [0] * num_registers

    def execute_program(self, reg_ip, instructions, instruction_index, register_index):
        """Executes the entire program given an instruction-pointer register and instructions."""
        registers = self.registers
        instruction_count = len(instructions)
        instruction_pointer = registers[reg_ip]

        while 0 <= instruction_pointer < instruction_count:
            instruction = instructions[instruction_pointer]
            instruction[0](registers, instruction[1],
                           instruction[2], instruction[3])

            instruction_pointer = registers[reg_ip] + 1
            registers[reg_ip] = instruction_pointer

            if instruction_index == instruction_pointer:
                break

        return registers[register_index]


def _parse_program(file_content):
    reg_ip = None
    instructions = []

    for line in file_content:
        match = _IP_REGEX.match(line)
        if match:
            reg_ip = int(match.group(1))

        match = _INSTRUCTION_REGEX.match(line)
        if match:
            instruction = [_INSTRUCTION_HANDLERS[match.group(1)]]
            instruction += [int(group) for group in match.groups()[1:]]
            instructions.append(instruction)

    return reg_ip, instructions


def run_part1(file_content):
    """Implmentation for Part 1."""
    reg_ip, instructions = _parse_program(file_content)
    device = Device(6)

    instr_index = 28
    reg_index = instructions[instr_index][1]
    return device.execute_program(reg_ip, instructions, instr_index, reg_index)


def run_part2(file_content):
    """Implmentation for Part 2."""
    reg_ip, instructions = _parse_program(file_content)
    device = Device(6)

    instr_index = 28
    reg_index = instructions[instr_index][1]
    seen = set()
    last = None

    while True:
        value = device.execute_program(
            reg_ip, instructions, instr_index, reg_index)
        if value in seen:
            return last

        seen.add(value)
        last = value

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
