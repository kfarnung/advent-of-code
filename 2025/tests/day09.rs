use aoc2025::days::day09::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2025/day09.txt");
    assert_eq!(part1(&content), 4749672288);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2025/day09.txt");
    assert_eq!(part2(&content), 1479665889);
}
