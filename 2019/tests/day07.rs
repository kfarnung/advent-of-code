use aoc2019::days::day07::{part1, part2};

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2019/day07.txt");
    assert_eq!(part1(&content), 368584);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2019/day07.txt");
    assert_eq!(part2(&content), 35993240);
}
