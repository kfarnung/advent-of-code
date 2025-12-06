pub fn part1(contents: &str) -> i64 {
    let split_lines: Vec<Vec<&str>> = contents
        .lines()
        .map(|line| line.split_whitespace().collect())
        .collect();

    let mut parsed_input = Vec::new();
    for i in 0..split_lines[0].len() {
        let mut column = Vec::new();
        for j in 0..split_lines.len() - 1 {
            let cell = &split_lines[j][i];
            column.push(cell.parse::<i64>().unwrap());
        }
        parsed_input.push((
            column,
            split_lines[split_lines.len() - 1][i]
                .chars()
                .next()
                .unwrap(),
        ));
    }

    let mut result = 0;
    for (column, operation) in parsed_input {
        match operation {
            '+' => result += column.iter().sum::<i64>(),
            '*' => result += column.iter().product::<i64>(),
            _ => panic!("Unknown operation"),
        }
    }

    result
}

pub fn part2(contents: &str) -> i64 {
    let lines: Vec<Vec<char>> = contents
        .lines()
        .map(|line| line.chars().collect())
        .collect();

    let mut result = 0;
    let mut numbers = Vec::new();

    for i in (0..lines[0].len()).rev() {
        let mut found_number = false;
        let mut number = 0;
        let mut multiplier = 1;

        for j in (0..lines.len() - 1).rev() {
            if lines[j][i] == ' ' {
                continue;
            }
            found_number = true;
            number += (lines[j][i].to_digit(10).unwrap() as i64) * multiplier;
            multiplier *= 10;
        }

        if found_number {
            numbers.push(number);
        }

        if lines[lines.len() - 1][i] != ' ' {
            match lines[lines.len() - 1][i] {
                '+' => result += numbers.iter().sum::<i64>(),
                '*' => result += numbers.iter().product::<i64>(),
                _ => panic!("Unknown operation"),
            }
            numbers.clear();
        }
    }

    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            "123 328  51 64 ",
            " 45 64  387 23 ",
            "  6 98  215 314",
            "*   +   *   +  ",
        ];

        assert_eq!(part1(&input.join("\n")), 4277556);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            "123 328  51 64 ",
            " 45 64  387 23 ",
            "  6 98  215 314",
            "*   +   *   +  ",
        ];

        assert_eq!(part2(&input.join("\n")), 3263827);
    }
}
