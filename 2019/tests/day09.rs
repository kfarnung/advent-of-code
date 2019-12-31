use aoc2019::days::day09::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day09");
    assert_eq!(part1(&content), 2789104029);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day09");
    assert_eq!(part2(&content), 32869);
}
