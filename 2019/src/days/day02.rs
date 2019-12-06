use crate::shared::intcode::IntcodeComputer;

pub fn part1(initial_memory: &str) -> i32 {
    let mut computer = IntcodeComputer::new_from_string(initial_memory);
    computer.set_value(1, 12);
    computer.set_value(2, 2);
    computer.run();

    return computer.get_memory(0);
}

pub fn part2(initial_memory: &str) -> i32 {
    let mut computer = IntcodeComputer::new_from_string(initial_memory);

    for noun in 0..99 {
        for verb in 0..99 {
            computer.reset();
            computer.set_value(1, noun);
            computer.set_value(2, verb);
            computer.run();

            if computer.get_memory(0) == 19690720 {
                return 100 * noun + verb;
            }
        }
    }

    panic!("We didn't find it!");
}
