package lib

import "fmt"

// Point2D represents a point in 2D space
type Point2D struct {
	X, Y int64
}

// NewPoint2D intitializes a Point2D struct
func NewPoint2D(x, y int64) Point2D {
	return Point2D{
		X: x,
		Y: y,
	}
}

// Add to the coordinates of the point
func (p *Point2D) Add(x, y int64) {
	p.X += x
	p.Y += y
}

// Rotate90DegreesClockwise rotates the point 90 degrees clockwise about the origin
func (p *Point2D) Rotate90DegreesClockwise(count int) {
	for i := 0; i < count; i++ {
		x := p.X
		p.X = p.Y
		p.Y = -x
	}
}

// Rotate90DegreesCounterClockwise rotates the point 90 degrees counter-clockwise about the origin
func (p *Point2D) Rotate90DegreesCounterClockwise(count int) {
	for i := 0; i < count; i++ {
		x := p.X
		p.X = -p.Y
		p.Y = x
	}
}

// ManhattanDistance calculates the manhattan distance between two points
func (p Point2D) ManhattanDistance(other Point2D) int64 {
	return AbsInt64(other.X-p.X) + AbsInt64(other.Y-p.Y)
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
