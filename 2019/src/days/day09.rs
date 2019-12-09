use crate::shared::intcode::IntcodeProcess;

pub fn part1(initial_memory: &str) -> i64 {
    return IntcodeProcess::execute(initial_memory, 1);
}

pub fn part2(initial_memory: &str) -> i64 {
    return IntcodeProcess::execute(initial_memory, 2);
}
