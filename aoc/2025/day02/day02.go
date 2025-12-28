package day02

import (
	"fmt"
	"strings"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2025/day02"

var gInput [][]int

func SolveP1() string {
	ans := 0
	for _, arr := range gInput {
		for num := arr[0]; num <= arr[1]; num++ {
			if isWierdPalindromeP1(num) {
				ans += num
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for _, arr := range gInput {
		for num := arr[0]; num <= arr[1]; num++ {
			if isWierdPalindromeP2(num) {
				ans += num
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func isWierdPalindromeP1(num int) bool {
	str := fmt.Sprintf("%v", num)
	mid := len(str) / 2
	return str[0:mid] == str[mid:]
}

/***** P2 Functions *****/

func isWierdPalindromeP2(num int) bool {
	str := fmt.Sprintf("%v", num)
	for i := 1; i <= len(str)/2; i++ {
		if isRepeatedByN(str, i) {
			return true
		}
	}
	return false
}

func isRepeatedByN(str string, n int) bool {
	if len(str)%n != 0 || len(str) <= n {
		return false
	}
	start := str[0:n]
	for i := n; i < len(str); i += n {
		next := str[i : i+n]
		if next != start {
			return false
		}
	}
	return true
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(strings.Split(lines[0], ","), "-", nil, -1)
}
