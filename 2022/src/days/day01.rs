use std::{
    cmp::Reverse,
    str::{FromStr, Lines},
};

pub fn part1(contents: &str) -> i64 {
    *find_totals(contents.lines()).first().unwrap()
}

pub fn part2(contents: &str) -> i64 {
    find_totals(contents.lines())
        .get(0..3)
        .unwrap()
        .iter()
        .fold(0, |sum, x| sum + x)
}

fn find_totals(lines: Lines<'_>) -> Vec<i64> {
    let mut totals = vec![];
    let mut sum = 0;
    for line in lines {
        if line.is_empty() {
            totals.push(sum);
            sum = 0;
        } else {
            sum += i64::from_str(line).unwrap();
        }
    }

    totals.push(sum);
    totals.sort_by_key(|w| Reverse(*w));
    totals
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_part1() {
        let lines = vec![
            "1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "",
            "10000",
        ];

        assert_eq!(part1(&lines.join("\n")), 24000);
    }

    #[test]
    fn test_part2() {
        let lines = vec![
            "1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "",
            "10000",
        ];

        assert_eq!(part2(&lines.join("\n")), 45000);
    }
}
