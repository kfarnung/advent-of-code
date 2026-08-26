pub fn part1(contents: &str) -> i32 {
    let lines: Vec<&str> = contents.lines().collect();
    lines
        .iter()
        .map(|l| l.parse::<i32>().unwrap())
        .map(get_fuel_requirements)
        .sum()
}

pub fn part2(contents: &str) -> i32 {
    let lines: Vec<&str> = contents.lines().collect();
    lines
        .iter()
        .map(|l| l.parse::<i32>().unwrap())
        .map(get_total_fuel_requirements)
        .sum()
}

pub fn get_fuel_requirements(mass: i32) -> i32 {
    (mass / 3) - 2
}

pub fn get_total_fuel_requirements(mass: i32) -> i32 {
    let mut total: i32 = 0;
    let mut current = mass;

    loop {
        current = get_fuel_requirements(current);
        if current > 0 {
            total += current
        } else {
            return total;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_fuel_requirements() {
        assert_eq!(get_fuel_requirements(12), 2);
        assert_eq!(get_fuel_requirements(14), 2);
        assert_eq!(get_fuel_requirements(1969), 654);
        assert_eq!(get_fuel_requirements(100756), 33583);
    }

    #[test]
    fn test_get_total_fuel_requirements() {
        assert_eq!(get_total_fuel_requirements(14), 2);
        assert_eq!(get_total_fuel_requirements(1969), 966);
        assert_eq!(get_total_fuel_requirements(100756), 50346);
    }
}
