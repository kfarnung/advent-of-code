use std::collections::HashSet;

pub fn part1(contents: &str) -> usize {
    find_marker(contents, 4)
}

pub fn part2(contents: &str) -> usize {
    find_marker(contents, 14)
}

fn find_marker(contents: &str, length: usize) -> usize {
    for i in length - 1..contents.len() {
        let set: HashSet<char> = contents.get(i - (length - 1)..=i).unwrap().chars().collect();
        if set.len() == length {
            return i + 1;
        }
    }

    0
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_part1() {
        let cases = vec![
            ("bvwbjplbgvbhsrlpgdmjqwftvncz", 5),
            ("nppdvjthqldpwncqszvftbrmjlhg", 6),
            ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10),
            ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11),
        ];

        for case in cases {
            assert_eq!(part1(case.0), case.1);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![
            ("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19),
            ("bvwbjplbgvbhsrlpgdmjqwftvncz", 23),
            ("nppdvjthqldpwncqszvftbrmjlhg", 23),
            ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29),
            ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26),
        ];

        for case in cases {
            assert_eq!(part2(case.0), case.1);
        }
    }
}
