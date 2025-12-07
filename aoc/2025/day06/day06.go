package day06

import (
	"fmt"
	"strings"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2025/day06"

var gInputLines []string
var gInputNums [][]int

var gColumnMarkers []int
var gInputNumsP2 [][]byte
var gLastLine string
var gInputOps []string

func SolveP1() string {
	ans := 0
	for j := 0; j < len(gInputNums[0]); j++ {
		temp := gInputNums[0][j]
		for i := 1; i < len(gInputNums); i++ {
			temp = performOperation(gInputOps[j], temp, gInputNums[i][j])
		}
		ans += temp
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	ans = process()
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func performOperation(op string, num1, num2 int) int {
	switch strings.TrimSpace(op) {
	case "+":
		return num1 + num2
	case "*":
		return num1 * num2
	}
	errz.HardAssert(false, "invalid op %v", op)
	return -1
}

/***** P2 Functions *****/

func process() int {
	ans := 0
	for i := 0; i < len(gColumnMarkers)-1; i++ {
		arr := extractByColumnMarkerIndex(i)
		temp := performOperationP2(string(gLastLine[gColumnMarkers[i]]), arr)
		ans += temp
	}
	return ans
}

func extractByColumnMarkerIndex(cmidx int) [][]byte {
	ans := make([][]byte, 0)
	for _, row := range gInputNumsP2 {
		start := gColumnMarkers[cmidx]
		end := gColumnMarkers[cmidx+1]
		ans = append(ans, row[start:end])
	}
	return ans
}

func performOperationP2(op string, arr [][]byte) int {
	ans := 0
	nums := make([]int, 0)
	for j := 0; j < len(arr[0]); j++ {
		tempArr := make([]byte, 0)
		for i := 0; i < len(arr); i++ {
			tempArr = append(tempArr, arr[i][j])
		}
		strr := strings.TrimSpace(string(tempArr))
		if strr != "" {
			nums = append(nums, stringz.AtoI(strr, 0))
		}
	}
	//fmt.Println(nums)
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		ans = performOperation(op, ans, nums[i])
	}
	return ans
}

func byteToInt(b byte, op string) int {
	if b == ' ' {
		if op == "+" {
			return 0
		} else if op == "*" {
			return 1
		}
		return 0
	}
	return int(b - '0')
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputNums = iutils.ExtractInt2DFromString1D(lines[0:len(lines)-1], " ", nil, -1)
	gInputOps = strings.Fields(lines[len(lines)-1])
	gColumnMarkers = findColumnMarkers(lines)
	gColumnMarkers = append(gColumnMarkers, len(lines[0]))
	gInputNumsP2 = iutils.ExtractByte2DFromString1D(lines[0:len(lines)-1], "", nil, 0)
	gLastLine = lines[len(lines)-1]
}

func findColumnMarkers(lines []string) []int {
	ans := make([]int, 0)
	lastLine := lines[len(lines)-1]
	for i := 0; i < len(lastLine); i++ {
		if lastLine[i] != ' ' {
			ans = append(ans, i)
		}
	}
	return ans
}
