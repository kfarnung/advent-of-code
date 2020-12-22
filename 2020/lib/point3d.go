package lib

import "fmt"

// Point3D represents a point in 2D space
type Point3D struct {
	X, Y, Z int64
}

// NewPoint3D intitializes a Point3D struct
func NewPoint3D(x, y, z int64) Point3D {
	return Point3D{
		X: x,
		Y: y,
		Z: z,
	}
}

// GetAdjacent returns the list of points adjacent to the current point
func (p Point3D) GetAdjacent() []Point3D {
	var points []Point3D
	for i := p.X - 1; i <= p.X+1; i++ {
		for j := p.Y - 1; j <= p.Y+1; j++ {
			for k := p.Z - 1; k <= p.Z+1; k++ {
				point := NewPoint3D(i, j, k)
				if p != point {
					points = append(points, point)
				}
			}
		}
	}

	return points
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}
