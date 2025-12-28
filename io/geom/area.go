package geom

import (
	"github.com/avertocle/contests/io/numz"
	"github.com/avertocle/contests/io/tpz"
)

func CalcRectArea2DByBounds[T tpz.Number](diag1, diag2 *Coord2D[T]) T {
	s1 := numz.Abs(diag2.X-diag1.X) + 1
	s2 := numz.Abs(diag2.Y-diag1.Y) + 1
	return s1 * s2
}
