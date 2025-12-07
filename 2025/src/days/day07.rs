use std::collections::{HashMap, HashSet, VecDeque};

pub fn part1(contents: &str) -> i64 {
    let grid = parse_grid(contents);
    let start = grid[0]
        .iter()
        .enumerate()
        .filter(|(_, c)| *c == &'S')
        .next()
        .unwrap();
    let mut count = 0;
    let mut queue = VecDeque::new();
    let mut visited = HashSet::new();

    let start = (0 as usize, start.0);
    queue.push_back(start);
    visited.insert(start);

    while !queue.is_empty() {
        let (y, x) = queue.pop_front().unwrap();
        if y == grid.len() - 1 {
            break;
        }

        match grid[y + 1][x] {
            '.' => {
                let next = (y + 1, x);
                if !visited.contains(&next) {
                    queue.push_back(next);
                    visited.insert(next);
                }
            }
            '^' => {
                let left = (y + 1, x - 1);
                if !visited.contains(&left) {
                    queue.push_back(left);
                    visited.insert(left);
                }
                let right = (y + 1, x + 1);
                if !visited.contains(&right) {
                    queue.push_back(right);
                    visited.insert(right);
                }
                count += 1;
            }
            _ => {
                panic!("Unexpected character")
            }
        }
    }
    count
}

pub fn part2(contents: &str) -> i64 {
    let grid = parse_grid(contents);
    let start = grid[0]
        .iter()
        .enumerate()
        .filter(|(_, c)| *c == &'S')
        .map(|(i, _)| (0 as usize, i))
        .next()
        .unwrap();
    let mut memo = HashMap::new();
    count_recursive(&grid, start, &mut memo)
}

fn parse_grid(contents: &str) -> Vec<Vec<char>> {
    contents
        .lines()
        .map(|line| line.chars().collect())
        .collect()
}

fn count_recursive(
    grid: &Vec<Vec<char>>,
    next: (usize, usize),
    memo: &mut HashMap<(usize, usize), i64>,
) -> i64 {
    if memo.contains_key(&next) {
        return memo.get(&next).unwrap().clone();
    }

    if next.0 == grid.len() - 1 {
        return 1;
    }

    let mut count = 0;
    match grid[next.0][next.1] {
        'S' => {
            let next = (next.0 + 1, next.1);
            count += count_recursive(grid, next, memo);
        }
        '^' => {
            let left = (next.0 + 1, next.1 - 1);
            let right = (next.0 + 1, next.1 + 1);
            count += count_recursive(grid, left, memo);
            count += count_recursive(grid, right, memo);
        }
        '.' => {
            let next = (next.0 + 1, next.1);
            count += count_recursive(grid, next, memo);
        }
        _ => {
            panic!("Unexpected character");
        }
    }

    memo.insert(next, count);
    count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            ".......S.......",
            "...............",
            ".......^.......",
            "...............",
            "......^.^......",
            "...............",
            ".....^.^.^.....",
            "...............",
            "....^.^...^....",
            "...............",
            "...^.^...^.^...",
            "...............",
            "..^...^.....^..",
            "...............",
            ".^.^.^.^.^...^.",
            "...............",
        ];

        assert_eq!(part1(&input.join("\n")), 21);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            ".......S.......",
            "...............",
            ".......^.......",
            "...............",
            "......^.^......",
            "...............",
            ".....^.^.^.....",
            "...............",
            "....^.^...^....",
            "...............",
            "...^.^...^.^...",
            "...............",
            "..^...^.....^..",
            "...............",
            ".^.^.^.^.^...^.",
            "...............",
        ];

        assert_eq!(part2(&input.join("\n")), 40);
    }
}
