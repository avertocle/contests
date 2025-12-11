package day10

import (
	"fmt"
	"math"
	"strings"

	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
)

const DirPath = "../2025/day10"

var gInputIndicators [][]byte
var gInputButtons [][][]int
var gInputJoltage [][]int

func SolveP1() string {
	ans := 0
	inputLen := len(gInputIndicators)
	for i := 0; i < inputLen; i++ {
		buttons, indicatorStateFinal := gInputButtons[i], gInputIndicators[i]
		smallestComboLen := toggleAndCheckAll(buttons, indicatorStateFinal)
		ans += smallestComboLen
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func toggleAndCheckAll(buttons [][]int, indicatorStateFinal []byte) int {
	buttonCombos := generateAllLengthCombinations(len(buttons))
	minComboLen := math.MaxInt32
	for _, bc := range buttonCombos {
		if toggleAndCheck(bc, buttons, indicatorStateFinal) {
			minComboLen = numz.Min(minComboLen, len(bc))
		}
	}
	errz.HardAssert(minComboLen < math.MaxInt32, "toggleAndCheckAll failed for input : %v %v", buttons, indicatorStateFinal)
	return minComboLen
}

func toggleAndCheck(buttonIdxs []int, buttons [][]int, indicatorResult []byte) bool {
	indicator := bytez.Init1D(len(indicatorResult), '.')
	for _, buttonIdx := range buttonIdxs {
		button := buttons[buttonIdx]
		handleButtonPress(indicator, button)
	}
	return string(indicator) == string(indicatorResult)
}

func handleButtonPress(indicator []byte, button []int) bool {
	for _, w := range button {
		if indicator[w] == '.' {
			indicator[w] = '#'
		} else {
			indicator[w] = '.'
		}
	}
	return false
}

/***** P2 Functions *****/

/***** Common Functions *****/

func generateAllLengthCombinations(maxButtons int) [][]int {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arr = arr[:maxButtons]
	allCombos := make([][]int, 0)
	for i := 0; i < maxButtons; i++ {
		generateCombinationsForArray(arr, i+1, 0, 0, new([]int), &allCombos)
	}
	return allCombos
}

func generateCombinationsForArray(arr []int, r int, start int, depth int, currentCombo *[]int, result *[][]int) {
	if depth == r {
		combo := make([]int, r)
		copy(combo, *currentCombo)
		*result = append(*result, combo)
		return
	}

	for i := start; i < len(arr); i++ {
		*currentCombo = append(*currentCombo, arr[i])
		generateCombinationsForArray(arr, r, i+1, depth+1, currentCombo, result)
		*currentCombo = (*currentCombo)[:len(*currentCombo)-1]
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputIndicators = make([][]byte, len(lines))
	gInputButtons = make([][][]int, len(lines))
	gInputJoltage = make([][]int, len(lines))
	maxButtons := 0
	for i, line := range lines {
		tokens := strings.Fields(line)
		gInputIndicators[i] = make([]byte, 0)
		gInputButtons[i] = make([][]int, 0)
		gInputJoltage[i] = make([]int, 0)
		for _, token := range tokens {
			tokenCleaned := strings.Trim(token, "()[]{}")
			if strings.HasPrefix(token, "[") {
				gInputIndicators[i] = []byte(tokenCleaned)
			} else if strings.HasPrefix(token, "(") {
				gInputButtons[i] = append(gInputButtons[i], iutils.ExtractInt1DFromString0D(tokenCleaned, ",", -1))
			} else if strings.HasPrefix(token, "{") {
				gInputJoltage[i] = iutils.ExtractInt1DFromString0D(tokenCleaned, ",", -1)
			}
		}
		maxButtons = numz.Max(maxButtons, len(gInputButtons[i]))
	}
}
