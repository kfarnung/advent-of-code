use aoc2022::days::day04::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day04.txt");
    assert_eq!(part1(&content), 483);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day04.txt");
    assert_eq!(part2(&content), 874);
}
