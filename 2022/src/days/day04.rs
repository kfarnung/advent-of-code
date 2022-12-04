use std::str::FromStr;

pub fn part1(contents: &str) -> i64 {
    let mut count = 0;
    for pair in contents.lines().map(|x| parse(x)) {
        if pair.0 >= pair.2 && pair.1 <= pair.3 {
            count += 1;
        } else if pair.2 >= pair.0 && pair.3 <= pair.1 {
            count += 1;
        }
    }

    count
}

pub fn part2(contents: &str) -> i64 {
    let mut count = 0;
    for pair in contents.lines().map(|x| parse(x)) {
        if (pair.0 >= pair.2 && pair.0 <= pair.3)
            || (pair.2 >= pair.0 && pair.2 <= pair.1)
        {
            count += 1;
        }
    }

    count
}

fn parse(line: &str) -> (i32, i32, i32, i32) {
    let ranges: Vec<&str> = line.split(',').collect();
    let first: Vec<&str> = ranges[0].split('-').collect();
    let second: Vec<&str> = ranges[1].split('-').collect();

    (
        i32::from_str(first[0]).unwrap(),
        i32::from_str(first[1]).unwrap(),
        i32::from_str(second[0]).unwrap(),
        i32::from_str(second[1]).unwrap(),
    )
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_part1() {
        let lines = vec![
            "2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8",
        ];

        assert_eq!(part1(&lines.join("\n")), 2);
    }

    #[test]
    fn test_part2() {
        let lines = vec![
            "2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8",
        ];

        assert_eq!(part2(&lines.join("\n")), 4);
    }
}
