use crate::shared::intcode::IntcodeProcess;
use std::collections::VecDeque;

pub fn part1(initial_memory: &str) -> i32 {
    let mut inputs = VecDeque::new();
    let mut outputs = VecDeque::new();
    let mut computer = IntcodeProcess::new_from_string(initial_memory);
    computer.set_value(1, 12);
    computer.set_value(2, 2);
    computer.run(&mut inputs, &mut outputs);

    return computer.get_memory(0);
}

pub fn part2(initial_memory: &str) -> i32 {
    for noun in 0..99 {
        for verb in 0..99 {
            let mut inputs = VecDeque::new();
            let mut outputs = VecDeque::new();
            let mut computer = IntcodeProcess::new_from_string(initial_memory);

            computer.set_value(1, noun);
            computer.set_value(2, verb);
            computer.run(&mut inputs, &mut outputs);

            if computer.get_memory(0) == 19690720 {
                return 100 * noun + verb;
            }
        }
    }

    panic!("We didn't find it!");
}
