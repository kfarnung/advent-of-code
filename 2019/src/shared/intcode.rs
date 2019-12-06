#[derive(Debug, PartialEq)]
struct IntcodeInstruction {
    opcode: i32,
    modes: Vec<i32>,
}

impl IntcodeInstruction {
    fn new(opcode: i32) -> Self {
        return IntcodeInstruction {
            opcode: opcode,
            modes: Vec::new(),
        };
    }

    pub fn parse(instruction: i32) -> Self {
        let mut current = instruction;
        let mut parsed = IntcodeInstruction::new(current % 100);

        // Remove the opcode.
        current /= 100;

        while current > 0 {
            parsed.modes.push(current % 10);
            current /= 10;
        }

        return parsed;
    }

    fn get_mode(&self, index: usize) -> i32 {
        let mode = self.modes.get(index);
        match mode {
            Some(x) => x.clone(),
            None => 0,
        }
    }
}

pub struct IntcodeComputer {
    initial: Vec<i32>,
    memory: Vec<i32>,
    inputs: Vec<i32>,
    outputs: Vec<i32>,
}

impl IntcodeComputer {
    pub fn new(initial: &Vec<i32>) -> Self {
        return IntcodeComputer {
            initial: initial.clone(),
            memory: initial.clone(),
            inputs: Vec::new(),
            outputs: Vec::new(),
        };
    }

    pub fn new_from_string(initial_memory: &str) -> Self {
        let initial_memory = initial_memory
            .trim()
            .split(',')
            .map(|i| i.parse::<i32>().unwrap())
            .collect();

        return Self::new(&initial_memory);
    }

    pub fn run(&mut self) {
        let mut ip = 0;

        loop {
            let instruction = IntcodeInstruction::parse(self.memory[ip]);
            match instruction.opcode {
                1 => ip = self.do_add(ip, &instruction),
                2 => ip = self.do_multiply(ip, &instruction),
                3 => ip = self.do_input(ip, &instruction),
                4 => ip = self.do_output(ip, &instruction),
                5 => ip = self.do_jump_if_true(ip, &instruction),
                6 => ip = self.do_jump_if_false(ip, &instruction),
                7 => ip = self.do_less_than(ip, &instruction),
                8 => ip = self.do_equals(ip, &instruction),
                99 => break,
                _ => panic!("Unexpected opcode!"),
            };
        }
    }

    pub fn reset(&mut self) {
        self.memory = self.initial.clone();
    }

    pub fn set_inputs(&mut self, values: &Vec<i32>) {
        self.inputs = values.clone();
    }

    pub fn get_outputs(&mut self) -> &Vec<i32> {
        return &self.outputs;
    }

    pub fn get_memory(&self, address: usize) -> i32 {
        return self.memory[address];
    }

    pub fn set_value(&mut self, address: usize, value: i32) {
        self.memory[address] = value;
    }

    fn load_parameter(&self, ip: usize, instruction: &IntcodeInstruction, index: usize) -> i32 {
        let param = self.memory[ip + 1 + index];
        match instruction.get_mode(index) {
            0 => self.memory[param as usize],
            1 => param,
            _ => panic!("Unexpected mode!"),
        }
    }

    fn store_parameter(
        &mut self,
        ip: usize,
        instruction: &IntcodeInstruction,
        index: usize,
        value: i32,
    ) {
        let param = self.memory[ip + 1 + index];
        match instruction.get_mode(index) {
            0 => self.memory[param as usize] = value,
            _ => panic!("Unexpected mode!"),
        }
    }

    fn do_add(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let a = self.load_parameter(ip, instruction, 0);
        let b = self.load_parameter(ip, instruction, 1);

        self.store_parameter(ip, instruction, 2, a + b);
        return ip + 4;
    }

    fn do_multiply(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let a = self.load_parameter(ip, instruction, 0);
        let b = self.load_parameter(ip, instruction, 1);

        self.store_parameter(ip, instruction, 2, a * b);
        return ip + 4;
    }

    fn do_input(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let value = self.inputs.remove(0);
        self.store_parameter(ip, instruction, 0, value);
        return ip + 2;
    }

    fn do_output(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let value = self.load_parameter(ip, instruction, 0);
        self.outputs.push(value);
        return ip + 2;
    }

    fn do_jump_if_true(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let a = self.load_parameter(ip, instruction, 0);
        if a != 0 {
            return self.load_parameter(ip, instruction, 1) as usize;
        }

        return ip + 3;
    }

    fn do_jump_if_false(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let a = self.load_parameter(ip, instruction, 0);
        if a == 0 {
            return self.load_parameter(ip, instruction, 1) as usize;
        }

        return ip + 3;
    }

    fn do_less_than(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let a = self.load_parameter(ip, instruction, 0);
        let b = self.load_parameter(ip, instruction, 1);

        let value = if a < b { 1 } else { 0 };
        self.store_parameter(ip, instruction, 2, value);
        return ip + 4;
    }

    fn do_equals(&mut self, ip: usize, instruction: &IntcodeInstruction) -> usize {
        let a = self.load_parameter(ip, instruction, 0);
        let b = self.load_parameter(ip, instruction, 1);

        let value = if a == b { 1 } else { 0 };
        self.store_parameter(ip, instruction, 2, value);
        return ip + 4;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_run() {
        let cases = vec![
            (
                vec![1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50],
                vec![3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50],
            ),
            (vec![1, 0, 0, 0, 99], vec![2, 0, 0, 0, 99]),
            (vec![2, 3, 0, 3, 99], vec![2, 3, 0, 6, 99]),
            (vec![2, 4, 4, 5, 99, 0], vec![2, 4, 4, 5, 99, 9801]),
            (
                vec![1, 1, 1, 4, 99, 5, 6, 0, 99],
                vec![30, 1, 1, 4, 2, 5, 6, 0, 99],
            ),
        ];

        for case in cases {
            let mut computer = IntcodeComputer::new(&case.0);
            computer.run();
            assert_eq!(computer.memory, case.1);
        }
    }

    #[test]
    fn test_parse_instruction() {
        let parsed = IntcodeInstruction::parse(1002);
        assert_eq!(parsed.opcode, 2);
        assert_eq!(parsed.get_mode(0), 0);
        assert_eq!(parsed.get_mode(1), 1);
        assert_eq!(parsed.get_mode(2), 0);
    }
}
