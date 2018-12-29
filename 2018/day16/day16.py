"""
Implementation for Advent of Code Day 16.

https://adventofcode.com/2018/day/16
"""

import re
from collections import defaultdict

_BEFORE_REGEX = re.compile(r'^Before: \[(\d+), (\d+), (\d+), (\d+)\]$')
_INSTRUCTION_REGEX = re.compile(r'^(\d+) (\d+) (\d+) (\d+)$')
_AFTER_REGEX = re.compile(r'^After:  \[(\d+), (\d+), (\d+), (\d+)\]$')

class Device:
    """Represents the current state of the device."""
    def __init__(self):
        self.registers = [0] * 4
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

    def get_opcodes(self):
        """Gets the available opcodes."""
        return self.opcodes.keys()

    def execute(self, name, input_a, input_b, output_c):
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

def _parse_samples(file_content):
    samples = []
    instructions = []
    before_state = None
    instruction = None

    for line in file_content:
        match = _BEFORE_REGEX.match(line)
        if match:
            before_state = [int(group) for group in match.groups()]

        match = _INSTRUCTION_REGEX.match(line)
        if match:
            instruction = [int(group) for group in match.groups()]
            if not before_state:
                instructions.append(instruction)
                instruction = None

        match = _AFTER_REGEX.match(line)
        if match:
            samples.append((before_state, instruction, [int(group) for group in match.groups()]))
            before_state = None
            instruction = None

    return samples, instructions

def _get_instruction_map(samples):
    probabilities = defaultdict(set)
    device = Device()

    for sample in samples:
        for name, func in device.opcodes.items():
            device.registers = list(sample[0])
            func(device, sample[1][1], sample[1][2], sample[1][3])
            if device.registers == sample[2]:
                probabilities[sample[1][0]].add(name)

    instruction_map = {}

    while probabilities:
        single_match = [(key, value) for key, value in probabilities.items() if len(value) == 1]
        for match in single_match:
            del probabilities[match[0]]
            item = [item for item in match[1]][0]
            instruction_map[match[0]] = item

            for value in probabilities.values():
                if item in value:
                    value.remove(item)

    return instruction_map

def run_part1(file_content):
    """Implmentation for Part 1."""
    samples, _ = _parse_samples(file_content)
    device = Device()
    sample_count = 0

    for sample in samples:
        opcode_count = 0
        for name in device.get_opcodes():
            device.registers = list(sample[0])
            device.execute(name, sample[1][1], sample[1][2], sample[1][3])
            if device.registers == sample[2]:
                opcode_count += 1
        if opcode_count >= 3:
            sample_count += 1

    return sample_count

def run_part2(file_content):
    """Implmentation for Part 2."""
    samples, instructions = _parse_samples(file_content)
    instruction_map = _get_instruction_map(samples)
    device = Device()

    for instruction in instructions:
        name = instruction_map[instruction[0]]
        device.execute(name, instruction[1], instruction[2], instruction[3])

    return device.registers[0]

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
