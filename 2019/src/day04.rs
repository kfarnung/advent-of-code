pub fn part1(contents: &str) -> i32 {
    let (lower, upper) = parse_range(contents);
    let mut valid_count = 0;

    for password in lower..=upper {
        if is_valid_password(password, None) {
            valid_count += 1;
        }
    }

    return valid_count;
}

pub fn part2(contents: &str) -> i32 {
    let (lower, upper) = parse_range(contents);
    let mut valid_count = 0;

    for password in lower..=upper {
        if is_valid_password(password, Some(2)) {
            valid_count += 1;
        }
    }

    return valid_count;
}

fn parse_range(contents: &str) -> (i32, i32) {
    let range: Vec<&str> = contents.trim().split('-').collect();
    return (range[0].parse().unwrap(), range[1].parse().unwrap());
}

fn is_valid_password(password: i32, max_repeat: Option<i32>) -> bool {
    if password > 999999 || password < 100000 {
        // Password must be six digits.
        return false;
    }

    let mut remaining = password;
    let mut repeat = false;
    let mut repeat_count = 1;
    let mut last_digit = i32::max_value();

    loop {
        let current_digit = remaining % 10;

        if current_digit > last_digit {
            // Since we're going backwards, if the current digit is larger than
            // the previous one we know the password is invalid.
            return false;
        } else if current_digit == last_digit {
            repeat_count += 1;
        } else {
            // We found a non-repeating digit, evaluate the repeated group to
            // decide whether it's valid.
            match max_repeat {
                Some(x) => {
                    if repeat_count == x {
                        repeat = true;
                    }
                }
                None => {
                    if repeat_count > 1 {
                        repeat = true;
                    }
                }
            };

            // Reset the repeat count.
            repeat_count = 1;
        }

        if remaining <= 0 {
            // We've run out of work to do.
            break;
        }

        last_digit = current_digit;
        remaining /= 10;
    }

    // Now we just need to make sure we found a repeat somewhere.
    return repeat;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_range() {
        assert_eq!(parse_range("123-456"), (123, 456));
    }

    #[test]
    fn test_is_valid_password() {
        assert_eq!(is_valid_password(99999, None), false);
        assert_eq!(is_valid_password(1000000, None), false);

        assert_eq!(is_valid_password(122345, None), true);
        assert_eq!(is_valid_password(111123, None), true);
        assert_eq!(is_valid_password(135679, None), false);

        assert_eq!(is_valid_password(111111, None), true);
        assert_eq!(is_valid_password(223450, None), false);
        assert_eq!(is_valid_password(123789, None), false);

        assert_eq!(is_valid_password(112233, Some(2)), true);
        assert_eq!(is_valid_password(123444, Some(2)), false);
        assert_eq!(is_valid_password(111122, Some(2)), true);
    }
}
