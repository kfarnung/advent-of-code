use aoc2022::days::day03::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day03.txt");
    assert_eq!(part1(&content), 8240);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day03.txt");
    assert_eq!(part2(&content), 2587);
}
