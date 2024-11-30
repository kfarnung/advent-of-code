use aoc2019::days::day12::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2019/day12.txt");
    assert_eq!(part1(&content), 12053);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2019/day12.txt");
    assert_eq!(part2(&content), 320380285873116);
}
