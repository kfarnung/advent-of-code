use crate::shared::intcode::IntcodeProcess;
use std::collections::VecDeque;

pub fn part1(initial_memory: &str) -> i32 {
    let mut inputs = VecDeque::new();
    let mut outputs = VecDeque::new();

    inputs.push_back(1);
    let mut computer = IntcodeProcess::new_from_string(initial_memory);
    computer.run(&mut inputs, &mut outputs);

    let outputs_vec: Vec<&i32> = outputs.iter().collect();
    if let Some((last, elements)) = outputs_vec.split_last() {
        for element in elements {
            if **element != 0 {
                panic!("Unexpected output!");
            }
        }

        return *last.clone();
    }

    panic!("No output was found!");
}

pub fn part2(initial_memory: &str) -> i32 {
    let mut inputs = VecDeque::new();
    let mut outputs = VecDeque::new();
    
    inputs.push_back(5);
    let mut computer = IntcodeProcess::new_from_string(initial_memory);
    computer.run(&mut inputs, &mut outputs);

    let outputs_vec: Vec<&i32> = outputs.iter().collect();
    if let Some((last, elements)) = outputs_vec.split_last() {
        for element in elements {
            if **element != 0 {
                panic!("Unexpected output!");
            }
        }

        return *last.clone();
    }

    panic!("No output was found!");
}
