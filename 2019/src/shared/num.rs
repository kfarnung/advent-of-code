use num::{NumCast, Signed};
use std::f32::consts::PI;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
pub struct Fraction<T>
where
    T: Signed + PartialOrd + Copy + NumCast,
{
    pub numerator: T,
    pub denomenator: T,
}

impl<T> Fraction<T>
where
    T: Signed + PartialOrd + Copy + NumCast,
{
    pub fn new(numerator: T, denomenator: T) -> Self {
        return Self {
            numerator: numerator,
            denomenator: denomenator,
        };
    }

    pub fn reduce(&self) -> Self {
        let gcd = gcd(self.numerator, self.denomenator);
        return Self::new(self.numerator / gcd, self.denomenator / gcd);
    }

    pub fn angle_relative_cw(&self, base: &Self) -> f32 {
        let angle = self.angle() - base.angle();
        if angle >= 0.0 {
            return angle;
        } else {
            return angle + (2.0 * PI);
        }
    }

    fn angle(&self) -> f32 {
        let numerator: f32 = num::cast(self.numerator).unwrap();
        let denomenator: f32 = num::cast(self.denomenator).unwrap();
        return numerator.atan2(denomenator);
    }
}

fn gcd<T>(first: T, second: T) -> T
where
    T: Signed + PartialOrd + Copy,
{
    let mut a = first.abs();
    let mut b = second.abs();

    while a.is_positive() {
        let remainder = b % a;
        b = a;
        a = remainder;
    }

    return b;
}

pub fn lcm<T>(first: T, second: T) -> T
where
    T: Signed + PartialOrd + Copy,
{
    return (first * second) / gcd(first, second);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_gcd() {
        assert_eq!(gcd(2, 4), 2);
        assert_eq!(gcd(2, -4), 2);
        assert_eq!(gcd(-2, -4), 2);
        assert_eq!(gcd(2i64, 4i64), 2i64);
        assert_eq!(gcd(2i64, -4i64), 2i64);
        assert_eq!(gcd(-2i64, -4i64), 2i64);
    }

    #[test]
    fn test_lcm64() {
        assert_eq!(lcm(3, 8), 24);
        assert_eq!(lcm(3i64, 8i64), 24i64);
    }

    #[test]
    fn test_reduce() {
        assert_eq!(Fraction::new(-2, -4).reduce(), Fraction::new(-1, -2));
    }

    #[test]
    fn test_angle() {
        let base_angle = Fraction::new(-1, 0);

        assert_eq!(Fraction::new(-2, 0).angle_relative_cw(&base_angle), 0.0);
        assert_eq!(
            Fraction::new(-2, 2).angle_relative_cw(&base_angle),
            PI / 4.0
        );
        assert_eq!(Fraction::new(0, 2).angle_relative_cw(&base_angle), PI / 2.0);
        assert_eq!(
            Fraction::new(2, 2).angle_relative_cw(&base_angle),
            3.0 * PI / 4.0
        );
        assert_eq!(Fraction::new(2, 0).angle_relative_cw(&base_angle), PI);
        assert_eq!(
            Fraction::new(2, -2).angle_relative_cw(&base_angle),
            5.0 * PI / 4.0
        );
        assert_eq!(
            Fraction::new(0, -2).angle_relative_cw(&base_angle),
            3.0 * PI / 2.0
        );
        assert_eq!(
            Fraction::new(-2, -2).angle_relative_cw(&base_angle),
            7.0 * PI / 4.0
        );
    }
}
