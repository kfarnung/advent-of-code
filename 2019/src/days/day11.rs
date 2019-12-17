use crate::shared::intcode::{IntcodeProcess, IntcodeProcessState};
use crate::shared::point::Point2D;
use std::collections::{HashMap, VecDeque};

pub fn part1(initial_memory: &str) -> usize {
    return paint(initial_memory, 0).len();
}

pub fn part2(initial_memory: &str) -> String {
    let grid = paint(initial_memory, 1);

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
                1 => "#",
                _ => panic!("Unexpected color"),
            });
        }

        lines.push(line.join(""));
    }

    return lines.join("\n");
}

fn paint(initial_memory: &str, starting_color: i64) -> HashMap<Point2D<i32>, i64> {
    let mut direction = 0;
    let mut position = Point2D::new(0, 0);
    let mut grid = HashMap::new();

    let mut input = VecDeque::new();
    let mut output = VecDeque::new();
    let mut process = IntcodeProcess::new_from_string(initial_memory);

    grid.insert(position.clone(), starting_color);

    loop {
        // Input the color of the current cell.
        let color = grid.get(&position).unwrap_or(&0);
        input.push_back(color.clone());
        // Run the process until is needs a new input.
        let state = process.run(&mut input, &mut output);

        if state == IntcodeProcessState::Halted {
            // The program exited.
            break;
        }
        // Update the color in the grid.
        let new_color = output.pop_front().unwrap();
        grid.insert(position.clone(), new_color);
        // Calculate the new direction.
        let new_direction = output.pop_front().unwrap();
        direction = turn(direction, new_direction);

        // Update the position.
        position += get_delta(direction);
    }

    return grid;
}

fn turn(current: usize, new: i64) -> usize {
    match new {
        0 => (((current as i32) + 3) % 4) as usize,
        1 => (((current as i32) + 1) % 4) as usize,
        _ => panic!("Invalid direction!"),
    }
}

fn get_delta(direction: usize) -> Point2D<i32> {
    match direction {
        0 => Point2D::new(0, -1),
        1 => Point2D::new(1, 0),
        2 => Point2D::new(0, 1),
        3 => Point2D::new(-1, 0),
        _ => panic!("Unexpected direction value"),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_turn_left() {
        assert_eq!(turn(0, 0), 3);
    }

    #[test]
    fn test_turn_right() {
        assert_eq!(turn(3, 1), 0);
    }
}
