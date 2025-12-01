use std::str::FromStr;

pub fn part1(contents: &str) -> i64 {
    let mut current: i64 = 50;
    let mut count = 0;
    for (dir, dist) in contents.lines().map(|line| parse_line(line)) {
        match dir {
            'L' => current = (current - dist).rem_euclid(100),
            'R' => current = (current + dist).rem_euclid(100),
            _ => panic!("Invalid direction"),
        }

        if current == 0 {
            count += 1;
        }
    }
    count
}

pub fn part2(contents: &str) -> i64 {
    let mut current = 50;
    let mut count = 0;
    for (dir, dist) in contents.lines().map(|line| parse_line(line)) {
        match dir {
            'L' => {
                count += (current - dist).div_euclid(100).abs();
                if current == 0 && dist > 0 {
                    // Moving left again will wrap around and count as another zero
                    // crossing. Remove one here to avoid double-counting.
                    count -= 1;
                }

                current = (current - dist).rem_euclid(100);
                if current == 0 {
                    // The division doesn't count landing on zero as a crossing,
                    // so we need to increment the count here.
                    count += 1;
                }
            }
            'R' => {
                count += (current + dist).div_euclid(100);
                current = (current + dist).rem_euclid(100);
            }
            _ => panic!("Invalid direction"),
        }
    }
    count
}

fn parse_line(line: &str) -> (char, i64) {
    let (dir, dist) = line.split_at(1);
    (dir.chars().next().unwrap(), i64::from_str(dist).unwrap())
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_part1() {
        let lines = vec![
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];

        assert_eq!(part1(&lines.join("\n")), 3);
    }

    #[test]
    fn test_part2() {
        let lines = vec![
            "L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82",
        ];

        assert_eq!(part2(&lines.join("\n")), 6);
    }
}
