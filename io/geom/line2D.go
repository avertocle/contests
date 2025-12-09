package geom

import (
	"math"

	"github.com/avertocle/contests/io/tpz"
)

type Line2D[T tpz.Number] struct {
	m T
	c T
}

func NewLine2D[T tpz.Number](cd, vel *Coord2D[T]) *Line2D[T] {
	m := (vel.Y) / (vel.X)
	c := (cd.Y) - m*(cd.X)
	//fmt.Println(m, c)
	return &Line2D[T]{m: m, c: c}
}

func LineIntersect2D[T tpz.Number](l1, l2 *Line2D[T]) *Coord2D[T] {
	x := (l2.c - l1.c) / (l1.m - l2.m)
	y := l1.m*x + l1.c
	return NewCoord2D[T](x, y)
}

func Dist2D[T tpz.Number](c1, c2 *Coord2D[T]) float64 {
	d2 := (c1.X-c2.X)*(c1.X-c2.X) + (c1.Y-c2.Y)*(c1.Y-c2.Y)
	return math.Sqrt(float64(d2))
}
