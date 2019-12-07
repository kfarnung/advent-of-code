use crate::shared::intcode::IntcodeProcess;
use crate::shared::intcode::IntcodeProcessState;
use permutohedron::heap_recursive;
use std::cmp::max;
use std::collections::VecDeque;

pub fn part1(initial_memory: &str) -> i32 {
    let mut phase_settings = vec![0, 1, 2, 3, 4];
    let mut max_value = 0;
    heap_recursive(&mut phase_settings, |setting| {
        max_value = max(max_value, do_part1(initial_memory, setting));
    });

    return max_value;
}

pub fn part2(initial_memory: &str) -> i32 {
    let mut phase_settings = vec![5, 6, 7, 8, 9];
    let mut max_value = 0;
    heap_recursive(&mut phase_settings, |setting| {
        max_value = max(max_value, do_part2(initial_memory, setting));
    });

    return max_value;
}

fn do_part1(initial_memory: &str, phase_settings: &[i32]) -> i32 {
    let mut last_output = 0;

    for setting in phase_settings {
        let mut inputs = VecDeque::new();
        inputs.push_back(*setting);
        inputs.push_back(last_output);

        let mut outputs = VecDeque::new();
        let mut computer = IntcodeProcess::new_from_string(initial_memory);
        computer.run(&mut inputs, &mut outputs);

        last_output = outputs.pop_front().unwrap();
    }
    return last_output;
}

fn do_part2(initial_memory: &str, phase_settings: &[i32]) -> i32 {
    let mut output_a = VecDeque::new();
    let mut output_b = VecDeque::new();
    let mut output_c = VecDeque::new();
    let mut output_d = VecDeque::new();
    let mut output_e = VecDeque::new();

    // Apply the phase settings to the inputs, because outputs are connected to
    // inputs this is slightly confusing.
    output_a.push_back(phase_settings[1]);
    output_b.push_back(phase_settings[2]);
    output_c.push_back(phase_settings[3]);
    output_d.push_back(phase_settings[4]);
    output_e.push_back(phase_settings[0]);

    // Apply zero to the input of the first process.
    output_e.push_back(0);

    let mut process_a = IntcodeProcess::new_from_string(initial_memory);
    let mut process_b = IntcodeProcess::new_from_string(initial_memory);
    let mut process_c = IntcodeProcess::new_from_string(initial_memory);
    let mut process_d = IntcodeProcess::new_from_string(initial_memory);
    let mut process_e = IntcodeProcess::new_from_string(initial_memory);

    loop {
        let state_a = process_a.run(&mut output_e, &mut output_a);
        let state_b = process_b.run(&mut output_a, &mut output_b);
        let state_c = process_c.run(&mut output_b, &mut output_c);
        let state_d = process_d.run(&mut output_c, &mut output_d);
        let state_e = process_e.run(&mut output_d, &mut output_e);

        if state_a == IntcodeProcessState::Halted
            && state_b == IntcodeProcessState::Halted
            && state_c == IntcodeProcessState::Halted
            && state_d == IntcodeProcessState::Halted
            && state_e == IntcodeProcessState::Halted
        {
            break;
        }
    }

    return output_e.pop_back().unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let cases = vec![
            (
                "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
                vec![4, 3, 2, 1, 0],
                43210,
            ),
            (
                "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
                vec![0, 1, 2, 3, 4],
                54321,
            ),
            (
                "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
                vec![1, 0, 4, 3, 2],
                65210,
            ),
        ];

        for case in cases {
            assert_eq!(do_part1(case.0, &case.1), case.2);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![
            (
                "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5",
                vec![9, 8, 7, 6, 5],
                139629729,
            ),
            (
                "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
                vec![9, 7, 8, 5, 6],
                18216,
            ),
        ];

        for case in cases {
            assert_eq!(do_part2(case.0, &case.1), case.2);
        }
    }
}
