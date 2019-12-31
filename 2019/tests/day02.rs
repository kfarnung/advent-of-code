use aoc2019::days::day02::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day02");
    assert_eq!(part1(&content), 8017076);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day02");
    assert_eq!(part2(&content), 3146);
}
