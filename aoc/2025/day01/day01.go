package day01

import (
	"fmt"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
)

const DirPath = "../2025/day01"

var gInput []int

func SolveP1() string {
	ans := 0
	start := 50
	for i := 0; i < len(gInput); i++ {
		start += gInput[i]
		start = fitToRange(start)
		if start == 0 {
			ans++
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	start := 50
	for i := 0; i < len(gInput); i++ {
		val, mult := reduceDistance(gInput[i])
		ans += mult
		start += val
		if start != val && (start <= 0 || start >= 100) {
			ans++
		}
		start = fitToRange(start)
	}
	return fmt.Sprintf("%v", ans)
}

func reduceDistance(val int) (int, int) {
	sign := val / numz.Abs(val)
	mult := numz.Abs(val) / 100
	val = (numz.Abs(val) % 100) * sign
	return val, mult
}

func fitToRange(val int) int {
	return (val + 100) % 100
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make([]int, len(lines))
	for i, line := range lines {
		steps := iutils.ExtractInt1DFromString0D(line[1:], " ", -1)[0]
		if line[0] == 'L' {
			gInput[i] = -1 * steps
		} else {
			gInput[i] = steps
		}
	}
}
