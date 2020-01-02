use aoc2019::days::day15::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day15");
    assert_eq!(part1(&content), 404);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day15");
    assert_eq!(part2(&content), 406);
}
