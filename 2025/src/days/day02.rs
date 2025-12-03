use std::str::FromStr;

pub fn part1(contents: &str) -> i64 {
    let mut sum = 0;
    for (start, end) in parse_data(contents) {
        for i in start..=end {
            let num_str = i.to_string();
            let split = num_str.split_at(num_str.len() / 2);
            if split.0 == split.1 {
                sum += i;
            }
        }
    }
    sum
}

pub fn part2(contents: &str) -> i64 {
    let mut sum = 0;
    for (start, end) in parse_data(contents) {
        for i in start..=end {
            let num_str = i.to_string();
            if match_pattern(&num_str) {
                sum += i;
            }
        }
    }
    sum
}

fn parse_data(contents: &str) -> Vec<(i64, i64)> {
    contents
        .trim()
        .split(',')
        .map(|range| {
            let mut bounds = range.split('-').map(|n| i64::from_str(n).unwrap());
            let start = bounds.next().unwrap();
            let end = bounds.next().unwrap();
            (start, end)
        })
        .collect()
}

fn match_pattern(s: &str) -> bool {
    for i in 1..=s.len() / 2 {
        if s.len() % i != 0 {
            continue;
        }

        let comparison = &s[0..i];
        for j in (i..s.len()).step_by(i) {
            let end = j + i;
            if &s[j..end] != comparison {
                break;
            }
            if end == s.len() {
                return true;
            }
        }
    }
    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let content = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

        assert_eq!(part1(&content), 1227775554);
    }

    #[test]
    fn test_part2() {
        let content = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

        assert_eq!(part2(&content), 4174379265);
    }
}
