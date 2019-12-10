use std::f32::consts::PI;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
pub struct Fraction {
    pub numerator: i32,
    pub denomenator: i32,
}

impl Fraction {
    pub fn new(numerator: i32, denomenator: i32) -> Self {
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
        return (self.numerator as f32).atan2(self.denomenator as f32);
    }
}

fn gcd(first: i32, second: i32) -> i32 {
    let mut a = first.abs();
    let mut b = second.abs();

    while a > 0 {
        let remainder = b % a;
        b = a;
        a = remainder;
    }

    return b;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_gcd() {
        assert_eq!(gcd(2, 4), 2);
        assert_eq!(gcd(2, -4), 2);
        assert_eq!(gcd(-2, -4), 2);
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
