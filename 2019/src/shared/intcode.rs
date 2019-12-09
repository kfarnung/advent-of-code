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
    opcode: i64,
    modes: Vec<i64>,
}

impl IntcodeInstruction {
    fn new(opcode: i64) -> Self {
        return IntcodeInstruction {
            opcode: opcode,
            modes: Vec::new(),
        };
    }

    pub fn parse(instruction: i64) -> Self {
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

    fn get_mode(&self, index: usize) -> i64 {
        let mode = self.modes.get(index);
        match mode {
            Some(x) => x.clone(),
            None => 0,
        }
    }
}

pub struct IntcodeProcess {
    ip: usize,
    memory: Vec<i64>,
    relative_base: i64,
    state: IntcodeProcessState,
}

impl IntcodeProcess {
    pub fn new(initial_memory: &Vec<i64>) -> Self {
        return IntcodeProcess {
            ip: 0,
            memory: initial_memory.clone(),
            relative_base: 0,
            state: IntcodeProcessState::Initialized,
        };
    }

    pub fn new_from_string(initial_memory: &str) -> Self {
        let initial_memory = initial_memory
            .trim()
            .split(',')
            .map(|i| i.parse::<i64>().unwrap())
            .collect();

        return Self::new(&initial_memory);
    }

    pub fn execute(initial_memory: &str, input: i64) -> i64 {
        let mut computer = IntcodeProcess::new_from_string(initial_memory);
        let mut inputs = VecDeque::new();
        let mut outputs = VecDeque::new();
    
        inputs.push_back(input);
        computer.run(&mut inputs, &mut outputs);
    
        return outputs.pop_back().unwrap();
    }

    pub fn run(
        &mut self,
        inputs: &mut VecDeque<i64>,
        outputs: &mut VecDeque<i64>,
    ) -> IntcodeProcessState {
        loop {
            self.run_step(inputs, outputs);

            if self.state != IntcodeProcessState::Running {
                break;
            }
        }

        return self.state;
    }

    fn run_step(&mut self, inputs: &mut VecDeque<i64>, outputs: &mut VecDeque<i64>) {
        let instruction = IntcodeInstruction::parse(self.get_memory(self.ip));
        self.state = match instruction.opcode {
            1 => self.do_add(&instruction),
            2 => self.do_multiply(&instruction),
            3 => self.do_input(&instruction, inputs),
            4 => self.do_output(&instruction, outputs),
            5 => self.do_jump_if_true(&instruction),
            6 => self.do_jump_if_false(&instruction),
            7 => self.do_less_than(&instruction),
            8 => self.do_equals(&instruction),
            9 => self.do_adjust_relative_base(&instruction),
            99 => IntcodeProcessState::Halted,
            _ => panic!("Unexpected opcode!"),
        };
    }

    pub fn get_memory(&self, address: usize) -> i64 {
        if address < self.memory.len() {
            return self.memory[address];
        } else {
            return 0;
        }
    }

    pub fn set_value(&mut self, address: usize, value: i64) {
        if address >= self.memory.len() {
            self.memory.resize_with(address + 1, Default::default);
        }

        self.memory[address] = value;
    }

    fn load_parameter(&self, instruction: &IntcodeInstruction, index: usize) -> i64 {
        let param = self.get_memory(self.ip + 1 + index);
        match instruction.get_mode(index) {
            0 => self.get_memory(param as usize),
            1 => param,
            2 => self.get_memory((self.relative_base + param) as usize),
            _ => panic!("Unexpected mode!"),
        }
    }

    fn store_parameter(&mut self, instruction: &IntcodeInstruction, index: usize, value: i64) {
        let param = self.get_memory(self.ip + 1 + index);
        match instruction.get_mode(index) {
            0 => self.set_value(param as usize, value),
            2 => self.set_value((self.relative_base + param) as usize, value),
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
        inputs: &mut VecDeque<i64>,
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
        outputs: &mut VecDeque<i64>,
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

    fn do_adjust_relative_base(&mut self, instruction: &IntcodeInstruction) -> IntcodeProcessState {
        self.relative_base += self.load_parameter(instruction, 0);
        self.ip += 2;

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
    fn test_output() {
        let cases = vec![
            (
                vec![
                    109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
                ],
                vec![
                    109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
                ],
            ),
            (
                vec![1102, 34915192, 34915192, 7, 4, 7, 99, 0],
                vec![1219070632396864],
            ),
            (vec![104, 1125899906842624, 99], vec![1125899906842624]),
        ];

        for case in cases {
            let mut inputs = VecDeque::new();
            let mut outputs = VecDeque::new();
            let mut computer = IntcodeProcess::new(&case.0);
            computer.run(&mut inputs, &mut outputs);
            assert_eq!(outputs, case.1);
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
