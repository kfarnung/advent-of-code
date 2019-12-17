use crate::shared::intcode::{IntcodeProcess, IntcodeProcessState};
use crate::shared::point::Point2D;
use ncurses::*;
use std::collections::{HashMap, VecDeque};
use std::{thread, time};

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

    // Render the progress using ncurses
    initscr();

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

        // Clear the screen and write the new output.
        clear();
        addstr(&format!("{}\n", render_display(&display)));
        addstr(&format!("Score: {}\n", score));
        refresh();

        let sleep_time = time::Duration::from_millis(10);
        thread::sleep(sleep_time);

        if state == IntcodeProcessState::Halted {
            // The program exited.
            break;
        }
    }

    // All done here!
    endwin();

    return score;
}

fn next_move(ball: &Point2D<i64>, paddle: &Point2D<i64>) -> i64 {
    return num::clamp(ball.x - paddle.x, -1, 1);
}

fn render_display(grid: &HashMap<Point2D<i64>, i64>) -> String {
    // Find the boundaries of the painting.
    let min_x = grid.keys().min_by_key(|x| x.x).unwrap().x;
    let min_y = grid.keys().min_by_key(|x| x.y).unwrap().y;
    let max_x = grid.keys().max_by_key(|x| x.x).unwrap().x;
    let max_y = grid.keys().max_by_key(|x| x.y).unwrap().y;

    let mut lines = Vec::new();

    for y in min_y..=max_y {
        let mut line = Vec::new();

        for x in min_x..=max_x {
            let point = Point2D::new(x, y);
            let color = grid.get(&point).unwrap_or(&0);

            line.push(match color {
                0 => " ",
                1 => "|",
                2 => "#",
                3 => "-",
                4 => "*",
                _ => panic!("Unexpected color"),
            });
        }

        lines.push(line.join(""));
    }

    return lines.join("\n");
}
