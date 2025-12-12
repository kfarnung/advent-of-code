use aoc2025::days::day10::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2025/day10.txt");
    assert_eq!(part1(&content), 419);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2025/day10.txt");
    assert_eq!(part2(&content), 18369);
}
