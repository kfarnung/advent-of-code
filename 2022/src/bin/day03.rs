use aoc2022::days::day03::{part1, part2};
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
    println!("Part 1: {}", part1(&contents));
    println!("Part 2: {}", part2(&contents));
}
