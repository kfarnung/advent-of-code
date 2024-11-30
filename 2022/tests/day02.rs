use aoc2022::days::day02::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2019/day02.txt.txt");
    assert_eq!(part1(&content), 12156);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2019/day02.txt.txt");
    assert_eq!(part2(&content), 10835);
}
