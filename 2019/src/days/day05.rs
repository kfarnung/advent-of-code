use crate::shared::intcode::IntcodeComputer;

pub fn part1(initial_memory: &str) -> i32 {
    let mut computer = IntcodeComputer::new_from_string(initial_memory);
    computer.set_inputs(&vec![1]);
    computer.run();

    let outputs = computer.get_outputs();

    if let Some((last, elements)) = outputs.split_last() {
        for element in elements {
            if element != &0 {
                panic!("Unexpected output!");
            }
        }

        return last.clone();
    }

    panic!("No output was found!");
}

pub fn part2(initial_memory: &str) -> i32 {
    let mut computer = IntcodeComputer::new_from_string(initial_memory);
    computer.set_inputs(&vec![5]);
    computer.run();

    let outputs = computer.get_outputs();

    if let Some((last, elements)) = outputs.split_last() {
        for element in elements {
            if element != &0 {
                panic!("Unexpected output!");
            }
        }

        return last.clone();
    }

    panic!("No output was found!");
}
