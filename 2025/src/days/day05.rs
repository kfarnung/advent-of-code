pub fn part1(contents: &str) -> i64 {
    let (ranges, available) = parse_input(contents);
    let mut count = 0;
    for b in available {
        for r in &ranges {
            if b >= r.0 && b <= r.1 {
                count += 1;
                break;
            }
        }
    }
    count
}

pub fn part2(contents: &str) -> i64 {
    let (mut ranges, _) = parse_input(contents);
    ranges.sort_by(|a, b| a.0.cmp(&b.0));
    let mut merged_ranges = Vec::new();
    let mut current_merged = ranges[0];
    for r in ranges.iter().skip(1) {
        if r.0 <= current_merged.1 {
            current_merged.1 = current_merged.1.max(r.1);
        } else {
            merged_ranges.push(current_merged);
            current_merged = r.clone();
        }
    }
    merged_ranges.push(current_merged);

    let mut count = 0;
    for r in merged_ranges {
        count += r.1 - r.0 + 1;
    }
    count
}

fn parse_input(contents: &str) -> (Vec<(i64, i64)>, Vec<i64>) {
    let lines = contents.lines();
    let mut ranges = Vec::new();
    let mut available = Vec::new();
    let mut found_blank = false;

    for l in lines {
        if l.is_empty() {
            found_blank = true;
            continue;
        }

        if !found_blank {
            let parts: Vec<&str> = l.split('-').collect();
            let start = parts[0].parse::<i64>().unwrap();
            let end = parts[1].parse::<i64>().unwrap();
            ranges.push((start, end));
        } else {
            let blocked_value = l.parse::<i64>().unwrap();
            available.push(blocked_value);
        }
    }

    (ranges, available)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            "3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32",
        ];

        assert_eq!(part1(&input.join("\n")), 3);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            "3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32",
        ];

        assert_eq!(part2(&input.join("\n")), 14);
    }
}
