use crate::shared::intcode::{IntcodeProcess, IntcodeProcessState};
use std::char;
use std::collections::VecDeque;

pub fn part1(initial_memory: &str) -> usize {
    let mut input = VecDeque::new();
    let mut output = VecDeque::new();
    let mut process = IntcodeProcess::new_from_string(initial_memory);

    let mut sum = 0;

    // Run the process until is needs a new input.
    let state = process.run(&mut input, &mut output);
    if state != IntcodeProcessState::Halted {
        // The program exited.
        panic!("The program didn't exit!");
    }

    let map = render(&output);
    output.clear();

    let lines: Vec<Vec<char>> = map.trim().lines().map(|x| x.chars().collect()).collect();
    let line_count = lines.len();
    let line_len = lines[0].len();

    for i in 1..line_count - 1 {
        for j in 1..line_len - 1 {
            if lines[i][j] == '#'
                && lines[i - 1][j] == '#'
                && lines[i][j - 1] == '#'
                && lines[i + 1][j] == '#'
                && lines[i][j + 1] == '#'
            {
                sum += i * j;
            }
        }
    }

    return sum;
}

pub fn part2(initial_memory: &str) -> i64 {
    let mut output = VecDeque::new();
    let mut process = IntcodeProcess::new_from_string(initial_memory);
    process.set_memory(0, 2);

    let instructions = vec![
        "A,A,B,C,B,C,B,C,C,A\n",
        "L,10,R,8,R,8\n",
        "L,10,L,12,R,8,R,10\n",
        "R,10,L,12,R,10\n",
        "n\n",
    ];
    let mut input: VecDeque<i64> = instructions.join("").chars().map(|x| x as i64).collect();

    // Run the process until is needs a new input.
    let state = process.run(&mut input, &mut output);
    if state != IntcodeProcessState::Halted {
        // The program exited.
        panic!("The program didn't exit!");
    }

    return output.pop_back().unwrap();
}

fn render(output: &VecDeque<i64>) -> String {
    return output
        .iter()
        .map(|x| char::from_u32(*x as u32).unwrap())
        .collect();
}
