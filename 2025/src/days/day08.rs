use core::panic;
use std::collections::{HashMap, HashSet};

#[derive(Eq, Hash, PartialEq)]
struct Point {
    x: i64,
    y: i64,
    z: i64,
}

impl Point {
    fn distance(&self, other: &Point) -> f64 {
        (((self.x - other.x).pow(2) + (self.y - other.y).pow(2) + (self.z - other.z).pow(2)) as f64)
            .sqrt()
    }
}

pub fn part1(contents: &str, count: usize) -> i64 {
    let junction_boxes = parse_input(contents);
    let mut distances = Vec::new();
    for i in 0..junction_boxes.len() {
        for j in i + 1..junction_boxes.len() {
            let distance = junction_boxes[i].distance(&junction_boxes[j]);
            distances.push((distance, i, j));
        }
    }
    distances.sort_by(|a, b| a.0.partial_cmp(&b.0).unwrap());
    let mut graph = HashMap::new();
    for distance in distances.iter().take(count) {
        graph
            .entry(distance.1)
            .or_insert(Vec::new())
            .push(distance.2);
        graph
            .entry(distance.2)
            .or_insert(Vec::new())
            .push(distance.1);
    }

    let mut circuits = Vec::new();
    let mut visited = vec![false; junction_boxes.len()];
    for i in 0..junction_boxes.len() {
        if !visited[i] {
            let mut count = 0;
            let mut stack = vec![i];
            while let Some(node) = stack.pop() {
                if !visited[node] {
                    count += 1;
                    visited[node] = true;
                    if let Some(neighbors) = graph.get(&node) {
                        for &neighbor in neighbors.iter() {
                            if !visited[neighbor] {
                                stack.push(neighbor);
                            }
                        }
                    }
                }
            }
            circuits.push(count);
        }
    }
    circuits.sort_by(|a, b| b.cmp(a));
    circuits.iter().take(3).product()
}

pub fn part2(contents: &str) -> i64 {
    let junction_boxes = parse_input(contents);
    let mut distances = Vec::new();
    for i in 0..junction_boxes.len() {
        for j in i + 1..junction_boxes.len() {
            let distance = junction_boxes[i].distance(&junction_boxes[j]);
            distances.push((distance, i, j));
        }
    }
    distances.sort_by(|a, b| a.0.partial_cmp(&b.0).unwrap());
    let mut graph = HashMap::new();
    for distance in distances {
        graph
            .entry(distance.1)
            .or_insert(Vec::new())
            .push(distance.2);
        graph
            .entry(distance.2)
            .or_insert(Vec::new())
            .push(distance.1);

        let junction_count = count_junction_boxes(&graph, distance.1);
        if junction_count == junction_boxes.len() {
            return junction_boxes[distance.1].x * junction_boxes[distance.2].x;
        }
    }
    panic!("No solution found");
}

fn parse_input(contents: &str) -> Vec<Point> {
    contents
        .lines()
        .map(|line| {
            let coords: Vec<i64> = line.split(',').map(|c| c.parse().unwrap()).collect();
            Point {
                x: coords[0],
                y: coords[1],
                z: coords[2],
            }
        })
        .collect()
}

fn count_junction_boxes(graph: &HashMap<usize, Vec<usize>>, start: usize) -> usize {
    let mut visited = HashSet::new();
    let mut count = 0;
    let mut stack = vec![start];
    while let Some(node) = stack.pop() {
        if !visited.contains(&node) {
            count += 1;
            visited.insert(node);
            if let Some(neighbors) = graph.get(&node) {
                for &neighbor in neighbors.iter() {
                    if !visited.contains(&neighbor) {
                        stack.push(neighbor);
                    }
                }
            }
        }
    }
    count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            "162,817,812",
            "57,618,57",
            "906,360,560",
            "592,479,940",
            "352,342,300",
            "466,668,158",
            "542,29,236",
            "431,825,988",
            "739,650,466",
            "52,470,668",
            "216,146,977",
            "819,987,18",
            "117,168,530",
            "805,96,715",
            "346,949,466",
            "970,615,88",
            "941,993,340",
            "862,61,35",
            "984,92,344",
            "425,690,689",
        ];

        assert_eq!(part1(&input.join("\n"), 10), 40);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            "162,817,812",
            "57,618,57",
            "906,360,560",
            "592,479,940",
            "352,342,300",
            "466,668,158",
            "542,29,236",
            "431,825,988",
            "739,650,466",
            "52,470,668",
            "216,146,977",
            "819,987,18",
            "117,168,530",
            "805,96,715",
            "346,949,466",
            "970,615,88",
            "941,993,340",
            "862,61,35",
            "984,92,344",
            "425,690,689",
        ];

        assert_eq!(part2(&input.join("\n")), 25272);
    }
}
