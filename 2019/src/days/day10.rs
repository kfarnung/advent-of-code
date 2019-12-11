use crate::shared::num::Fraction;
use crate::shared::point::Point2D;
use ordered_float::OrderedFloat;
use std::collections::HashMap;
use std::collections::HashSet;

pub fn part1(map: &str) -> usize {
    let asteroids = parse_map(map);
    return find_location(&asteroids).1;
}

pub fn part2(map: &str) -> i32 {
    let asteroids = parse_map(map);
    let best_location = find_location(&asteroids).0;

    let base_angle = Fraction::new(-1, 0);
    let mut hits_map = find_hits(&asteroids, &best_location);
    let mut angles = get_sorted_angles(&hits_map);
    angles.sort_by_cached_key(|x| OrderedFloat(x.angle_relative_cw(&base_angle)));

    let mut count = 0;

    loop {
        for angle in &angles {
            let hits = hits_map.get_mut(&angle).unwrap();
            hits.sort_by_cached_key(|x| x.manhattan_distance(&best_location));

            let point = hits.remove(0);
            count += 1;

            if count == 200 {
                return point.x * 100 + point.y;
            }
        }
    }
}

fn get_sorted_angles(hits_map: &HashMap<Fraction, Vec<Point2D>>) -> Vec<Fraction> {
    let keys: Vec<&Fraction> = hits_map.keys().collect();
    return keys.iter().cloned().cloned().collect();
}

fn parse_map(map: &str) -> HashSet<Point2D> {
    let mut asteroids = HashSet::new();

    for (y, line) in map.lines().enumerate() {
        for (x, value) in line.chars().enumerate() {
            if value == '#' {
                asteroids.insert(Point2D::new(x as i32, y as i32));
            }
        }
    }

    return asteroids;
}

fn find_location(asteroids: &HashSet<Point2D>) -> (Point2D, usize) {
    let mut positions = Vec::new();

    for point in asteroids {
        let hits_map = find_hits(&asteroids, point);
        positions.push((point.clone(), hits_map.len()));
    }

    return positions.iter().max_by_key(|x| x.1).unwrap().clone();
}

fn find_hits(asteroids: &HashSet<Point2D>, point: &Point2D) -> HashMap<Fraction, Vec<Point2D>> {
    let mut hits_map: HashMap<Fraction, Vec<Point2D>> = HashMap::new();

    for other in asteroids {
        if point == other {
            continue;
        }

        let slope = point.slope(&other);
        if let Some(x) = hits_map.get_mut(&slope) {
            x.push(other.clone());
        } else {
            hits_map.insert(slope, vec![other.clone()]);
        }
    }

    return hits_map;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let cases = vec![
            (vec![".#..#", ".....", "#####", "....#", "...##"], 8),
            (
                vec![
                    "......#.#.",
                    "#..#.#....",
                    "..#######.",
                    ".#.#.###..",
                    ".#..#.....",
                    "..#....#.#",
                    "#..#....#.",
                    ".##.#..###",
                    "##...#..#.",
                    ".#....####",
                ],
                33,
            ),
        ];

        for case in cases {
            assert_eq!(part1(&case.0.join("\n")), case.1);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![(
            vec![
                ".#..##.###...#######",
                "##.############..##.",
                ".#.######.########.#",
                ".###.#######.####.#.",
                "#####.##.#.##.###.##",
                "..#####..#.#########",
                "####################",
                "#.####....###.#.#.##",
                "##.#################",
                "#####.##.###..####..",
                "..######..##.#######",
                "####.##.####...##..#",
                ".#####..#.######.###",
                "##...#.##########...",
                "#.##########.#######",
                ".####.#.###.###.#.##",
                "....##.##.###..#####",
                ".#.#.###########.###",
                "#.#.#.#####.####.###",
                "###.##.####.##.#..##",
            ],
            802,
        )];

        for case in cases {
            assert_eq!(part2(&case.0.join("\n")), case.1);
        }
    }
}
