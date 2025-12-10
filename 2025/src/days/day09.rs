struct Point {
    x: i64,
    y: i64,
}

impl Point {
    fn new(x: i64, y: i64) -> Self {
        Point { x, y }
    }

    fn area(&self, other: &Point) -> i64 {
        ((self.x - other.x).abs() + 1) * ((self.y - other.y).abs() + 1)
    }
}

pub fn part1(contents: &str) -> i64 {
    let points = parse_input(contents);
    let mut max_area = 0;
    for (i, p1) in points.iter().enumerate() {
        for p2 in &points[i + 1..] {
            let area = p1.area(p2);
            if area > max_area {
                max_area = area;
            }
        }
    }
    max_area
}

pub fn part2(contents: &str) -> i64 {
    let points = parse_input(contents);
    let mut max_area = 0;
    for (i, p1) in points.iter().enumerate() {
        for p2 in &points[i + 1..] {
            let area = p1.area(p2);
            if area > max_area && is_interior(p1, p2, &points) {
                max_area = area;
            }
        }
    }
    max_area
}

fn parse_input(input: &str) -> Vec<Point> {
    input
        .lines()
        .map(|line| {
            let coords: Vec<i64> = line.split(',').map(|s| s.parse().unwrap()).collect();
            Point::new(coords[0], coords[1])
        })
        .collect()
}

fn is_interior(p1: &Point, p2: &Point, points: &[Point]) -> bool {
    let max_x = if p1.x > p2.x { p1.x } else { p2.x };
    let min_x = if p1.x < p2.x { p1.x } else { p2.x };
    let max_y = if p1.y > p2.y { p1.y } else { p2.y };
    let min_y = if p1.y < p2.y { p1.y } else { p2.y };
    for (i, p) in points.iter().enumerate() {
        if (p.x > min_x && p.x < max_x) && (p.y > min_y && p.y < max_y) {
            // Fully inside
            return false;
        } else {
            let next = &points[(i + 1) % points.len()];
            let min_line_x = if p.x < next.x { p.x } else { next.x };
            let max_line_x = if p.x > next.x { p.x } else { next.x };
            let min_line_y = if p.y < next.y { p.y } else { next.y };
            let max_line_y = if p.y > next.y { p.y } else { next.y };
            if (min_line_x <= min_x
                && max_line_x >= max_x
                && min_line_y > min_y
                && max_line_y < max_y)
                || (min_line_y <= min_y
                    && max_line_y >= max_y
                    && min_line_x > min_x
                    && max_line_x < max_x)
            {
                // Line crosses rectangle
                return false;
            }
        }
    }
    true
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec!["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];

        assert_eq!(part1(&input.join("\n")), 50);
    }

    #[test]
    fn test_part2() {
        let input = vec!["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];

        assert_eq!(part2(&input.join("\n")), 24);
    }
}
