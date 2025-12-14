use aoc2025::days::day11::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2025/day11.txt");
    assert_eq!(part1(&content), 658);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2025/day11.txt");
    assert_eq!(part2(&content), 371113003846800);
}
