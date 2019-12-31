use aoc2019::days::day11::{part1, part2};

const PART2_EXPECTED: &'static str = r" #### ####  ##  #  # #  # ####  ##   ##    
 #    #    #  # # #  #  # #    #  # #  #   
 ###  ###  #    ##   #  # ###  #    #      
 #    #    #    # #  #  # #    # ## #      
 #    #    #  # # #  #  # #    #  # #  #   
 #### #     ##  #  #  ##  ####  ###  ##    ";

#[test]
fn test_part1() {
    let content = std::include_str!("../data/day11");
    assert_eq!(part1(&content), 2211);
}

#[test]
fn test_part2() {
    let content = std::include_str!("../data/day11");
    assert_eq!(part2(&content), PART2_EXPECTED);
}
