use std::collections::{HashMap, VecDeque};

pub fn part1(contents: &str) -> i64 {
    let graph = parse_input(contents);
    let mut count = 0;
    let mut queue = VecDeque::new();
    queue.push_back("you");
    while !queue.is_empty() {
        let current = queue.pop_front().unwrap();
        if current == "out" {
            count += 1;
            continue;
        }
        if let Some(neighbors) = graph.get(current) {
            for &neighbor in neighbors {
                queue.push_back(neighbor);
            }
        }
    }
    count
}

pub fn part2(contents: &str) -> i64 {
    let graph = parse_input(contents);

    // Determine which path exists between "dac" and "fft"
    let paths_dac_fft = count_paths(&graph, "dac", "fft", &mut HashMap::new());
    let paths_fft_dac = count_paths(&graph, "fft", "dac", &mut HashMap::new());

    if paths_dac_fft > 0 {
        // There is a path from "dac" to "fft"
        let paths_srv_dac = count_paths(&graph, "svr", "dac", &mut HashMap::new());
        let paths_fft_out = count_paths(&graph, "fft", "out", &mut HashMap::new());
        paths_dac_fft * paths_fft_out * paths_srv_dac
    } else {
        // There is a path from "fft" to "dac"
        let paths_srv_fft = count_paths(&graph, "svr", "fft", &mut HashMap::new());
        let paths_dac_out = count_paths(&graph, "dac", "out", &mut HashMap::new());
        paths_fft_dac * paths_dac_out * paths_srv_fft
    }
}

fn parse_input(contents: &str) -> HashMap<&str, Vec<&str>> {
    let mut graph = HashMap::new();
    for l in contents.lines() {
        let parts: Vec<&str> = l.split(": ").collect();
        let device = parts[0];
        let outputs: Vec<&str> = parts[1].split(' ').collect();
        graph.insert(device, outputs);
    }
    graph
}

fn count_paths(
    graph: &HashMap<&str, Vec<&str>>,
    start: &str,
    end: &str,
    memo: &mut HashMap<String, i64>,
) -> i64 {
    if let Some(&cached) = memo.get(start) {
        return cached;
    }

    if start == end {
        return 1;
    }

    let mut total_paths = 0;
    if let Some(neighbors) = graph.get(start) {
        for &neighbor in neighbors {
            total_paths += count_paths(graph, neighbor, end, memo);
        }
    }

    memo.insert(start.to_string(), total_paths);
    total_paths
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            "aaa: you hhh",
            "you: bbb ccc",
            "bbb: ddd eee",
            "ccc: ddd eee fff",
            "ddd: ggg",
            "eee: out",
            "fff: out",
            "ggg: out",
            "hhh: ccc fff iii",
            "iii: out",
        ];

        assert_eq!(part1(&input.join("\n")), 5);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            "svr: aaa bbb",
            "aaa: fft",
            "fft: ccc",
            "bbb: tty",
            "tty: ccc",
            "ccc: ddd eee",
            "ddd: hub",
            "hub: fff",
            "eee: dac",
            "dac: fff",
            "fff: ggg hhh",
            "ggg: out",
            "hhh: out",
        ];

        assert_eq!(part2(&input.join("\n")), 2);
    }
}
