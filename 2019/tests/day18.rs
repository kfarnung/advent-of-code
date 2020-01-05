use aoc2019::days::day18::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day18");
    assert_eq!(part1(&content), 3832);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day18");
    assert_eq!(part2(&content), 1724);
}
