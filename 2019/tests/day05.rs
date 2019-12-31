use aoc2019::days::day05::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day05");
    assert_eq!(part1(&content), 7157989);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day05");
    assert_eq!(part2(&content), 7873292);
}
