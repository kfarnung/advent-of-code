use std::collections::HashSet;


const CODE_LOWERCASE_A: u8 = 97;
const CODE_UPPERCASE_A: u8 = 65;

pub fn part1(contents: &str) -> i64 {
    let mut score: i64 = 0;
    for line in contents.lines() {
        let mut compartment1 = HashSet::new();
        let mut compartment2 = HashSet::new();
        let middle = line.len() / 2;
        for (index, char) in line.chars().enumerate() {
            if index < middle {
                compartment1.insert(char);
            } else {
                compartment2.insert(char);
            }
        }

        let intersection = compartment1.intersection(&compartment2);
        for i in intersection {
            score += score_char(i);
        }
    }

    score
}

pub fn part2(contents: &str) -> i64 {
    let mut score: i64 = 0;
    let mut current = HashSet::new();
    for (index, line) in contents.lines().enumerate() {
        let rucksack: HashSet<_> = line.chars().collect();

        if index % 3 == 0 {
            current = rucksack;
        } else {
            current = current.intersection(&rucksack).map(|x| *x).collect();
        }

        if index % 3 == 2 {
            let list: Vec<&char> = current.iter().collect();
            score += score_char(list.first().unwrap());
        }
    }

    score
}

fn get_ascii_code(char: &char) -> u8 {
    let mut code = [0; 1];
    char.encode_utf8(&mut code);
    code[0]
}

fn score_char(char: &char) -> i64 {
    let code = get_ascii_code(char);
    if code >= CODE_LOWERCASE_A {
        i64::from(code - CODE_LOWERCASE_A + 1)
    } else if code >= CODE_UPPERCASE_A {
        i64::from(code - CODE_UPPERCASE_A + 27)
    } else {
        panic!("Unexpected character found!");
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_part1() {
        let lines = vec![
            "vJrwpWtwJgWrhcsFMMfFFhFp",
            "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
            "PmmdzqPrVvPwwTWBwg",
            "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
            "ttgJtRGJQctTZtZT",
            "CrZsJsPPZsGzwwsLwLmpwMDw",
        ];

        assert_eq!(part1(&lines.join("\n")), 157);
    }

    #[test]
    fn test_part2() {
        let lines = vec![
            "vJrwpWtwJgWrhcsFMMfFFhFp",
            "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
            "PmmdzqPrVvPwwTWBwg",
            "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
            "ttgJtRGJQctTZtZT",
            "CrZsJsPPZsGzwwsLwLmpwMDw",
        ];

        assert_eq!(part2(&lines.join("\n")), 70);
    }
}
