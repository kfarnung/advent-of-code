pub struct IntcodeComputer {
    memory: Vec<usize>,
}

impl IntcodeComputer {
    pub fn new(initial_memory: &Vec<usize>) -> Self {
        return IntcodeComputer {
            memory: initial_memory.clone(),
        };
    }

    pub fn run(&mut self, noun: Option<usize>, verb: Option<usize>) {
        let mut ip = 0;
    
        match noun {
            Some(x) => self.memory[1] = x,
            None => (),
        };
    
        match verb {
            Some(x) => self.memory[2] = x,
            None => (),
        };
    
        loop {
            match self.memory[ip] {
                1 => ip = do_add(&mut self.memory, ip),
                2 => ip = do_multiply(&mut self.memory, ip),
                99 => break,
                _ => panic!("Unexpected opcode!"),
            };
        }
    }

    pub fn get_value(&self, address: usize) -> usize {
        return self.memory[address];
    }
}

fn do_add(memory: &mut Vec<usize>, ip: usize) -> usize {
    let a = memory[ip + 1];
    let b = memory[ip + 2];
    let r = memory[ip + 3];

    memory[r] = memory[a] + memory[b];
    return ip + 4;
}

fn do_multiply(memory: &mut Vec<usize>, ip: usize) -> usize {
    let a = memory[ip + 1];
    let b = memory[ip + 2];
    let r = memory[ip + 3];

    memory[r] = memory[a] * memory[b];
    return ip + 4;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_run() {
        let cases = vec![
            (
                vec![1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50],
                vec![3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50],
            ),
            (vec![1, 0, 0, 0, 99], vec![2, 0, 0, 0, 99]),
            (vec![2, 3, 0, 3, 99], vec![2, 3, 0, 6, 99]),
            (vec![2, 4, 4, 5, 99, 0], vec![2, 4, 4, 5, 99, 9801]),
            (
                vec![1, 1, 1, 4, 99, 5, 6, 0, 99],
                vec![30, 1, 1, 4, 2, 5, 6, 0, 99],
            ),
        ];

        for case in cases {
            let mut computer = IntcodeComputer::new(&case.0);
            computer.run(None, None);
            assert_eq!(computer.memory, case.1);
        }
    }
}
