package day07

import (
	"fmt"

	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/boolz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2025/day07"

var gInput [][]byte
var gStart []int

func SolveP1() string {
	ans := 0
	visitedSplitters := boolz.Init2D(len(gInput), len(gInput[0]), false)
	simulateBeamP1(gInput, gStart[0], gStart[1], visitedSplitters)
	ans = arrz.Count2D(visitedSplitters, true)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	memory := intz.Init2D(len(gInput), len(gInput[0]), -1)
	ans = simulateBeamP2(gInput, gStart[0], gStart[1], memory)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func simulateBeamP1(grid [][]byte, i int, j int, splitHit [][]bool) {
	if i+1 == len(grid) || splitHit[i][j] {
		return
	}
	if grid[i+1][j] == '^' {
		splitHit[i][j] = true
		simulateBeamP1(grid, i+1, j-1, splitHit)
		simulateBeamP1(grid, i+1, j+1, splitHit)
	} else {
		simulateBeamP1(grid, i+1, j, splitHit)
	}
}

/***** P2 Functions *****/

func simulateBeamP2(grid [][]byte, i int, j int, memory [][]int) int {
	if i+1 == len(grid) {
		return 1
	} else if memory[i][j] != -1 {
		return memory[i][j]
	}
	ans := 0
	if grid[i+1][j] == '^' {
		ans += simulateBeamP2(grid, i+1, j-1, memory)
		ans += simulateBeamP2(grid, i+1, j+1, memory)
	} else {
		ans += simulateBeamP2(grid, i+1, j, memory)
	}
	memory[i][j] = ans
	return ans
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gStart = arrz.Find2D(gInput, 'S')[0]
}
