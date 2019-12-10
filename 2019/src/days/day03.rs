use crate::shared::point::Point2D;
use std::collections::HashMap;

#[derive(Debug, PartialEq)]
struct PathStep {
    direction: String,
    distance: i32,
}

pub fn part1(lines: &Vec<&str>) -> i32 {
    let wire_a = follow_path(lines[0]);
    let wire_b = follow_path(lines[1]);

    let intersections = find_intersections(wire_a, wire_b);
    let min_distance = intersections
        .iter()
        .min_by_key(|x| x.0.manhattan_distance(&Point2D::new(0, 0)))
        .unwrap()
        .0;

    return min_distance.manhattan_distance(&Point2D::new(0, 0));
}

pub fn part2(lines: &Vec<&str>) -> i32 {
    let wire_a = follow_path(lines[0]);
    let wire_b = follow_path(lines[1]);

    let intersections = find_intersections(wire_a, wire_b);
    let min_distance = intersections.iter().min_by_key(|x| x.1).unwrap().1;

    return min_distance.clone();
}

fn find_intersections(
    wire_a: HashMap<Point2D, i32>,
    wire_b: HashMap<Point2D, i32>,
) -> HashMap<Point2D, i32> {
    let mut intersections: HashMap<Point2D, i32> = HashMap::new();

    for point_a in wire_a {
        match wire_b.get(&point_a.0) {
            Some(x) => intersections.insert(point_a.0, point_a.1 + x),
            None => None,
        };
    }

    return intersections;
}

fn follow_path(line: &str) -> HashMap<Point2D, i32> {
    let mut visited: HashMap<Point2D, i32> = HashMap::new();
    let mut current = Point2D::new(0, 0);
    let mut total_distance = 0;

    for step in get_steps(line) {
        let delta = get_delta(&step.direction);

        for _ in 0..step.distance {
            total_distance += 1;
            current += delta.clone();
            visited.insert(current.clone(), total_distance);
        }
    }

    return visited;
}

fn get_steps(line: &str) -> Vec<PathStep> {
    return line.split(",").map(|s| get_step(s)).collect();
}

fn get_step(step: &str) -> PathStep {
    let (direction, distance) = step.split_at(1);

    return PathStep {
        direction: direction.to_string(),
        distance: distance.parse().unwrap(),
    };
}

fn get_delta(direction: &str) -> Point2D {
    match direction {
        "U" => Point2D::new(0, 1),
        "D" => Point2D::new(0, -1),
        "L" => Point2D::new(-1, 0),
        "R" => Point2D::new(1, 0),
        _ => panic!("Unexpected direction!"),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_steps() {
        assert_eq!(
            get_steps("R75,D30,U83,L12"),
            vec![
                PathStep {
                    direction: "R".to_string(),
                    distance: 75,
                },
                PathStep {
                    direction: "D".to_string(),
                    distance: 30,
                },
                PathStep {
                    direction: "U".to_string(),
                    distance: 83,
                },
                PathStep {
                    direction: "L".to_string(),
                    distance: 12,
                },
            ]
        )
    }

    #[test]
    fn test_part1() {
        let cases = vec![
            (
                vec![
                    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
                    "U62,R66,U55,R34,D71,R55,D58,R83",
                ],
                159,
            ),
            (
                vec![
                    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
                    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
                ],
                135,
            ),
        ];

        for case in cases {
            assert_eq!(part1(&case.0), case.1);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![
            (
                vec![
                    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
                    "U62,R66,U55,R34,D71,R55,D58,R83",
                ],
                610,
            ),
            (
                vec![
                    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
                    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
                ],
                410,
            ),
        ];

        for case in cases {
            assert_eq!(part2(&case.0), case.1);
        }
    }
}
