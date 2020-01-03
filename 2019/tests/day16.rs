use aoc2019::days::day16::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day16");
    assert_eq!(part1(&content), 11833188);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day16");
    assert_eq!(part2(&content), 55005000);
}
