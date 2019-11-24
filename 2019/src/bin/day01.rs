use aoc2019::day01;
use std::env;
use std::fs;
use std::process;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("Missing required arguments");
        process::exit(1);
    }

    let contents = fs::read_to_string(&args[1])
        .expect("Something went wrong");

    let lines: Vec<&str> = contents.lines().collect();
    println!("Part 1: {}", day01::part1(&lines));
    println!("Part 2: {}", day01::part2(&lines));
}
