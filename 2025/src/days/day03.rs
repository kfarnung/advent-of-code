use std::vec;

pub fn part1(contents: &str) -> i64 {
    calculate_joltage(contents, 2)
}

pub fn part2(contents: &str) -> i64 {
    calculate_joltage(contents, 12)
}

fn calculate_joltage(contents: &str, digits: usize) -> i64 {
    let mut sum = 0;
    for l in contents.lines() {
        let nums: Vec<i64> = l.chars().map(|c| c.to_digit(10).unwrap() as i64).collect();
        let mut indices = vec![];
        find_max_recursive(&nums, 0, nums.len(), &mut indices, digits);
        indices.sort();

        sum += indices
            .iter()
            .map(|e| nums[*e])
            .reduce(|acc, e| acc * 10 + e)
            .unwrap();
    }
    sum
}

fn find_max_recursive(
    nums: &[i64],
    start: usize,
    end: usize,
    indices: &mut Vec<usize>,
    digits: usize,
) {
    if start == end || indices.len() == digits {
        return;
    }

    let max_index = find_max_index(nums, start, end);
    indices.push(max_index);
    find_max_recursive(nums, max_index + 1, end, indices, digits);
    find_max_recursive(nums, start, max_index, indices, digits);
}

fn find_max_index(nums: &[i64], start: usize, end: usize) -> usize {
    let mut max = -1;
    let mut max_index: usize = 0;
    for (i, n) in nums[start..end].iter().enumerate() {
        if *n > max {
            max = *n;
            max_index = start + i;
        }
    }
    max_index
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            "987654321111111",
            "811111111111119",
            "234234234234278",
            "818181911112111",
        ];

        assert_eq!(part1(&input.join("\n")), 357);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            "987654321111111",
            "811111111111119",
            "234234234234278",
            "818181911112111",
        ];

        assert_eq!(part2(&input.join("\n")), 3121910778619);
    }
}
