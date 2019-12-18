use std::cmp::max;
use std::collections::HashMap;
use std::collections::VecDeque;

struct Chemical {
    pub name: String,
    pub qty: i64,
}

struct Reaction {
    pub inputs: Vec<Chemical>,
    pub output: Chemical,
}

struct Reactions {
    map: HashMap<String, Reaction>,
}

impl Chemical {
    pub fn parse(line: &str) -> Self {
        let segments: Vec<&str> = line.split_whitespace().collect();
        let qty = segments[0].parse::<i64>().unwrap();
        let name = segments[1];

        return Self {
            name: name.to_string(),
            qty: qty,
        };
    }
}

impl Reaction {
    pub fn parse(line: &str) -> Self {
        let parts: Vec<&str> = line.split(" => ").collect();
        let inputs: Vec<Chemical> = parts[0].split(", ").map(|x| Chemical::parse(x)).collect();
        let output = Chemical::parse(parts[1]);

        return Self {
            inputs: inputs,
            output: output,
        };
    }
}

impl Reactions {
    fn new() -> Self {
        return Self {
            map: HashMap::new(),
        };
    }

    pub fn load(contents: &str) -> Self {
        let reactions: Vec<Reaction> = contents.lines().map(|x| Reaction::parse(x)).collect();

        let mut result = Reactions::new();
        for reaction in reactions {
            result.map.insert(reaction.output.name.clone(), reaction);
        }

        return result;
    }

    pub fn get(&self, name: &str) -> &Reaction {
        return self.map.get(name).unwrap();
    }
}

pub fn part1(contents: &str) -> i64 {
    let reactions = Reactions::load(contents);
    return calculate_ore(&reactions, 1);
}

pub fn part2(contents: &str) -> i64 {
    let reactions = Reactions::load(contents);

    // Find the upper bounds.
    let mut upper = 1;

    while calculate_ore(&reactions, upper) <= 1000000000000 {
        upper *= 2;
    }

    let mut lower = upper / 2;

    while lower < upper - 1 {
        let fuel_count = lower + (upper - lower) / 2;
        if calculate_ore(&reactions, fuel_count) > 1000000000000 {
            upper = fuel_count;
        } else {
            lower = fuel_count;
        }
    }

    return lower;
}

