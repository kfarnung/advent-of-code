use crate::shared::point::Point2D;
use core::cmp::Ordering;
use std::collections::BinaryHeap;
use std::collections::HashMap;
use std::collections::HashSet;
use std::collections::VecDeque;

const GRID_FLOOR: char = '.';
const GRID_START: char = '@';
const GRID_WALL: char = '#';
const KEY_MIN: char = 'a';
const KEY_MAX: char = 'z';
const DOOR_MIN: char = 'A';
const DOOR_MAX: char = 'Z';
const SPLIT_START: [&str; 3] = ["@#@", "###", "@#@"];

fn is_key(value: &char) -> bool {
    *value >= KEY_MIN && *value <= KEY_MAX
}

fn is_door(value: &char) -> bool {
    *value >= DOOR_MIN && *value <= DOOR_MAX
}

type ReachableKey = (Point2D<i32>, char, usize, TunnelKeys);
type ReachableKeyMap = HashMap<Point2D<i32>, Vec<ReachableKey>>;

struct Tunnels {
    all_keys: TunnelKeys,
    cache: ReachableKeyMap,
    grid: HashMap<Point2D<i32>, char>,
    start: Vec<Point2D<i32>>,
}

impl Tunnels {
    fn parse(content: &str) -> Self {
        let mut grid = HashMap::new();
        let mut start = Vec::new();
        let mut all_keys = TunnelKeys::new();

        for (y, line) in content.lines().enumerate() {
            for (x, value) in line.chars().enumerate() {
                let point = Point2D::new(x as i32, y as i32);
                if value == GRID_START {
                    start.push(point);

                    // Treat is as floor instead.
                    grid.insert(point, GRID_FLOOR);
                } else {
                    if is_key(&value) {
                        all_keys.add_key(&value);
                    }

                    grid.insert(point, value);
                }
            }
        }

        Self {
            all_keys,
            cache: HashMap::new(),
            grid,
            start,
        }
    }

    fn replace_start(&mut self) {
        if self.start.len() != 1 {
            panic!("Can't replace multiple start points")
        }

        let base = self.start[0] + Point2D::new(-1, -1);
        self.start = Vec::new();

        for (y, line) in SPLIT_START.iter().enumerate() {
            for (x, ch) in line.chars().enumerate() {
                let point = base + Point2D::new(x as i32, y as i32);
                if ch == GRID_START {
                    self.start.push(point);
                    self.grid.insert(point, GRID_FLOOR);
                } else {
                    self.grid.insert(point, ch);
                }
            }
        }
    }

    fn find_reachable_keys(
        &mut self,
        start: &Point2D<i32>,
    ) -> Vec<(Point2D<i32>, char, usize, TunnelKeys)> {
        if let Some(cache_result) = self.cache.get(start) {
            return cache_result.clone();
        }

        let mut queue = VecDeque::new();
        let mut visited = HashSet::new();
        let mut reachable_keys = Vec::new();
        queue.push_back((*start, 0, TunnelKeys::new()));

        while !queue.is_empty() {
            let (current_position, current_distance, current_keys) = queue.pop_front().unwrap();
            visited.insert(current_position);

            let value = self.grid.get(&current_position).unwrap_or(&GRID_WALL);
            if is_key(value) && current_position != *start {
                reachable_keys.push((current_position, *value, current_distance, current_keys));

                // No need to search further down this path.
                continue;
            }

            let mut required_keys = current_keys;
            if is_door(value) {
                // Add a required key to the set.
                required_keys.add_key(&value.to_ascii_lowercase());
            }

            let neighbors = vec![
                current_position + Point2D::new(0, -1),
                current_position + Point2D::new(0, 1),
                current_position + Point2D::new(-1, 0),
                current_position + Point2D::new(1, 0),
            ];

            for neighbor in neighbors {
                if visited.contains(&neighbor) {
                    continue;
                }

                let value = self.grid.get(&neighbor).unwrap_or(&GRID_WALL);
                if value == &GRID_WALL {
                    continue;
                }

                queue.push_back((neighbor, current_distance + 1, required_keys));
            }
        }

        self.cache.insert(*start, reachable_keys.clone());
        reachable_keys
    }
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
struct TunnelKeys {
    bit_field: u32,
}

impl TunnelKeys {
    fn new() -> Self {
        TunnelKeys { bit_field: 0 }
    }

    fn add_key(&mut self, other: &char) -> bool {
        if self.has_key(other) {
            return false;
        }

        match TunnelKeys::key_to_mask(other) {
            Some(value) => {
                self.bit_field |= value;
                true
            }
            None => false,
        }
    }

    fn has_key(&self, other: &char) -> bool {
        match TunnelKeys::key_to_mask(other) {
            Some(value) => (self.bit_field & value) == value,
            None => false,
        }
    }

    fn intersects(&self, other: &TunnelKeys) -> bool {
        (self.bit_field & other.bit_field) == other.bit_field
    }

