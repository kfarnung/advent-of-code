use aoc2022::days::day01::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2022/day01.txt");
    assert_eq!(part1(&content), 71780);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2022/day01.txt");
    assert_eq!(part2(&content), 212489);
}
