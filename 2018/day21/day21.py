"""
Implementation for Advent of Code Day 21.

https://adventofcode.com/2018/day/21
"""

import re

_IP_REGEX = re.compile(r'^#ip (\d)$')
_INSTRUCTION_REGEX = re.compile(r'^([a-z]+) (\d+) (\d+) (\d+)$')

class Device(object):
    """Represents the current state of the device."""
    def __init__(self, num_registers):
        self.registers = [0] * num_registers

    def execute_program(self, reg_ip, instructions, instruction_index, register_index):
        """Executes the entire program given an instruction-pointer register and instructions."""
        instruction_pointer = 0
        while instruction_pointer >= 0 and instruction_pointer < len(instructions):
            if instruction_index == instruction_pointer:
                yield self.registers[register_index]
            self.registers[reg_ip] = instruction_pointer
            instruction = instructions[instruction_pointer]
            self._execute(instruction[0], instruction[1], instruction[2], instruction[3])
            instruction_pointer = self.registers[reg_ip] + 1

    def _execute(self, name, input_a, input_b, output_c):
        """Executes the given instruction."""
        if name == 'addr':
            self.registers[output_c] = self.registers[input_a] + self.registers[input_b]
        elif name == 'addi':
            self.registers[output_c] = self.registers[input_a] + input_b
        elif name == 'mulr':
            self.registers[output_c] = self.registers[input_a] * self.registers[input_b]
        elif name == 'muli':
            self.registers[output_c] = self.registers[input_a] * input_b
        elif name == 'banr':
            self.registers[output_c] = self.registers[input_a] & self.registers[input_b]
        elif name == 'bani':
            self.registers[output_c] = self.registers[input_a] & input_b
        elif name == 'borr':
            self.registers[output_c] = self.registers[input_a] | self.registers[input_b]
        elif name == 'bori':
            self.registers[output_c] = self.registers[input_a] | input_b
        elif name == 'setr':
            self.registers[output_c] = self.registers[input_a]
        elif name == 'seti':
            self.registers[output_c] = input_a
        elif name == 'gtir':
            self.registers[output_c] = 1 if input_a > self.registers[input_b] else 0
        elif name == 'gtri':
            self.registers[output_c] = 1 if self.registers[input_a] > input_b else 0
        elif name == 'gtrr':
            self.registers[output_c] = 1 if self.registers[input_a] > self.registers[input_b] else 0
        elif name == 'eqir':
            self.registers[output_c] = 1 if input_a == self.registers[input_b] else 0
        elif name == 'eqri':
            self.registers[output_c] = 1 if self.registers[input_a] == input_b else 0
        elif name == 'eqrr':
            self.registers[output_c] = 1 if self.registers[input_a] == self.registers[input_b] else 0

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
    reg_index = instructions[28][1]
    for value in device.execute_program(reg_ip, instructions, instr_index, reg_index):
        return value

def run_part2(file_content):
    """Implmentation for Part 2."""
    reg_ip, instructions = _parse_program(file_content)
    device = Device(6)

    instr_index = 28
    reg_index = instructions[28][1]
    seen = set()
    last = None
    for value in device.execute_program(reg_ip, instructions, instr_index, reg_index):
        if value in seen:
            return last

        seen.add(value)
        last = value

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
