use aoc2025::days::day04::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2025/day04.txt");
    assert_eq!(part1(&content), 1564);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2025/day04.txt");
    assert_eq!(part2(&content), 9401);
}
