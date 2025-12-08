package day07

import (
	"fmt"

	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2025/day07"

var gInput [][]byte
var gStart []int

func SolveP1() string {
	ans := 0
	grid := arrz.Copy2D(gInput)
	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' || grid[i][j] == '|' {
				// assumes first and last cell in a row will not have ^
				// assumes first and last row doesn't have ^
				if grid[i+1][j] == '^' {
					grid[i+1][j-1] = '|'
					grid[i+1][j+1] = '|'
					ans++
				} else {
					grid[i+1][j] = '|'
				}
			}
		}
	}

	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	grid := arrz.Copy2D(gInput)
	// visited := make([][][]byte, 0)
	memory := intz.Init2D(len(grid), len(grid[0]), -1)
	ans = foo(grid, gStart[0], gStart[1], 0, memory)
	return fmt.Sprintf("%v", ans)
}

func foo(grid [][]byte, i int, j int, t int, memory [][]int) int {
	if i+1 == len(grid) {
		return 1
	}
	if memory[i][j] != -1 {
		return memory[i][j]
	}
	ans := 0
	if grid[i+1][j] == '^' {
		ans_1 := foo(grid, i+1, j-1, t+1, memory)
		ans_2 := foo(grid, i+1, j+1, t+1, memory)
		ans = ans_1 + ans_2
	} else {
		ans = foo(grid, i+1, j, t+1, memory)
	}
	memory[i][j] = ans
	return ans
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gStart = arrz.Find2D(gInput, 'S')[0]
}
