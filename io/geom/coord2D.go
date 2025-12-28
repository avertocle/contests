package geom

import (
	"fmt"

	"github.com/avertocle/contests/io/tpz"
)

type Coord2D[T tpz.Number] struct {
	X T
	Y T
}

func (o *Coord2D[T]) IsInside(boundTL, boundBR *Coord2D[T]) bool {
	return o.X >= boundTL.X &&
		o.X <= boundBR.X &&
		o.Y <= boundTL.Y &&
		o.Y >= boundBR.Y
}

func (o *Coord2D[T]) MoveBy(dx, dy T) *Coord2D[T] {
	o.X += dx
	o.Y += dy
	return o
}

func (o *Coord2D[T]) Str() string {
	return fmt.Sprintf("%v,%v", o.X, o.Y)
}

func (o *Coord2D[T]) IsEqual(o1 *Coord2D[T]) bool {
	return o.X == o1.X && o.Y == o1.Y

}

func PPrintCoord2D[T tpz.Number](coords []*Coord2D[T]) {
	for _, c := range coords {
		fmt.Println(c.Str())
	}
}

func NewCoord2D[T tpz.Number](x, y T) *Coord2D[T] {
	return &Coord2D[T]{X: x, Y: y}
}
