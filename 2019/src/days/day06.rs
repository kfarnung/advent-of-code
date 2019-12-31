use crate::shared::tree::NamedTree;
use std::collections::HashSet;
use std::collections::VecDeque;

pub fn part1(connections: &str) -> i32 {
    let mut tree = NamedTree::new();

    for line in connections.lines() {
        let parts: Vec<&str> = line.split(")").collect();
        tree.add_child(parts[0], parts[1]);
    }

    let mut orbits = 0;
    let mut queue = VecDeque::new();
    queue.push_back(("COM", 0));

    while let Some((name, depth)) = queue.pop_front() {
        orbits += depth;

        for child in tree.get_children(name) {
            queue.push_back((child, depth + 1));
        }
    }

    return orbits;
}

pub fn part2(connections: &str) -> i32 {
    let mut tree = NamedTree::new();

    for line in connections.lines() {
        let parts: Vec<&str> = line.split(")").collect();
        tree.add_child(parts[0], parts[1]);
    }

    let mut visited = HashSet::new();
    let mut queue = VecDeque::new();
    queue.push_back(("YOU", 0));

    while let Some((name, distance)) = queue.pop_front() {
        visited.insert(name.to_string());

        if let Some(parent) = tree.get_parent(name) {
            if parent == "SAN" {
                return distance - 1;
            }

            if !visited.contains(parent) {
                queue.push_back((parent, distance + 1));
            }
        }

        for child in tree.get_children(name) {
            if child == "SAN" {
                return distance - 1;
            }

            if visited.contains(child) {
                continue;
            }

            queue.push_back((child, distance + 1));
        }
    }

    panic!("We couldn't find Santa!");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let cases = vec![(
            vec![
                "COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L",
            ],
            42,
        )];

        for case in cases {
            assert_eq!(part1(&case.0.join("\n")), case.1);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![(
            vec![
                "COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L",
                "K)YOU", "I)SAN",
            ],
            4,
        )];

        for case in cases {
            assert_eq!(part2(&case.0.join("\n")), case.1);
        }
    }
}
