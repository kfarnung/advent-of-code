use std::vec;

pub fn part1(contents: &str) -> i64 {
    let grid = parse_input(contents);
    let mut count = 0;

    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] == '@' && is_movable(&grid, x, y) {
                count += 1;
            }
        }
    }
    count
}

pub fn part2(contents: &str) -> i64 {
    let mut grid = parse_input(contents);
    let mut count = 0;
    let mut found = true;

    while found {
        found = false;
        let mut new_grid = grid.clone();

        for y in 0..grid.len() {
            for x in 0..grid[0].len() {
                if grid[y][x] == '@' && is_movable(&grid, x, y) {
                    new_grid[y][x] = '.';
                    found = true;
                    count += 1;
                }
            }
        }
        grid = new_grid;
    }
    count
}

fn parse_input(contents: &str) -> Vec<Vec<char>> {
    let mut grid = vec![];
    for l in contents.lines() {
        grid.push(l.chars().collect());
    }
    grid
}

fn is_movable(grid: &Vec<Vec<char>>, x: usize, y: usize) -> bool {
    let mut count: usize = 0;
    let rows = grid.len();
    let cols = grid[0].len();

    if x > 0 && y > 0 && grid[y - 1][x - 1] == '@' {
        count += 1;
    }
    if x > 0 && grid[y][x - 1] == '@' {
        count += 1;
    }
    if x > 0 && y < rows - 1 && grid[y + 1][x - 1] == '@' {
        count += 1;
    }
    if x < cols - 1 && grid[y][x + 1] == '@' {
        count += 1;
    }
    if x < cols - 1 && y > 0 && grid[y - 1][x + 1] == '@' {
        count += 1;
    }
    if y > 0 && grid[y - 1][x] == '@' {
        count += 1;
    }
    if y < rows - 1 && grid[y + 1][x] == '@' {
        count += 1;
    }
    if x < cols - 1 && y < rows - 1 && grid[y + 1][x + 1] == '@' {
        count += 1;
    }
    count < 4
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = vec![
            "..@@.@@@@.",
            "@@@.@.@.@@",
            "@@@@@.@.@@",
            "@.@@@@..@.",
            "@@.@@@@.@@",
            ".@@@@@@@.@",
            ".@.@.@.@@@",
            "@.@@@.@@@@",
            ".@@@@@@@@.",
            "@.@.@@@.@.",
        ];

        assert_eq!(part1(&input.join("\n")), 13);
    }

    #[test]
    fn test_part2() {
        let input = vec![
            "..@@.@@@@.",
            "@@@.@.@.@@",
            "@@@@@.@.@@",
            "@.@@@@..@.",
            "@@.@@@@.@@",
            ".@@@@@@@.@",
            ".@.@.@.@@@",
            "@.@@@.@@@@",
            ".@@@@@@@@.",
            "@.@.@@@.@.",
        ];

        assert_eq!(part2(&input.join("\n")), 43);
    }
}
