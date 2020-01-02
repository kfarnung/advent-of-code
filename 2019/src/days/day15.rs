use crate::shared::intcode::IntcodeProcess;
use crate::shared::point::Point2D;
use std::collections::{HashMap, VecDeque};

pub fn part1(initial_memory: &str) -> i64 {
    let map = generate_map(initial_memory);
    let generator_location = find_oxygen_generator(&map);
    let distances = calculate_distances(&map, &generator_location);
    return distances.get(&Point2D::new(0, 0)).unwrap().clone();
}

pub fn part2(initial_memory: &str) -> i64 {
    let map = generate_map(initial_memory);
    let generator_location = find_oxygen_generator(&map);
    let distances = calculate_distances(&map, &generator_location);
    let max_point = distances.iter().max_by_key(|x| x.1).unwrap();
    return *max_point.1;
}

fn generate_map(initial_memory: &str) -> HashMap<Point2D<i64>, i64> {
    let mut input = VecDeque::new();
    let mut output = VecDeque::new();
    let mut process = IntcodeProcess::new_from_string(initial_memory);

    let mut visited = HashMap::new();
    let mut backtrack = VecDeque::new();

    // Start at the center of the universe.
    let mut is_backtrack = false;
    let mut direction = 1;
    let mut position = Point2D::new(0, 0);
    visited.insert(position.clone(), 1);

    loop {
        if direction <= 4 {
            let next_position = get_position(&position, direction);
            if visited.contains_key(&next_position) {
                // Already visited, try another direction.
                direction += 1;
                continue;
            }
        } else if !backtrack.is_empty() {
            // Backtrack to the last place we were and try again.
            direction = backtrack.pop_front().unwrap();
            is_backtrack = true;
        } else {
            // We've exhausted all options.
            break;
        }

        // Run the process until is needs a new input.
        input.push_back(direction);
        process.run(&mut input, &mut output);

        let result = output.pop_front().unwrap();

        let next_position = get_position(&position, direction);
        visited.insert(next_position, result);

        if result > 2 {
            panic!("Unexpected output!");
        } else if result > 0 {
            if !is_backtrack {
                backtrack.push_front(reverse_direction(direction));
            }
            is_backtrack = false;
            position = next_position;
            direction = 1;
        }
    }

    return visited;
}

fn find_oxygen_generator(map: &HashMap<Point2D<i64>, i64>) -> Point2D<i64> {
    return map
        .iter()
        .filter_map(|(k, v)| if *v == 2 { Some(k) } else { None })
        .next()
        .unwrap()
        .clone();
}

fn calculate_distances(
    map: &HashMap<Point2D<i64>, i64>,
    start: &Point2D<i64>,
) -> HashMap<Point2D<i64>, i64> {
    let mut distances = HashMap::new();
    let mut search_queue = VecDeque::new();

    distances.insert(start.clone(), 0);
    search_queue.push_back((start.clone(), 0));

    while !search_queue.is_empty() {
        let current = search_queue.pop_front().unwrap();

        for direction in 1..=4 {
            let next_position = get_position(&current.0, direction);
            let new_distance = current.1 + 1;

            if let Some(distance) = distances.get_mut(&next_position) {
                // We've already been here, see if we found a shorter path.
                if new_distance < *distance {
                    search_queue.push_back((next_position, new_distance));
                    *distance = new_distance;
                }
            } else if let Some(entity) = map.get(&next_position) {
                if *entity != 0 {
                    // We're not moving into a wall.
                    search_queue.push_back((next_position, new_distance));
                    distances.insert(next_position.clone(), new_distance);
                }
            }
        }
    }

    return distances;
}

fn reverse_direction(direction: i64) -> i64 {
    return match direction {
        1 => 2,
        2 => 1,
        3 => 4,
        4 => 3,
        _ => panic!("Invalid direction"),
    };
}

fn get_position(current: &Point2D<i64>, direction: i64) -> Point2D<i64> {
    let offset = match direction {
        1 => Point2D::new(0, -1),
        2 => Point2D::new(0, 1),
        3 => Point2D::new(-1, 0),
        4 => Point2D::new(1, 0),
        _ => panic!("Invalid direction"),
    };

    return *current + offset;
}