fn calculate_ore(reactions: &Reactions, fuel_count: i64) -> i64 {
    let mut ore_count = 0;
    let mut queue = VecDeque::new();
    let mut leftover_chemicals = HashMap::new();

    queue.push_back(Chemical {
        name: "FUEL".to_string(),
        qty: fuel_count,
    });

    while !queue.is_empty() {
        let current = queue.pop_front().unwrap();
        let mut current_qty = current.qty;

        // Check for any leftover chemicals.
        if let Some(current_leftover) = leftover_chemicals.get_mut(&current.name) {
            let updated_leftover = max(0, *current_leftover - current_qty);
            current_qty -= *current_leftover;
            *current_leftover = updated_leftover;
        }

        if current_qty > 0 {
            let reaction = reactions.get(&current.name);
            let multiplier = (current_qty + reaction.output.qty - 1) / reaction.output.qty;

            // Figure out how many extras we end up with.
            let remainder = (reaction.output.qty * multiplier) - current_qty;
            if remainder > 0 {
                leftover_chemicals.insert(current.name, remainder);
            }

            for input in &reaction.inputs {
                let required_qty = input.qty * multiplier;

                if input.name != "ORE" {
                    queue.push_back(Chemical {
                        name: input.name.clone(),
                        qty: required_qty,
                    });
                } else {
                    ore_count += required_qty;
                }
            }
        }
    }

    return ore_count;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let cases = vec![
            (
                vec![
                    "10 ORE => 10 A",
                    "1 ORE => 1 B",
                    "7 A, 1 B => 1 C",
                    "7 A, 1 C => 1 D",
                    "7 A, 1 D => 1 E",
                    "7 A, 1 E => 1 FUEL",
                ],
                31,
            ),
            (
                vec![
                    "9 ORE => 2 A",
                    "8 ORE => 3 B",
                    "7 ORE => 5 C",
                    "3 A, 4 B => 1 AB",
                    "5 B, 7 C => 1 BC",
                    "4 C, 1 A => 1 CA",
                    "2 AB, 3 BC, 4 CA => 1 FUEL",
                ],
                165,
            ),
            (
                vec![
                    "157 ORE => 5 NZVS",
                    "165 ORE => 6 DCFZ",
                    "44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
                    "12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ",
                    "179 ORE => 7 PSHF",
                    "177 ORE => 5 HKGWZ",
                    "7 DCFZ, 7 PSHF => 2 XJWVT",
                    "165 ORE => 2 GPVTF",
                    "3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
                ],
                13312,
            ),
            (
                vec![
                    "2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG",
                    "17 NVRVD, 3 JNWZP => 8 VPVL",
                    "53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL",
                    "22 VJHF, 37 MNCFX => 5 FWMGM",
                    "139 ORE => 4 NVRVD",
                    "144 ORE => 7 JNWZP",
                    "5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC",
                    "5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV",
                    "145 ORE => 6 MNCFX",
                    "1 NVRVD => 8 CXFTF",
                    "1 VJHF, 6 MNCFX => 4 RFSQX",
                    "176 ORE => 6 VJHF",
                ],
                180697,
            ),
            (
                vec![
                    "171 ORE => 8 CNZTR",
                    "7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL",
                    "114 ORE => 4 BHXH",
                    "14 VRPVC => 6 BMBT",
                    "6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL",
                    "6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT",
                    "15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW",
                    "13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW",
                    "5 BMBT => 4 WPTQ",
                    "189 ORE => 9 KTJDG",
                    "1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP",
                    "12 VRPVC, 27 CNZTR => 2 XDBXC",
                    "15 KTJDG, 12 BHXH => 5 XCVML",
                    "3 BHXH, 2 VRPVC => 7 MZWV",
                    "121 ORE => 7 VRPVC",
                    "7 XCVML => 6 RJRHP",
                    "5 BHXH, 4 VRPVC => 5 LTCX",
                ],
                2210736,
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
                    "157 ORE => 5 NZVS",
                    "165 ORE => 6 DCFZ",
                    "44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
                    "12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ",
                    "179 ORE => 7 PSHF",
                    "177 ORE => 5 HKGWZ",
                    "7 DCFZ, 7 PSHF => 2 XJWVT",
                    "165 ORE => 2 GPVTF",
                    "3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
                ],
                82892753,
            ),
            (
                vec![
                    "2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG",
                    "17 NVRVD, 3 JNWZP => 8 VPVL",
                    "53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL",
                    "22 VJHF, 37 MNCFX => 5 FWMGM",
                    "139 ORE => 4 NVRVD",
                    "144 ORE => 7 JNWZP",
                    "5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC",
                    "5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV",
                    "145 ORE => 6 MNCFX",
                    "1 NVRVD => 8 CXFTF",
                    "1 VJHF, 6 MNCFX => 4 RFSQX",
                    "176 ORE => 6 VJHF",
                ],
                5586022,
            ),
            (
                vec![
                    "171 ORE => 8 CNZTR",
                    "7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL",
                    "114 ORE => 4 BHXH",
                    "14 VRPVC => 6 BMBT",
                    "6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL",
                    "6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT",
                    "15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW",
                    "13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW",
                    "5 BMBT => 4 WPTQ",
                    "189 ORE => 9 KTJDG",
                    "1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP",
                    "12 VRPVC, 27 CNZTR => 2 XDBXC",
                    "15 KTJDG, 12 BHXH => 5 XCVML",
                    "3 BHXH, 2 VRPVC => 7 MZWV",
                    "121 ORE => 7 VRPVC",
                    "7 XCVML => 6 RJRHP",
                    "5 BHXH, 4 VRPVC => 5 LTCX",
                ],
                460664,
            ),
        ];

        for case in cases {
            assert_eq!(part2(&case.0.join("\n")), case.1);
        }
    }
}
