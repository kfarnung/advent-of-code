use aoc2022::days::day05::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day05.txt");
    assert_eq!(part1(&content), "ZBDRNPMVH");
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day05.txt");
    assert_eq!(part2(&content), "WDLPFNNNB");
}
