use crate::shared::intcode::{IntcodeProcess, IntcodeProcessState};
use crate::shared::point::Point2D;
use std::collections::{HashMap, VecDeque};

pub fn part1(initial_memory: &str) -> usize {
    let mut input = VecDeque::new();
    let mut output = VecDeque::new();
    let mut process = IntcodeProcess::new_from_string(initial_memory);

    let mut display = HashMap::new();

    loop {
        // Run the process until is needs a new input.
        let state = process.run(&mut input, &mut output);

        while !output.is_empty() {
            let x = output.pop_front().unwrap();
            let y = output.pop_front().unwrap();
            let tile_id = output.pop_front().unwrap();

            if x == -1 && y == 0 {
                println!("{}", tile_id);
            } else {
                display.insert(Point2D::new(x, y), tile_id);
            }
        }

        if state == IntcodeProcessState::Halted {
            // The program exited.
            break;
        }
    }

    return display.iter().filter(|x| *x.1 == 2).count();
}

pub fn part2(initial_memory: &str) -> i64 {
    let mut input = VecDeque::new();
    let mut output = VecDeque::new();
    let mut display = HashMap::new();
    let mut score = 0;

    let mut ball = Point2D::new(0, 0);
    let mut paddle = Point2D::new(0, 0);

    let mut process = IntcodeProcess::new_from_string(initial_memory);
    process.set_memory(0, 2);

    loop {
        // Run the process until is needs a new input.
        let state = process.run(&mut input, &mut output);

        while !output.is_empty() {
            let x = output.pop_front().unwrap();
            let y = output.pop_front().unwrap();
            let tile_id = output.pop_front().unwrap();

            match tile_id {
                3 => paddle = Point2D::new(x, y),
                4 => ball = Point2D::new(x, y),
                _ => (),
            };

            if x == -1 && y == 0 {
                score = tile_id;
            } else {
                display.insert(Point2D::new(x, y), tile_id);
            }
        }

        if input.is_empty() {
            input.push_back(next_move(&ball, &paddle));
        }

        if state == IntcodeProcessState::Halted {
            // The program exited.
            break;
        }
    }

    return score;
}

fn next_move(ball: &Point2D<i64>, paddle: &Point2D<i64>) -> i64 {
    return num::clamp(ball.x - paddle.x, -1, 1);
}
