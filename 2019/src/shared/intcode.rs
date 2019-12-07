use std::collections::VecDeque;

#[derive(Clone, Copy, Debug, PartialEq)]
pub enum IntcodeProcessState {
    Initialized,
    Running,
    Waiting,
    Halted,
}

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

pub struct IntcodeProcess {
    ip: usize,
    memory: Vec<i32>,
    state: IntcodeProcessState,
}

impl IntcodeProcess {
    pub fn new(initial_memory: &Vec<i32>) -> Self {
        return IntcodeProcess {
            ip: 0,
            memory: initial_memory.clone(),
            state: IntcodeProcessState::Initialized,
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

    pub fn run(
        &mut self,
        inputs: &mut VecDeque<i32>,
        outputs: &mut VecDeque<i32>,
    ) -> IntcodeProcessState {
        loop {
            self.run_step(inputs, outputs);

            if self.state != IntcodeProcessState::Running {
                break;
            }
        }

        return self.state;
    }

    fn run_step(&mut self, inputs: &mut VecDeque<i32>, outputs: &mut VecDeque<i32>) {
        let instruction = IntcodeInstruction::parse(self.memory[self.ip]);
        self.state = match instruction.opcode {
            1 => self.do_add(&instruction),
            2 => self.do_multiply(&instruction),
            3 => self.do_input(&instruction, inputs),
            4 => self.do_output(&instruction, outputs),
            5 => self.do_jump_if_true(&instruction),
            6 => self.do_jump_if_false(&instruction),
            7 => self.do_less_than(&instruction),
            8 => self.do_equals(&instruction),
            99 => IntcodeProcessState::Halted,
            _ => panic!("Unexpected opcode!"),
        };
    }

    pub fn get_memory(&self, address: usize) -> i32 {
        return self.memory[address];
    }

    pub fn set_value(&mut self, address: usize, value: i32) {
        self.memory[address] = value;
    }

    fn load_parameter(&self, instruction: &IntcodeInstruction, index: usize) -> i32 {
        let param = self.memory[self.ip + 1 + index];
        match instruction.get_mode(index) {
            0 => self.memory[param as usize],
            1 => param,
            _ => panic!("Unexpected mode!"),
        }
    }

    fn store_parameter(&mut self, instruction: &IntcodeInstruction, index: usize, value: i32) {
        let param = self.memory[self.ip + 1 + index];
        match instruction.get_mode(index) {
            0 => self.memory[param as usize] = value,
            _ => panic!("Unexpected mode!"),
        }
    }

    fn do_add(&mut self, instruction: &IntcodeInstruction) -> IntcodeProcessState {
        let a = self.load_parameter(instruction, 0);
        let b = self.load_parameter(instruction, 1);

        self.store_parameter(instruction, 2, a + b);
        self.ip += 4;

        return IntcodeProcessState::Running;
    }

    fn do_multiply(&mut self, instruction: &IntcodeInstruction) -> IntcodeProcessState {
        let a = self.load_parameter(instruction, 0);
        let b = self.load_parameter(instruction, 1);

        self.store_parameter(instruction, 2, a * b);
        self.ip += 4;

        return IntcodeProcessState::Running;
    }

    fn do_input(
        &mut self,
        instruction: &IntcodeInstruction,
        inputs: &mut VecDeque<i32>,
    ) -> IntcodeProcessState {
        if let Some(value) = inputs.pop_front() {
            self.store_parameter(instruction, 0, value);

            self.ip += 2;
            return IntcodeProcessState::Running;
        } else {
            return IntcodeProcessState::Waiting;
        }
    }

    fn do_output(
        &mut self,
        instruction: &IntcodeInstruction,
        outputs: &mut VecDeque<i32>,
    ) -> IntcodeProcessState {
        let value = self.load_parameter(instruction, 0);
        outputs.push_back(value);
        self.ip += 2;

        return IntcodeProcessState::Running;
    }

    fn do_jump_if_true(&mut self, instruction: &IntcodeInstruction) -> IntcodeProcessState {
        let a = self.load_parameter(instruction, 0);
        if a != 0 {
            self.ip = self.load_parameter(instruction, 1) as usize;
        } else {
            self.ip += 3;
        }

        return IntcodeProcessState::Running;
    }

    fn do_jump_if_false(&mut self, instruction: &IntcodeInstruction) -> IntcodeProcessState {
        let a = self.load_parameter(instruction, 0);
        if a == 0 {
            self.ip = self.load_parameter(instruction, 1) as usize;
        } else {
            self.ip += 3;
        }

        return IntcodeProcessState::Running;
    }

    fn do_less_than(&mut self, instruction: &IntcodeInstruction) -> IntcodeProcessState {
        let a = self.load_parameter(instruction, 0);
        let b = self.load_parameter(instruction, 1);

        let value = if a < b { 1 } else { 0 };
        self.store_parameter(instruction, 2, value);
        self.ip += 4;

        return IntcodeProcessState::Running;
    }

    fn do_equals(&mut self, instruction: &IntcodeInstruction) -> IntcodeProcessState {
        let a = self.load_parameter(instruction, 0);
        let b = self.load_parameter(instruction, 1);

        let value = if a == b { 1 } else { 0 };
        self.store_parameter(instruction, 2, value);
        self.ip += 4;

        return IntcodeProcessState::Running;
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
            let mut inputs = VecDeque::new();
            let mut outputs = VecDeque::new();
            let mut computer = IntcodeProcess::new(&case.0);
            computer.run(&mut inputs, &mut outputs);
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
