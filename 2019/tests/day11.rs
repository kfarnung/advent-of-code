use aoc2019::days::day11::{part1, part2};

const PART2_EXPECTED: &'static str = r" #### ####  ##  #  # #  # ####  ##   ##    
 #    #    #  # # #  #  # #    #  # #  #   
 ###  ###  #    ##   #  # ###  #    #      
 #    #    #    # #  #  # #    # ## #      
 #    #    #  # # #  #  # #    #  # #  #   
 #### #     ##  #  #  ##  ####  ###  ##    ";

#[test]
fn test_part1() {
    let content = std::include_str!("../../private/inputs/2019/day11.txt");
    assert_eq!(part1(&content), 2211);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../../private/inputs/2019/day11.txt");
    assert_eq!(part2(&content), PART2_EXPECTED);
}
