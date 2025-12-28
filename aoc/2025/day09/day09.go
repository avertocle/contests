package day09

import (
	"fmt"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
)

const DirPath = "../2025/day09"

var gInput []*geom.Coord2D[int]

func SolveP1() string {
	ans := 0
	maxArea := 0
	for i := 0; i < len(gInput); i++ {
		for j := i + 1; j < len(gInput); j++ {
			area := geom.CalcRectArea2DByBounds(gInput[i], gInput[j])
			maxArea = numz.Max(area, maxArea)
		}
	}
	ans = maxArea
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	maxArea := 0
	polygon := gInput
	sides := makeSides(polygon)
	pointCache := newPointCache()
	for i := 0; i < len(polygon); i++ {
		for j := i + 1; j < len(polygon); j++ {
			diag1, diag2 := polygon[i], polygon[j]
			area := geom.CalcRectArea2DByBounds(diag1, diag2)
			if area <= maxArea {
				continue
			}
			vertices := getRectVertices(diag1, diag2)
			edges := getRectEdges(diag1, diag2)
			isInside := isRectInsidePolygon(vertices, edges, sides, pointCache)
			if isInside {
				maxArea = area
				// fmt.Printf("new-area : %v : [%v - %v]\n", area, diag1.Str(), diag2.Str())
			}
		}
	}
	ans = maxArea
	// fmt.Printf("cache : %v\n", pointCache.Str())
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func getRectVertices(diag1, diag2 *geom.Coord2D[int]) []*geom.Coord2D[int] {
	return []*geom.Coord2D[int]{
		geom.NewCoord2D(diag1.X, diag2.Y),
		geom.NewCoord2D(diag2.X, diag1.Y),
	}
}

func getRectEdges(diag1, diag2 *geom.Coord2D[int]) []*line {
	x := []int{numz.Min(diag1.X, diag2.X), numz.Max(diag1.X, diag2.X)}
	y := []int{numz.Min(diag1.Y, diag2.Y), numz.Max(diag1.Y, diag2.Y)}
	edges := make([]*line, 4)
	edges[0] = newLine(geom.NewCoord2D(x[0], y[0]), geom.NewCoord2D(x[0], y[1]))
	edges[1] = newLine(geom.NewCoord2D(x[0], y[1]), geom.NewCoord2D(x[1], y[1]))
	edges[2] = newLine(geom.NewCoord2D(x[1], y[1]), geom.NewCoord2D(x[1], y[0]))
	edges[3] = newLine(geom.NewCoord2D(x[1], y[0]), geom.NewCoord2D(x[0], y[0]))
	return edges
}

func isRectInsidePolygon(vertices []*geom.Coord2D[int], edges []*line, sides []*line, pointCache *pointCache) bool {
	if !areAllPointInsidePolygon(vertices, sides, pointCache) {
		return false
	}
	for _, edge := range edges {
		if isEdgeFullyContainedInAnySide(edge, sides) {
			continue
		}
		edgePoints := edge.getAllPoints()
		if !areAllPointInsidePolygon(edgePoints, sides, pointCache) {
			return false
		}
	}
	return true
}

func areAllPointInsidePolygon(points []*geom.Coord2D[int], sides []*line, pointCache *pointCache) bool {
	for _, point := range points {
		if !isPointInsidePolygon(point, sides, pointCache) {
			return false
		}
	}
	return true
}

func isPointInsidePolygon(c2d *geom.Coord2D[int], sides []*line, pointCache *pointCache) bool {
	if val, ok := pointCache.get(c2d); ok {
		return val
	}
	intersections := 0
	for _, side := range sides {
		if side.containsPoint(c2d) {
			return true
		} else if checkPointIntersectionWithSide(c2d, side) {
			intersections++
		}
	}
	pointCache.set(c2d, intersections%2 == 1)
	return intersections%2 == 1
}

func checkPointIntersectionWithSide(c2d *geom.Coord2D[int], side *line) bool {
	if side.isHorizontal || c2d.X >= side.start.X {
		return false
	}
	return c2d.Y > side.start.Y && c2d.Y <= side.end.Y
}

func isEdgeFullyContainedInAnySide(edge *line, sides []*line) bool {
	for _, side := range sides {
		if side.containsLine(edge) {
			return true
		}
	}
	return false
}

func makeSides(vertices []*geom.Coord2D[int]) []*line {
	sides := make([]*line, len(vertices))
	for i := 0; i < len(vertices); i++ {
		sides[i] = newLine(vertices[i], vertices[(i+1)%len(vertices)])
	}
	return sides
}

/***** Structs *****/

type pointCache struct {
	cache map[int]map[int]bool
	hits  int64
	miss  int64
	size  int64
}

func (pc *pointCache) Str() string {
	return fmt.Sprintf("[hits = %v : miss = %v : size = %v]", pc.hits, pc.miss, pc.size)
}

func newPointCache() *pointCache {
	return &pointCache{
		cache: make(map[int]map[int]bool),
		hits:  0,
		miss:  0,
		size:  0,
	}
}

func (pc *pointCache) get(p *geom.Coord2D[int]) (bool, bool) {
	if val, ok := pc.cache[p.X][p.Y]; ok {
		pc.hits++
		return val, true
	}
	pc.miss++
	return false, false
}

func (pc *pointCache) set(p *geom.Coord2D[int], val bool) {
	if _, ok := pc.cache[p.X]; !ok {
		pc.cache[p.X] = make(map[int]bool)
	}
	pc.size++
	pc.cache[p.X][p.Y] = val
}

type line struct {
	start        *geom.Coord2D[int]
	end          *geom.Coord2D[int]
	isHorizontal bool
}

func (l *line) Str() string {
	return fmt.Sprintf("%v -> %v", l.start.Str(), l.end.Str())
}

func newLine(p1, p2 *geom.Coord2D[int]) *line {
	isHorizontal := p1.Y == p2.Y
	var start, end *geom.Coord2D[int]
	if isHorizontal {
		if p1.X < p2.X {
			start = p1
			end = p2
		} else {
			start = p2
			end = p1
		}
	} else {
		if p1.Y < p2.Y {
			start = p1
			end = p2
		} else {
			start = p2
			end = p1
		}
	}
	return &line{
		start:        start,
		end:          end,
		isHorizontal: isHorizontal,
	}
}

func (l *line) length() int {
	if l.isHorizontal {
		return l.end.X - l.start.X + 1
	} else {
		return l.end.Y - l.start.Y + 1
	}
}

func (l line) containsPoint(p *geom.Coord2D[int]) bool {
	if l.isHorizontal {
		return p.Y == l.start.Y && p.X >= l.start.X && p.X <= l.end.X
	} else {
		return p.X == l.start.X && p.Y >= l.start.Y && p.Y <= l.end.Y
	}
}

// assumes l2 is smaller than l1 : input guarentees this
func (l *line) containsLine(l2 *line) bool {
	if l.isHorizontal && l2.isHorizontal {
		return l.start.Y == l2.start.Y && l.start.X <= l2.start.X && l.end.X >= l2.end.X
	} else if !l.isHorizontal && !l2.isHorizontal {
		return l.start.X == l2.start.X && l.start.Y <= l2.start.Y && l.end.Y >= l2.end.Y
	} else {
		return false
	}
}

func (l *line) getAllPoints() []*geom.Coord2D[int] {
	points := make([]*geom.Coord2D[int], l.length())
	if l.isHorizontal {
		for i := l.start.X; i <= l.end.X; i++ {
			points[i-l.start.X] = geom.NewCoord2D(i, l.start.Y)
		}
	} else {
		for i := l.start.Y; i <= l.end.Y; i++ {
			points[i-l.start.Y] = geom.NewCoord2D(l.start.X, i)
		}
	}
	return points
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	temp := iutils.ExtractInt2DFromString1D(lines, ",", nil, -1)
	gInput = make([]*geom.Coord2D[int], len(temp))
	for i, row := range temp {
		gInput[i] = geom.NewCoord2D(row[0], row[1])
	}
}

/***** DISCARDED FUNCTIONS *****/

// unused
func reduceEdgeLengthToBeChecked(edge *line, sides []*line) *line {
	reducedEdge := edge
	for _, side := range sides {
		reducedEdge, _ = side.eatOverlapAndReduce(reducedEdge)
	}
	return reducedEdge
}

// assumes l2 is smaller than l1
func (l *line) eatOverlapAndReduce(l2 *line) (*line, bool) {
	if l.isHorizontal && l2.isHorizontal && l.start.Y == l2.start.Y {
		if l2.start.X < l.start.X && l2.end.X >= l.start.X && l2.end.X <= l.end.X {
			// return l2.end.X - l.start.X + 1, true
			return newLine(l2.start, l.start), true
		} else if l2.end.X > l.end.X && l2.start.X >= l.start.X && l2.start.X <= l.end.X {
			// return l.end.X - l2.start.X + 1, true
			return newLine(l.end, l2.end), true
		}
	} else if !l.isHorizontal && !l2.isHorizontal && l.start.X == l2.start.X {
		if l2.start.Y < l.start.Y && l2.end.Y >= l.start.Y && l2.end.Y <= l.end.Y {
			// return l2.end.Y - l.start.Y + 1, true
			return newLine(l2.start, l.start), true
		} else if l2.end.Y > l.end.Y && l2.start.Y >= l.start.Y && l2.start.Y <= l.end.Y {
			// return l.end.Y - l2.start.Y + 1, true
			return newLine(l.end, l2.end), true
		}
	}
	return l2, false
}
