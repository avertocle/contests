package geom

import (
	"fmt"
	"math"

	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/numz"
)

/*
returns ref for easy chaining
*/

type Coord3d struct {
	X int
	Y int
	Z int
}

func NewCoord3d(x, y, z int) *Coord3d {
	return &Coord3d{X: x, Y: y, Z: z}
}

func NewCoord3dFromVec(v []int) *Coord3d {
	return &Coord3d{X: v[0], Y: v[1], Z: v[2]}
}

func (o *Coord3d) MoveBy(vec []int) *Coord3d {
	o.X += vec[0]
	o.Y += vec[1]
	o.Z += vec[2]
	return o
}

func (o *Coord3d) Trim(bounds [][]int) *Coord3d {
	o.X = numz.Trim(o.X, bounds[0])
	o.Y = numz.Trim(o.Y, bounds[1])
	o.Z = numz.Trim(o.Z, bounds[2])
	return o
}

func (o *Coord3d) InBounds(bounds [][]int) bool {
	return intz.InBounds3D(o.Arr(), bounds)
}

func (o *Coord3d) Arr() []int {
	return []int{o.X, o.Y, o.Z}
}

func (o *Coord3d) Str() string {
	return fmt.Sprintf("%v,%v,%v", o.X, o.Y, o.Z)
}

func C3DToBounds(cb, ce *Coord3d) [][]int {
	return [][]int{
		{cb.X, ce.X},
		{cb.Y, ce.Y},
		{cb.Z, ce.Z},
	}
}

func EuclidDist3D(c1, c2 *Coord3d) float64 {
	return math.Sqrt(
		math.Pow(float64(c1.X-c2.X), 2) +
			math.Pow(float64(c1.Y-c2.Y), 2) +
			math.Pow(float64(c1.Z-c2.Z), 2))
}
