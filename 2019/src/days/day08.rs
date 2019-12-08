pub fn part1(image: &str) -> usize {
    let width = 25;
    let height = 6;

    let pixels = get_pixels(image);
    let digit_counts = pixels.chunks(width * height).map(|x| count_digits(x));
    let min = digit_counts.min_by_key(|x| x.0).unwrap();
    return min.1 * min.2;
}

pub fn part2(image: &str) -> String {
    let width = 25;
    let height = 6;

    let pixels = get_pixels(image);
    let layers = pixels.chunks(width * height);

    let mut output = vec![2; width * height];
    for layer in layers {
        for (pos, value) in output.iter_mut().enumerate() {
            if *value == 2 {
                *value = layer[pos];
            }
        }
    }

    return get_output(&output, width);
}

fn get_pixels(image: &str) -> Vec<i32> {
    return image
        .trim()
        .chars()
        .map(|x| x.to_digit(10).unwrap() as i32)
        .collect::<Vec<i32>>();
}

fn count_digits(digits: &[i32]) -> (usize, usize, usize) {
    return (
        count_digit(digits, 0),
        count_digit(digits, 1),
        count_digit(digits, 2),
    );
}

fn count_digit(digits: &[i32], digit: i32) -> usize {
    return digits.iter().filter(|&x| *x == digit).count();
}

fn get_output(output: &Vec<i32>, width: usize) -> String {
    return output
        .chunks(width)
        .map(|x| get_output_line(x))
        .collect::<Vec<String>>()
        .join("\n");
}

fn get_output_line(digits: &[i32]) -> String {
    let line_chars: Vec<&str> = digits
        .iter()
        .map(|x| match x {
            0 => " ",
            1 => "#",
            _ => panic!("Invalid color"),
        })
        .collect();

    return line_chars.join("");
}
