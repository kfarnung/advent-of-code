package lib

import "fmt"

// Point4D represents a point in 2D space
type Point4D struct {
	W, X, Y, Z int64
}

// NewPoint4D intitializes a Point4D struct
func NewPoint4D(w, x, y, z int64) Point4D {
	return Point4D{
		W: w,
		X: x,
		Y: y,
		Z: z,
	}
}

// GetAdjacent returns the list of points adjacent to the current point
func (p Point4D) GetAdjacent() []Point4D {
	var points []Point4D
	for h := p.W - 1; h <= p.W+1; h++ {
		for i := p.X - 1; i <= p.X+1; i++ {
			for j := p.Y - 1; j <= p.Y+1; j++ {
				for k := p.Z - 1; k <= p.Z+1; k++ {
					point := NewPoint4D(h, i, j, k)
					if p != point {
						points = append(points, point)
					}
				}
			}
		}
	}

	return points
}

func (p Point4D) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}
