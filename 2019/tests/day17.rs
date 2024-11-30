use aoc2019::days::day17::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2019/day17.txt");
    assert_eq!(part1(&content), 6212);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2019/day17.txt");
    assert_eq!(part2(&content), 1016741);
}
