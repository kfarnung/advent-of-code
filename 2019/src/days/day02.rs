use crate::shared::intcode;

pub fn part1(initial_memory: &Vec<usize>) -> usize {
    let memory = intcode::run_program(initial_memory, Some(12), Some(2));
    return memory[0];
}

pub fn part2(initial_memory: &Vec<usize>) -> usize {
    for noun in 0..99 {
        for verb in 0..99 {
            let memory = intcode::run_program(initial_memory, Some(noun), Some(verb));
            if memory[0] == 19690720 {
                return 100 * noun + verb;
            }
        }
    }

    panic!("We didn't find it!");
}
