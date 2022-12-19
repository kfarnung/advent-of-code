use regex::Regex;
use std::collections::VecDeque;

pub fn part1(contents: &str) -> String {
    let (mut stacks, moves) = parse(contents);

    for action in moves {
        for _x in 0..action.0 {
            let value = stacks[action.1 - 1].pop_front().unwrap();
            stacks[action.2 - 1].push_front(value);
        }
    }

    stack_tops(stacks)
}

pub fn part2(contents: &str) -> String {
    let (mut stacks, moves) = parse(contents);

    for action in moves {
        for x in 0..action.0 {
            let value = stacks[action.1 - 1][action.0 - x - 1];
            stacks[action.1 - 1].remove(action.0 - x - 1);
            stacks[action.2 - 1].push_front(value);
        }
    }

    stack_tops(stacks)
}

fn parse(contents: &str) -> (Vec<VecDeque<char>>, Vec<(usize, usize, usize)>) {
    let re = Regex::new(r"^move (\d+) from (\d+) to (\d+)$").unwrap();

    let mut stacks: Vec<VecDeque<char>> = vec![];
    let mut moves = vec![];
    let mut found_break = false;
    for line in contents.lines() {
        if line.is_empty() {
            found_break = true;
        } else if found_break {
            let caps = re.captures(line).unwrap();
            moves.push((
                caps[1].parse::<usize>().unwrap(),
                caps[2].parse::<usize>().unwrap(),
                caps[3].parse::<usize>().unwrap(),
            ));
        } else {
            for (i, char) in line.chars().skip(1).step_by(4).enumerate() {
                if char >= 'A' && char <= 'Z' {
                    while stacks.len() <= i {
                        stacks.push(VecDeque::new());
                    }

                    stacks[i].push_back(char);
                }
            }
        }
    }

    (stacks, moves)
}

fn stack_tops(stacks: Vec<VecDeque<char>>) -> String {
    stacks
        .iter()
        .map(|x| *x.front().unwrap())
        .collect::<String>()
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_part1() {
        let lines = vec![
            "    [D]    ",
            "[N] [C]    ",
            "[Z] [M] [P]",
            " 1   2   3 ",
            "",
            "move 1 from 2 to 1",
            "move 3 from 1 to 3",
            "move 2 from 2 to 1",
            "move 1 from 1 to 2",
        ];

        assert_eq!(part1(&lines.join("\n")), "CMZ");
    }

    #[test]
    fn test_part2() {
        let lines = vec![
            "    [D]    ",
            "[N] [C]    ",
            "[Z] [M] [P]",
            " 1   2   3 ",
            "",
            "move 1 from 2 to 1",
            "move 3 from 1 to 3",
            "move 2 from 2 to 1",
            "move 1 from 1 to 2",
        ];

        assert_eq!(part2(&lines.join("\n")), "MCD");
    }
}
