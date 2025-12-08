use aoc2025::days::day08::{part1, part2};
use std::env;
use std::fs;
use std::process;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 3 {
        println!("Missing required arguments");
        process::exit(1);
    }

    let contents = fs::read_to_string(&args[1]).expect("Something went wrong");
    let count: usize = args[2].parse().expect("Invalid number for count");
    println!("Part 1: {}", part1(&contents, count));
    println!("Part 2: {}", part2(&contents));
}
