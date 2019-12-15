use crate::shared::num::lcm;
use crate::shared::point::Point3D;
use regex::Regex;
use std::cmp::max;
use std::cmp::min;

pub fn part1(moons: &str) -> i32 {
    return do_part1(moons, 1000);
}

pub fn part2(moons: &str) -> i64 {
    let mut planets = parse_planets(moons);
    let initial_state: Vec<Planet> = planets.iter().map(|x| x.clone()).collect();
    let mut step = 0;
    let mut cycle_counts = (0, 0, 0);

    while cycle_counts.0 == 0 || cycle_counts.1 == 0 || cycle_counts.2 == 0 {
        do_step(&mut planets);
        step += 1;

        if cycle_counts.0 == 0
            && detect_cycle(&initial_state, &planets, |x| x.position.x, |x| x.velocity.x)
        {
            cycle_counts.0 = step;
        }

        if cycle_counts.1 == 0
            && detect_cycle(&initial_state, &planets, |x| x.position.y, |x| x.velocity.y)
        {
            cycle_counts.1 = step;
        }

        if cycle_counts.2 == 0
            && detect_cycle(&initial_state, &planets, |x| x.position.z, |x| x.velocity.z)
        {
            cycle_counts.2 = step;
        }
    }

    return lcm(cycle_counts.0, lcm(cycle_counts.1, cycle_counts.2));
}

fn do_part1(moons: &str, steps: usize) -> i32 {
    let mut planets = parse_planets(moons);

    for _count in 0..steps {
        do_step(&mut planets);
    }

    return planets.iter().fold(0, |acc, x| acc + x.potential_energy());
}

fn do_step(planets: &mut Vec<Planet>) {
    let positions: Vec<Point3D> = planets.iter().map(|x| x.position.clone()).collect();

    for position in positions {
        for planet in planets.iter_mut() {
            update_velocity(planet, &position);
        }
    }

    for planet in planets.iter_mut() {
        let velocity = planet.velocity.clone();
        planet.position += velocity;
    }
}

#[derive(Clone, Debug, PartialEq)]
struct Planet {
    position: Point3D,
    velocity: Point3D,
}

impl Planet {
    fn potential_energy(&self) -> i32 {
        let origin = Point3D::new(0, 0, 0);

        return self.position.manhattan_distance(&origin)
            * self.velocity.manhattan_distance(&origin);
    }
}

fn parse_planets(moons: &str) -> Vec<Planet> {
    return moons.lines().map(|x| parse_planet(x)).collect();
}

fn parse_planet(moon: &str) -> Planet {
    let re = Regex::new(r"<x=(-?\d+), y=(-?\d+), z=(-?\d+)>").unwrap();
    let caps = re.captures(moon).unwrap();
    return Planet {
        position: Point3D::new(
            caps[1].parse::<i32>().unwrap(),
            caps[2].parse::<i32>().unwrap(),
            caps[3].parse::<i32>().unwrap(),
        ),
        velocity: Point3D::new(0, 0, 0),
    };
}

fn update_velocity(first: &mut Planet, second: &Point3D) {
    let x = velocity_delta(first.position.x, second.x);
    let y = velocity_delta(first.position.y, second.y);
    let z = velocity_delta(first.position.z, second.z);

    first.velocity += Point3D::new(x, y, z);
}

fn velocity_delta(first: i32, second: i32) -> i32 {
    return max(-1, min(1, second - first));
}

fn detect_cycle(
    initial_state: &Vec<Planet>,
    current_state: &Vec<Planet>,
    position_key: fn(&Planet) -> i32,
    velocity_key: fn(&Planet) -> i32,
) -> bool {
    let initial = initial_state
        .iter()
        .map(|x| (position_key(x), velocity_key(x)));
    let current = current_state
        .iter()
        .map(|x| (position_key(x), velocity_key(x)));
    return initial.eq(current);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse_planet() {
        assert_eq!(
            parse_planet("<x=2, y=-10, z=-7>"),
            Planet {
                position: Point3D::new(2, -10, -7),
                velocity: Point3D::new(0, 0, 0)
            }
        );
    }

    #[test]
    fn test_velocity_delta() {
        assert_eq!(velocity_delta(3, 5), 1);
        assert_eq!(velocity_delta(5, 3), -1);
        assert_eq!(velocity_delta(3, 3), 0);
    }

    #[test]
    fn test_detect_cycle() {
        let initial_state = vec![
            Planet {
                position: Point3D::new(1, 2, 3),
                velocity: Point3D::new(4, 5, 6),
            },
            Planet {
                position: Point3D::new(7, 8, 9),
                velocity: Point3D::new(10, 11, 12),
            },
        ];

        let current_state = vec![
            Planet {
                position: Point3D::new(1, 2, 3),
                velocity: Point3D::new(4, 5, 6),
            },
            Planet {
                position: Point3D::new(7, 8, 9),
                velocity: Point3D::new(10, 11, 12),
            },
        ];

        let current_state2 = vec![
            Planet {
                position: Point3D::new(1, 2, 3),
                velocity: Point3D::new(4, 5, 6),
            },
            Planet {
                position: Point3D::new(7, 13, 9),
                velocity: Point3D::new(10, 11, 12),
            },
        ];

        assert_eq!(
            detect_cycle(
                &initial_state,
                &current_state,
                |x| x.position.x,
                |x| x.velocity.x
            ),
            true
        );
        assert_eq!(
            detect_cycle(
                &initial_state,
                &current_state2,
                |x| x.position.y,
                |x| x.velocity.y
            ),
            false
        );
    }

    #[test]
    fn test_part1() {
        let cases = vec![
            (
                vec![
                    "<x=-1, y=0, z=2>",
                    "<x=2, y=-10, z=-7>",
                    "<x=4, y=-8, z=8>",
                    "<x=3, y=5, z=-1>",
                ],
                10,
                179,
            ),
            (
                vec![
                    "<x=-8, y=-10, z=0>",
                    "<x=5, y=5, z=10>",
                    "<x=2, y=-7, z=3>",
                    "<x=9, y=-8, z=-3>",
                ],
                100,
                1940,
            ),
        ];

        for case in cases {
            assert_eq!(do_part1(&case.0.join("\n"), case.1), case.2);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![
            (
                vec![
                    "<x=-1, y=0, z=2>",
                    "<x=2, y=-10, z=-7>",
                    "<x=4, y=-8, z=8>",
                    "<x=3, y=5, z=-1>",
                ],
                2772,
            ),
            (
                vec![
                    "<x=-8, y=-10, z=0>",
                    "<x=5, y=5, z=10>",
                    "<x=2, y=-7, z=3>",
                    "<x=9, y=-8, z=-3>",
                ],
                4686774924,
            ),
        ];

        for case in cases {
            assert_eq!(part2(&case.0.join("\n")), case.1);
        }
    }
}
