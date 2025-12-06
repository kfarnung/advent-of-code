use aoc2025::days::day05::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2025/day05.txt");
    assert_eq!(part1(&content), 679);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2025/day05.txt");
    assert_eq!(part2(&content), 358155203664116);
}
