const BASE_PATTERN: [i32; 4] = [0, 1, 0, -1];

pub fn part1(contents: &str) -> usize {
    let mut digits: Vec<u8> = contents
        .trim()
        .chars()
        .map(|x| x.to_digit(10).unwrap() as u8)
        .collect();

    for _i in 0..100 {
        digits = fft_phase(&digits);
    }

    return to_number(&digits, 8);
}

pub fn part2(contents: &str) -> usize {
    let digits: Vec<u8> = contents
        .trim()
        .chars()
        .map(|x| x.to_digit(10).unwrap() as u8)
        .collect();

    let total_digits = digits.len() * 10000;
    let offset = to_number(&digits, 7);

    let mut combined: Vec<u8> = digits
        .iter()
        .cycle()
        .skip(offset % digits.len())
        .take(total_digits - offset)
        .map(|x| *x)
        .collect();

    for _i in 0..100 {
        fft_fast(&mut combined);
    }

    return to_number(&combined, 8);
}

fn fft_phase(input_signal: &Vec<u8>) -> Vec<u8> {
    let mut output_signal = Vec::new();

    for i in 0..input_signal.len() {
        let pattern = expand_pattern(i);
        let pattern_len = pattern.len();
        let mut new_digit = 0;

        for (i, digit) in input_signal.iter().enumerate() {
            new_digit += (*digit as i32) * pattern[(i + 1) % pattern_len];
        }

        new_digit = new_digit.abs();
        new_digit %= 10;
        output_signal.push(new_digit as u8);
    }

    return output_signal;
}

fn expand_pattern(index: usize) -> Vec<i32> {
    let mut expanded = Vec::new();

    for digit in BASE_PATTERN.iter() {
        for _i in 0..=index {
            expanded.push(*digit);
        }
    }

    return expanded;
}

fn fft_fast(input_signal: &mut Vec<u8>) {
    let len = input_signal.len();
    for i in 1..len {
        input_signal[len - i - 1] = (input_signal[len - i] + input_signal[len - i - 1]) % 10;
    }
}

fn to_number(input_signal: &Vec<u8>, count: usize) -> usize {
    return input_signal
        .iter()
        .take(count)
        .fold(0usize, |sum, x| (sum * 10) + (*x as usize));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let cases = vec![
            ("80871224585914546619083218645595", 24176176),
            ("19617804207202209144916044189917", 73745418),
            ("69317163492948606335995924319873", 52432133),
        ];

        for case in cases {
            assert_eq!(part1(&case.0), case.1);
        }
    }

    #[test]
    fn test_part2() {
        let cases = vec![
            ("03036732577212944063491565474664", 84462026),
            ("02935109699940807407585447034323", 78725270),
            ("03081770884921959731165446850517", 53553731),
        ];

        for case in cases {
            assert_eq!(part2(&case.0), case.1);
        }
    }
}
