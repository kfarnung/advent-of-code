use crate::shared::intcode::IntcodeComputer;

pub fn part1(initial_memory: &Vec<usize>) -> usize {
    let mut computer = IntcodeComputer::new(initial_memory);
    computer.run(Some(12), Some(2));
    return computer.get_value(0);
}

pub fn part2(initial_memory: &Vec<usize>) -> usize {
    for noun in 0..99 {
        for verb in 0..99 {
            let mut computer = IntcodeComputer::new(initial_memory);
            computer.run(Some(noun), Some(verb));
            if computer.get_value(0) == 19690720 {
                return 100 * noun + verb;
            }
        }
    }

    panic!("We didn't find it!");
}
