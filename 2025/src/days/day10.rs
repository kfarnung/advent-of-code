use std::{
    collections::{HashMap, VecDeque},
    vec,
};

pub fn part1(contents: &str) -> i64 {
    let mut count = 0;
    let diagrams = parse_input(contents);
    for diagram in &diagrams {
        let combinations = generate_combinations(&diagram.buttons);
        let possible_presses = combinations.get(&diagram.lights).unwrap();
        count += possible_presses
            .iter()
            .map(|v| v.len() as i64)
            .min()
            .unwrap();
    }
    count
}

pub fn part2(contents: &str) -> i64 {
    let mut count = 0;
    let diagrams = parse_input(contents);
    for diagram in &diagrams {
        let mut combinations = generate_combinations(&diagram.buttons);
        count += count_presses(&diagram.powers, &mut combinations).unwrap();
    }
    count
}

fn parse_input(input: &str) -> Vec<Diagram> {
    input.lines().map(|line| Diagram::from_line(line)).collect()
}

fn generate_combinations(buttons: &Vec<u32>) -> HashMap<u32, Vec<Vec<u32>>> {
    let mut map = HashMap::new();

    // Include the empty combination to handle the case of zero lights
    map.insert(0, vec![vec![]]);

    let mut queue = VecDeque::new();
    for (i, button) in buttons.iter().enumerate() {
        queue.push_back((*button, i + 1, vec![*button]));
    }

    while !queue.is_empty() {
        let (current_lights, index, presses) = queue.pop_front().unwrap();
        map.entry(current_lights)
            .or_insert_with(Vec::new)
            .push(presses.clone());

        for (i, button) in buttons.iter().skip(index).enumerate() {
            let mut presses = presses.clone();
            presses.push(*button);
            queue.push_back((current_lights ^ button, index + i + 1, presses));
        }
    }
    map
}

fn get_parity(powers: &Vec<i32>) -> u32 {
    let mut parity: u32 = 0;
    for (i, &power) in powers.iter().enumerate() {
        if power % 2 == 1 {
            parity |= 1 << i;
        }
    }
    parity
}

fn calculate_powers(powers: &Vec<i32>, buttons: &Vec<u32>) -> Option<Vec<i32>> {
    let mut powers = powers.clone();
    for (i, power) in powers.iter_mut().enumerate() {
        for button in buttons {
            if (button & (1 << i)) != 0 {
                *power -= 1;
                if *power < 0 {
                    // We've gone below zero, no need to continue
                    return None;
                }
            }
        }
        *power /= 2;
    }
    Some(powers)
}

fn found_solution(powers: &Vec<i32>) -> bool {
    for power in powers {
        if *power != 0 {
            return false;
        }
    }
    true
}

fn count_presses(powers: &Vec<i32>, map: &HashMap<u32, Vec<Vec<u32>>>) -> Option<i64> {
    if found_solution(powers) {
        return Some(0);
    }

    let parity = get_parity(powers);
    if let Some(possible_presses) = map.get(&parity) {
        let mut counts = vec![];
        for presses in possible_presses {
            if let Some(new_powers) = calculate_powers(powers, &presses) {
                if let Some(sub_count) = count_presses(&new_powers, map) {
                    counts.push(sub_count * 2 + presses.len() as i64);
                }
            }
        }

        if let Some(min) = counts.iter().min() {
            return Some(*min);
        }
    }
    None
}

struct Diagram {
    lights: u32,
    buttons: Vec<u32>,
    powers: Vec<i32>,
}

impl Diagram {
    fn from_line(line: &str) -> Self {
        let parts: Vec<&str> = line.split(' ').collect();
        let mut lights: u32 = 0;
        let mut buttons = vec![];
        let mut powers = vec![];
        for part in parts {
            if part.starts_with('[') {
                let mut shift = 0;
                for c in part.chars().skip(1).take(part.len() - 2) {
                    if c == '#' {
                        lights |= 1 << shift;
                    }
                    shift += 1;
                }
            } else if part.starts_with('(') {
                let button_values: Vec<u32> = part[1..part.len() - 1]
                    .split(',')
                    .map(|s| s.parse().unwrap())
                    .collect();
                let mut button: u32 = 0;
                for v in button_values {
                    button |= 1 << v;
                }
                buttons.push(button);
            } else if part.starts_with('{') {
                let power_values: Vec<i32> = part[1..part.len() - 1]
                    .split(',')
                    .map(|s| s.parse().unwrap())
                    .collect();
                powers.extend(power_values);
            }
        }
        Diagram {
            lights: lights,
            buttons: buttons,
            powers: powers,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
            "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
            "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
        ];

        assert_eq!(part1(&input.join("\n")), 7);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
            "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
            "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
        ];

        assert_eq!(part2(&input.join("\n")), 33);
    }
}
