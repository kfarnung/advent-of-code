package lib

// Point2D represents a point in 2D space
type Point2D struct {
	X int
	Y int
}

// NewPoint2D intitializes a Point2D struct
func NewPoint2D(x int, y int) Point2D {
	return Point2D{
		X: x,
		Y: y,
	}
}

func (p *Point2D) Add(x int, y int) {
	p.X += x
	p.Y += y
}

func (p *Point2D) Rotate90DegreesClockwise(count int) {
	for i := 0; i < count; i++ {
		x := p.X
		p.X = p.Y
		p.Y = -x
	}
}

func (p *Point2D) Rotate90DegreesCounterClockwise(count int) {
	for i := 0; i < count; i++ {
		x := p.X
		p.X = -p.Y
		p.Y = x
	}
}

func (p Point2D) ManhattanDistance(other Point2D) int {
	return AbsInt(other.X-p.X) + AbsInt(other.Y-p.Y)
}
