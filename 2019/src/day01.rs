pub fn part1(lines: &Vec<&str>) -> i32 {
    lines.iter()
        .map(|l| l.parse::<i32>().unwrap())
        .map(|m| get_fuel_requirements(m))
        .sum()
}

pub fn part2(lines: &Vec<&str>) -> i32 {
    lines.iter()
        .map(|l| l.parse::<i32>().unwrap())
        .map(|m| get_total_fuel_requirements(m))
        .sum()
}

pub fn get_fuel_requirements(mass: i32) -> i32 {
    return (mass / 3) - 2;
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
