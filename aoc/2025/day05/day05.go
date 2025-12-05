package day05

import (
	"fmt"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/rangez"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2025/day05"

var gInput []int
var gInputRanges [][]int

func SolveP1() string {
	ans := 0
	for _, x := range gInput {
		for _, r := range gInputRanges {
			if rangez.IsInside(r, x) {
				ans++
				break
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	baseRanges := make([][]int, 0)
	for _, r := range gInputRanges {
		baseRanges = rangez.Union1D(baseRanges, r)
	}
	for _, r := range baseRanges {
		ans += r[1] - r[0] + 1
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	breakIdx := stringz.FindEmpty1D(lines)[0]
	gInputRanges = make([][]int, 0)
	gInput = iutils.ExtractInt1DFromString1D(lines[breakIdx+2:], "", -1, int(-1))
	gInputRanges = iutils.ExtractInt2DFromString1D(lines[0:breakIdx], "-", nil, int(-1))
	// fmt.Printf("gInput: %v\n", gInput)
	// fmt.Printf("gInputRanges: %v\n", gInputRanges)
}
