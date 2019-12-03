use std::ops::AddAssign;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
pub struct Point2D {
    pub x: i32,
    pub y: i32,
}

impl Point2D {
    pub fn manhattan_distance(&self, other: Self) -> i32 {
        return (self.x - other.x).abs() + (self.y - other.y).abs();
    }
}

impl AddAssign for Point2D {
    fn add_assign(&mut self, other: Self) {
        *self = Self {
            x: self.x + other.x,
            y: self.y + other.y,
        };
    }
}
