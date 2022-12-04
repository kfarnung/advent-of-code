pub fn part1(contents: &str) -> i64 {
    let mut score = 0;
    for line in contents.lines() {
        let round: Vec<&str> = line.split(' ').collect();
        score += score_round(round[0], round[1]);
        score += score_play(round[1]);
    }

    score
}

pub fn part2(contents: &str) -> i64 {
    let mut score = 0;
    for line in contents.lines() {
        let round: Vec<&str> = line.split(' ').collect();
        if round[1] == "X" {
            score += ((score_play(round[0]) + 1) % 3) + 1
        } else if round[1] == "Y" {
            score += 3;
            score += score_play(round[0]);
        } else if round[1] == "Z" {
            score += 6;
            score += (score_play(round[0]) % 3) + 1
        } else {
            panic!("Unexpected outcome!");
        }
    }

    score
}

fn score_round(opponent: &str, you: &str) -> i64 {
    if (opponent == "A" && you == "Y")
        || (opponent == "B" && you == "Z")
        || (opponent == "C" && you == "X")
    {
        6
    } else if (opponent == "A" && you == "X")
        || (opponent == "B" && you == "Y")
        || (opponent == "C" && you == "Z")
    {
        3
    } else {
        0
    }
}

fn score_play(you: &str) -> i64 {
    match you {
        "A" => 1,
        "B" => 2,
        "C" => 3,
        "X" => 1,
        "Y" => 2,
        "Z" => 3,
        _ => panic!("Unexpected input!"),
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn test_part1() {
        let lines = vec!["A Y", "B X", "C Z"];

        assert_eq!(part1(&lines.join("\n")), 15);
    }

    #[test]
    fn test_part2() {
        let lines = vec!["A Y", "B X", "C Z"];

        assert_eq!(part2(&lines.join("\n")), 12);
    }
}
