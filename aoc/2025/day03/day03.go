package day03

import (
	"fmt"

	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2025/day03"

var gInput [][]int

func SolveP1() string {
	ans := int64(0)
	batCount := 2
	for _, arr := range gInput {
		ans += findMaxPossibleJoltage(arr, batCount)
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := int64(0)
	batCount := 12
	for _, arr := range gInput {
		ans += findMaxPossibleJoltage(arr, batCount)
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func findMaxPossibleJoltage(arr []int, batCount int) int64 {
	batArr := make([]int, batCount)
	startIdx, endIdx := 0, len(arr)-batCount
	for i := range batArr {
		batArr[i], startIdx = findHighestNumberAfterIndex(arr, startIdx, endIdx)
		startIdx += 1
		endIdx += 1
	}
	num := stringz.AtoI64(arrz.ToStr1D(batArr, ""), 0)
	return num
}

// both startIdx and endIdx are inclusive
func findHighestNumberAfterIndex(arr []int, startIdx int, endIdx int) (int, int) {
	highest := 0
	highestIdx := -1
	for i := startIdx; i <= endIdx; i++ {
		if arr[i] > highest {
			highest = arr[i]
			highestIdx = i
		}
	}
	return highest, highestIdx
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, "", nil, -1)
}
