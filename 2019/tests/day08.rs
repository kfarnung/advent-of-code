use aoc2019::days::day08::{part1, part2};

const PART2_EXPECTED: &'static str = r"###   ##  #  # #     ##  
#  # #  # #  # #    #  # 
#  # #    #  # #    #  # 
###  #    #  # #    #### 
#    #  # #  # #    #  # 
#     ##   ##  #### #  # ";

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2019/day08.txt");
    assert_eq!(part1(&content), 1920);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2019/day08.txt");
    assert_eq!(part2(&content), PART2_EXPECTED);
}