    fn key_to_mask(other: &char) -> Option<u32> {
        if !is_key(other) {
            return None;
        }

        let position = (*other as u8) - (KEY_MIN as u8);
        Some(1_u32 << position)
    }
}

#[derive(Clone, Eq, PartialEq)]
struct State {
    distance: usize,
    keys: TunnelKeys,
    positions: Vec<Point2D<i32>>,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        other.distance.cmp(&self.distance)
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn search_tunnels(tunnels: &mut Tunnels) -> usize {
    let mut seen_states = HashMap::new();
    let mut priority_queue = BinaryHeap::new();

    priority_queue.push(State {
        distance: 0,
        keys: TunnelKeys::new(),
        positions: tunnels.start.clone(),
    });

    seen_states.insert((tunnels.start.clone(), TunnelKeys::new()), 0);

    while !priority_queue.is_empty() {
        let current = priority_queue.pop().unwrap();

        // We are always getting the next shortest distance, if it has all the
        // keys collected, then it's the shortest path to collect all keys.
        if current.keys == tunnels.all_keys {
            return current.distance;
        }

        for (i, current_position) in current.positions.iter().enumerate() {
            let mut positions = current.positions.clone();

            let found_keys = tunnels.find_reachable_keys(current_position);
            for (position, key, distance, required_keys) in found_keys {
                if !current.keys.intersects(&required_keys) {
                    // Location is inaccessible currently.
                    continue;
                }

                positions[i] = position;
                let mut keys = current.keys;
                keys.add_key(&key);

                let distance = current.distance + distance;
                let seen_distance = seen_states
                    .entry((positions.clone(), keys))
                    .or_insert(usize::MAX);
                if *seen_distance > distance {
                    *seen_distance = distance;
                    priority_queue.push(State {
                        distance,
                        keys,
                        positions: positions.clone(),
                    });
                }
            }
        }
    }

    0
}

pub fn part1(contents: &str) -> usize {
    let mut tunnels = Tunnels::parse(contents);
    search_tunnels(&mut tunnels)
}

pub fn part2(contents: &str) -> usize {
    let mut tunnels = Tunnels::parse(contents);
    tunnels.replace_start();
    search_tunnels(&mut tunnels)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse() {
        let case = ["#########", "#b.A.@.a#", "#########"];
        let tunnels = Tunnels::parse(&case.join("\n"));
        assert_eq!(tunnels.grid.len(), 27);
        assert_eq!(tunnels.start[0], Point2D::new(5, 1));
    }

    #[test]
    fn test_replace_start() {
        let input = [
            "#######", "#a.#Cd#", "##...##", "##.@.##", "##...##", "#cB#Ab#", "#######",
        ];
        let output = [
            "#######", "#a.#Cd#", "##@#@##", "#######", "##@#@##", "#cB#Ab#", "#######",
        ];

        let mut tunnels = Tunnels::parse(&input.join("\n"));
        tunnels.replace_start();

        let output_tunnels = Tunnels::parse(&output.join("\n"));
        assert_eq!(tunnels.grid, output_tunnels.grid);
        assert_eq!(tunnels.start, output_tunnels.start);
    }

    #[test]
    fn test_part1() {
        let cases = vec![
            (vec!["#########", "#b.A.@.a#", "#########"], 8),
            (
                vec![
                    "########################",
                    "#f.D.E.e.C.b.A.@.a.B.c.#",
                    "######################.#",
                    "#d.....................#",
                    "########################",
                ],
                86,
            ),
            (
                vec![
                    "########################",
                    "#...............b.C.D.f#",
                    "#.######################",
                    "#.....@.a.B.c.d.A.e.F.g#",
                    "########################",
                ],
                132,
            ),
            (
                vec![
                    "#################",
                    "#i.G..c...e..H.p#",
                    "########.########",
                    "#j.A..b...f..D.o#",
                    "########@########",
                    "#k.E..a...g..B.n#",
                    "########.########",
                    "#l.F..d...h..C.m#",
                    "#################",
                ],
                136,
            ),
            (
                vec![
                    "########################",
                    "#@..............ac.GI.b#",
                    "###d#e#f################",
                    "###A#B#C################",
                    "###g#h#i################",
                    "########################",
                ],
                81,
            ),
        ];

        for case in cases {
            assert_eq!(part1(&case.0.join("\n")), case.1);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![
            (
                vec![
                    "#######", "#a.#Cd#", "##...##", "##.@.##", "##...##", "#cB#Ab#", "#######",
                ],
                8,
            ),
            (
                vec![
                    "###############",
                    "#d.ABC.#.....a#",
                    "######...######",
                    "######.@.######",
                    "######...######",
                    "#b.....#.....c#",
                    "###############",
                ],
                24,
            ),
            (
                vec![
                    "#############",
                    "#DcBa.#.GhKl#",
                    "#.###...#I###",
                    "#e#d#.@.#j#k#",
                    "###C#...###J#",
                    "#fEbA.#.FgHi#",
                    "#############",
                ],
                32,
            ),
            (
                vec![
                    "#############",
                    "#g#f.D#..h#l#",
                    "#F###e#E###.#",
                    "#dCba...BcIJ#",
                    "#####.@.#####",
                    "#nK.L...G...#",
                    "#M###N#H###.#",
                    "#o#m..#i#jk.#",
                    "#############",
                ],
                72,
            ),
        ];

        for case in cases {
            assert_eq!(part2(&case.0.join("\n")), case.1);
        }
    }
}
