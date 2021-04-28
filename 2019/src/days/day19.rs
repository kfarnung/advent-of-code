use crate::shared::intcode::{IntcodeProcess, IntcodeProcessState};
use std::collections::VecDeque;

fn run_program(initial_memory: &str, x: i64, y: i64) -> i64 {
    let mut input = VecDeque::new();
    let mut output = VecDeque::new();

    input.push_back(x);
    input.push_back(y);

    let mut process = IntcodeProcess::new_from_string(initial_memory);
    let state = process.run(&mut input, &mut output);
    if state != IntcodeProcessState::Halted {
        // The program exited.
        panic!("The program didn't exit!");
    }

    let result = output.pop_front().unwrap();

    println!("({}, {}) => {}", x, y, result);

    return result;
}

pub fn part1(initial_memory: &str) -> i64 {
    let mut sum = 0;

    for x in 0..50 {
        for y in 0..50 {
            sum += run_program(initial_memory, x, y)
        }
    }

    return sum;
}

pub fn part2(initial_memory: &str) -> i64 {
    let mut x = 0;
    let mut y = 99;

    loop {
        while run_program(initial_memory, x, y) == 0 {
            x += 1;
        }

        if run_program(initial_memory, x + 99, y - 99) == 1 {
            return x * 10000 + (y - 99);
        }

        y += 1;
    }
}
