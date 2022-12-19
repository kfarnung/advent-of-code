use aoc2022::days::day06::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day06.txt");
    assert_eq!(part1(&content), 1262);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day06.txt");
    assert_eq!(part2(&content), 3444);
}
