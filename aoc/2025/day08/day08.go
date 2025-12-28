package day08

import (
	"fmt"
	"slices"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2025/day08"

var gInput [][]int
var gInputCoords []*geom.Coord3d
var gPairsToProcessP1 int = 0

func SolveP1() string {
	ans := 0
	pairs := makePairsSortedByDistance()
	circuits := groupPairsToCircuitsP1(pairs)
	slices.SortFunc(circuits, circuitSortFuncDesc)
	ans = 1
	for i := 0; i < 3; i++ {
		ans *= circuits[i].length()
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	pairs := makePairsSortedByDistance()
	p := groupPairsToCircuitsP2(pairs, len(gInputCoords))
	ans = p.coord1.X * p.coord2.X
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func groupPairsToCircuitsP1(sortedPairs []*pair) []*circuit {
	circuits := make([]*circuit, 0)
	for _, p := range sortedPairs[0:gPairsToProcessP1] {
		circuits, _ = addOnePairToCircuits(circuits, p)
	}
	return circuits
}

/***** P2 Functions *****/

func groupPairsToCircuitsP2(sortedPairs []*pair, cutOffCircuitSize int) *pair {
	circuits := make([]*circuit, 0)
	var circModified *circuit
	for _, p := range sortedPairs {
		circuits, circModified = addOnePairToCircuits(circuits, p)
		if circModified.length() >= cutOffCircuitSize {
			return p
		}
	}
	return nil
}

/***** Common Functions *****/

func makePairsSortedByDistance() []*pair {
	count := len(gInputCoords)
	distPairs := make([]*pair, 0)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			c3d1, c3d2 := gInputCoords[i], gInputCoords[j]
			dist := geom.EuclidDist3D(c3d1, c3d2)
			distPairs = append(distPairs, newPair(c3d1, c3d2, dist))
		}
	}
	slices.SortFunc(distPairs, func(a, b *pair) int {
		return int(a.dist - b.dist)
	})
	return distPairs
}

func addOnePairToCircuits(circuits []*circuit, p *pair) ([]*circuit, *circuit) {
	circ1 := findCircuitForC3d(circuits, p.coord1)
	circ2 := findCircuitForC3d(circuits, p.coord2)
	if circ1 == nil && circ2 == nil {
		circNew := newCircuit(len(circuits), p)
		circuits = append(circuits, circNew)
		return circuits, circNew
	} else if circ1 == nil {
		circ2.absorbPair(p)
		return circuits, circ2
	} else if circ2 == nil {
		circ1.absorbPair(p)
		return circuits, circ1
	} else {
		circ1.absorbCircuit(circ2)
		return circuits, circ1
	}
}

func findCircuitForC3d(circuits []*circuit, c3d *geom.Coord3d) *circuit {
	for _, c := range circuits {
		if c.hasCoord(c3d) {
			return c
		}
	}
	return nil
}

func printAllCircuits(circuits []*circuit) {
	for _, c := range circuits {
		if c.length() > 0 {
			fmt.Println(c.str())
		}
	}
}

/***** Structs *****/

type pair struct {
	dist   float64
	coord1 *geom.Coord3d
	coord2 *geom.Coord3d
}

func newPair(coord1, coord2 *geom.Coord3d, dist float64) *pair {
	return &pair{dist: dist, coord1: coord1, coord2: coord2}
}

type circuit struct {
	id       int
	coordMap map[string]*geom.Coord3d
}

func newCircuit(id int, dp *pair) *circuit {
	c := &circuit{id: id, coordMap: make(map[string]*geom.Coord3d)}
	c.coordMap[dp.coord1.Str()] = dp.coord1
	c.coordMap[dp.coord2.Str()] = dp.coord2
	return c
}

func (c *circuit) hasCoord(coord *geom.Coord3d) bool {
	_, ok := c.coordMap[coord.Str()]
	return ok
}

func (c *circuit) str() string {
	cmapStr := ""
	for _, c3d := range c.coordMap {
		cmapStr += c3d.Str() + "|"
	}
	return fmt.Sprintf("%v, %v, %v", c.id, c.length(), cmapStr)
}

func (c *circuit) length() int {
	return len(c.coordMap)
}

func (c *circuit) absorbCircuit(c2 *circuit) {
	if c.id == c2.id {
		return
	}
	for _, c3d := range c2.coordMap {
		c.coordMap[c3d.Str()] = c3d
		delete(c2.coordMap, c3d.Str())
	}
}

func (c *circuit) absorbPair(pair *pair) bool {
	c3d1, c3d2 := pair.coord1, pair.coord2
	if c.hasCoord(c3d1) || c.hasCoord(c3d2) {
		c.coordMap[c3d1.Str()] = c3d1
		c.coordMap[c3d2.Str()] = c3d2
		return true
	}
	return false
}

func circuitSortFuncDesc(a, b *circuit) int {
	return int(len(b.coordMap) - len(a.coordMap))
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, ",", nil, 0)
	gInputCoords = make([]*geom.Coord3d, 0)
	for _, input := range gInput {
		gInputCoords = append(gInputCoords, geom.NewCoord3dFromVec(input))
	}
	gPairsToProcessP1 = 10 // input-small
	if len(lines) > 100 {
		gPairsToProcessP1 = 1000 // input-final
	}
}

// 6341067600 = too high
