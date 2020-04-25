"""
Implementation for Advent of Code Day 21.

https://adventofcode.com/2018/day/21
"""

from __future__ import print_function

import re

_IP_REGEX = re.compile(r'^#ip (\d)$')
_INSTRUCTION_REGEX = re.compile(r'^([a-z]+) (\d+) (\d+) (\d+)$')

class Device:
    """Represents the current state of the device."""
    def __init__(self, num_registers):
        self.registers = [0] * num_registers

    def execute_program(self, reg_ip, instructions, instruction_index, register_index):
        """Executes the entire program given an instruction-pointer register and instructions."""
        instruction_pointer = 0
        while 0 <= instruction_pointer < len(instructions):
            if instruction_index == instruction_pointer:
                yield self.registers[register_index]
            self.registers[reg_ip] = instruction_pointer
            instruction = instructions[instruction_pointer]
            self._execute(instruction[0], instruction[1], instruction[2], instruction[3])
            instruction_pointer = self.registers[reg_ip] + 1

    def _execute(self, name, in_a, in_b, out_c):
        """Executes the given instruction."""
        if name == 'addr':
            self.registers[out_c] = self.registers[in_a] + self.registers[in_b]
        elif name == 'addi':
            self.registers[out_c] = self.registers[in_a] + in_b
        elif name == 'mulr':
            self.registers[out_c] = self.registers[in_a] * self.registers[in_b]
        elif name == 'muli':
            self.registers[out_c] = self.registers[in_a] * in_b
        elif name == 'banr':
            self.registers[out_c] = self.registers[in_a] & self.registers[in_b]
        elif name == 'bani':
            self.registers[out_c] = self.registers[in_a] & in_b
        elif name == 'borr':
            self.registers[out_c] = self.registers[in_a] | self.registers[in_b]
        elif name == 'bori':
            self.registers[out_c] = self.registers[in_a] | in_b
        elif name == 'setr':
            self.registers[out_c] = self.registers[in_a]
        elif name == 'seti':
            self.registers[out_c] = in_a
        elif name == 'gtir':
            self.registers[out_c] = 1 if in_a > self.registers[in_b] else 0
        elif name == 'gtri':
            self.registers[out_c] = 1 if self.registers[in_a] > in_b else 0
        elif name == 'gtrr':
            self.registers[out_c] = 1 if self.registers[in_a] > self.registers[in_b] else 0
        elif name == 'eqir':
            self.registers[out_c] = 1 if in_a == self.registers[in_b] else 0
        elif name == 'eqri':
            self.registers[out_c] = 1 if self.registers[in_a] == in_b else 0
        elif name == 'eqrr':
            self.registers[out_c] = 1 if self.registers[in_a] == self.registers[in_b] else 0
        else:
            raise Exception('Invalid opcode')

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

    instr_index = 28
    reg_index = instructions[instr_index][1]
    for value in device.execute_program(reg_ip, instructions, instr_index, reg_index):
        return value

def run_part2(file_content):
    """Implmentation for Part 2."""
    reg_ip, instructions = _parse_program(file_content)
    device = Device(6)

    instr_index = 28
    reg_index = instructions[instr_index][1]
    seen = set()
    last = None
    for value in device.execute_program(reg_ip, instructions, instr_index, reg_index):
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
