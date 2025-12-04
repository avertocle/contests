package day04

import (
	"fmt"

	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2025/day04"

var gInput [][]byte

func SolveP1() string {
	ans := 0
	for i := 0; i < len(gInput); i++ {
		for j := 0; j < len(gInput[i]); j++ {
			if isAccessible(gInput, i, j) {
				ans++
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	grid := bytez.Copy2D(gInput)
	for {
		accessible := findAccessibles(grid)
		if len(accessible) == 0 {
			break
		}
		removeAccessibles(grid, accessible)
		ans += len(accessible)
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func findAccessibles(grid [][]byte) []*arrz.Idx2D[int] {
	accessible := make([]*arrz.Idx2D[int], 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isAccessible(grid, i, j) {
				accessible = append(accessible, arrz.NewIdx2D(i, j))
			}
		}
	}
	return accessible
}

func removeAccessibles(grid [][]byte, idx2ds []*arrz.Idx2D[int]) {
	for _, idx2d := range idx2ds {
		grid[idx2d.I][idx2d.J] = '.'
	}
}

/***** Common Functions *****/

func isAccessible(grid [][]byte, r, c int) bool {
	if grid[r][c] != '@' {
		return false
	}
	nbrs := arrz.NewIdx2D(r, c).Neighbours(true)
	emptyNbrCtr := 0
	for _, nbr := range nbrs {
		if nbr.IsInBounds(len(grid), len(grid[0])) && grid[nbr.I][nbr.J] == '@' {
			emptyNbrCtr++
		}
	}
	return emptyNbrCtr < 4
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
