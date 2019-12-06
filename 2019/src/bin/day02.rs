use aoc2019::day02;
use std::env;
use std::fs;
use std::process;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("Missing required arguments");
        process::exit(1);
    }

    let contents = fs::read_to_string(&args[1]).expect("Something went wrong");

    let initial_memory = contents
        .trim()
        .split(',')
        .map(|i| i.parse::<usize>().unwrap())
        .collect();

    println!("Part 1: {}", day02::part1(&initial_memory));
    println!("Part 2: {}", day02::part2(&initial_memory));
}
