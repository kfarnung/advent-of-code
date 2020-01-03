use aoc2019::days::day16::{part1, part2};
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
    println!("Part 1: {:0>8}", part1(&contents));
    println!("Part 2: {:0>8}", part2(&contents));
}
