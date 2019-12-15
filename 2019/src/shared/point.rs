use crate::shared::num::Fraction;
use std::ops::AddAssign;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
pub struct Point2D {
    pub x: i32,
    pub y: i32,
}

impl Point2D {
    pub fn new(x: i32, y: i32) -> Self {
        return Self { x: x, y: y };
    }

    pub fn manhattan_distance(&self, other: &Self) -> i32 {
        return (self.x - other.x).abs() + (self.y - other.y).abs();
    }

    pub fn slope(&self, other: &Self) -> Fraction {
        return Fraction::new(other.y - self.y, other.x - self.x).reduce();
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

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
pub struct Point3D {
    pub x: i32,
    pub y: i32,
    pub z: i32,
}

impl Point3D {
    pub fn new(x: i32, y: i32, z: i32) -> Self {
        return Self { x: x, y: y, z: z };
    }

    pub fn manhattan_distance(&self, other: &Self) -> i32 {
        return (self.x - other.x).abs() + (self.y - other.y).abs() + (self.z - other.z).abs();
    }
}

impl AddAssign for Point3D {
    fn add_assign(&mut self, other: Self) {
        *self = Self {
            x: self.x + other.x,
            y: self.y + other.y,
            z: self.z + other.z,
        };
    }
}
