use aoc2019::days::day04;
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

    println!("Part 1: {}", day04::part1(&contents));
    println!("Part 2: {}", day04::part2(&contents));
}
