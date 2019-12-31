use aoc2019::days::day03::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day03");
    assert_eq!(part1(&content), 1211);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day03");
    assert_eq!(part2(&content), 101386);
}
