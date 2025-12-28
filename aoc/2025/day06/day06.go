package day06

import (
	"fmt"

	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2025/day06"

var gInputGrid [][]byte
var gColMarkers []int // ops are also column-end markers
var gInputOps []byte

func SolveP1() string {
	ans := 0
	for i := 0; i < len(gColMarkers); i++ {
		chunk := extractChunk(gInputGrid, gColMarkers, i)
		ans += calculateChunkValue(chunk, gInputOps[i])
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for i := 0; i < len(gColMarkers); i++ {
		chunk := extractChunk(gInputGrid, gColMarkers, i)
		chunk = arrz.Transpose2D(chunk)
		ans += calculateChunkValue(chunk, gInputOps[i])
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

// cmIdx : column-end marker index
func extractChunk(grid [][]byte, colMarkers []int, chunkIdx int) [][]byte {
	sCol := colMarkers[chunkIdx]
	eCol := len(grid[0])
	if chunkIdx+1 < len(colMarkers) {
		eCol = colMarkers[chunkIdx+1] - 1 // skip the empty column between chunks
	}
	ans := make([][]byte, 0)
	for _, row := range grid {
		ans = append(ans, row[sCol:eCol])
	}
	return ans
}

func calculateChunkValue(chunk [][]byte, op byte) int {
	ans := getInitValueByOp(op)
	for _, row := range chunk {
		ans = performOp(op, ans, bytez.AtoI(row))
	}
	return ans
}

func performOp(op byte, n1, n2 int) int {
	ans := 0
	switch op {
	case '+':
		ans = n1 + n2
	case '*':
		ans = n1 * n2
	default:
		errz.HardAssert(false, "invalid op %v", op)
	}
	// fmt.Printf("%v %v %v = %v\n", n1, string(op), n2, ans)
	return ans
}

func getInitValueByOp(op byte) int {
	switch op {
	case '+':
		return 0
	case '*':
		return 1
	}
	errz.HardAssert(false, "invalid op %v", op)
	return -1
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputGrid = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	opLine := gInputGrid[len(gInputGrid)-1]
	gInputGrid = gInputGrid[0 : len(gInputGrid)-1] // remove op line
	gColMarkers, gInputOps = findColumnMarkersAndOps(opLine)
}

func findColumnMarkersAndOps(opLine []byte) ([]int, []byte) {
	colmarkers := make([]int, 0)
	ops := make([]byte, 0)
	for i, b := range opLine {
		if b != ' ' {
			colmarkers = append(colmarkers, i)
			ops = append(ops, b)
		}
	}
	return colmarkers, ops
}
