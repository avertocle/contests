package day10

import (
	"fmt"
	"math"
	"strings"

	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
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
		indicatorResult := gInputIndicators[i]
		buttons := gInputButtons[i]
		smallestComboLen, success := toggleAndCheckAll(buttons, indicatorResult)
		errz.HardAssert(success, "toggleAndCheckAll failed for input %v", i)
		ans += smallestComboLen
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	inputLen := len(gInputIndicators)
	for i := 0; i < inputLen; i++ {
		joltage := gInputJoltage[i]
		buttons := gInputButtons[i]
		smallestComboLen, success := toggleAndCheckAllP2(buttons, joltage)
		errz.HardAssert(success, "toggleAndCheckAllP2 failed for input %v", i)
		ans += smallestComboLen
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func toggleAndCheckAll(buttons [][]int, indicatorResult []byte) (int, bool) {
	allButtonCombos := generateButtonCombinations(len(buttons))
	smallestComboLen := math.MaxInt32
	success := false
	// fmt.Printf("allButtonCombos : %v\n", allButtonCombos)
	for _, bc := range allButtonCombos {
		buttonsToToggle := extractButtonAsPerCombo(buttons, bc)
		// fmt.Printf("toggleAndCheckAll : %v, %v\n", bc, buttonsToToggle)
		if toggleAndCheck(buttonsToToggle, indicatorResult) {
			smallestComboLen = numz.Min(smallestComboLen, len(bc))
			success = true
		}
	}
	return smallestComboLen, success
}

func toggleAndCheck(buttons [][]int, indicatorResult []byte) bool {
	indicator := bytez.Init1D(len(indicatorResult), '.')
	for _, button := range buttons {
		handleButtonPress(indicator, button)
	}
	return string(indicator) == string(indicatorResult)
}

func handleButtonPress(indicator []byte, wiring []int) bool {
	for _, w := range wiring {
		if indicator[w] == '.' {
			indicator[w] = '#'
		} else {
			indicator[w] = '.'
		}
	}
	return false
}

/***** P2 Functions *****/

func toggleAndCheckAllP2(buttons [][]int, joltage []int) (int, bool) {
	allButtonCombos := generateButtonCombinationsP2(len(buttons))
	fmt.Printf("allButtonCombos : %v\n", allButtonCombos)
	smallestComboLen := math.MaxInt32
	success := false
	// fmt.Printf("allButtonCombos : %v\n", allButtonCombos)
	for _, bc := range allButtonCombos {
		buttonsToToggle := extractButtonAsPerCombo(buttons, bc)
		// fmt.Printf("toggleAndCheckAll : %v, %v\n", bc, buttonsToToggle)
		if toggleAndCheckP2(buttonsToToggle, joltage) {
			smallestComboLen = numz.Min(smallestComboLen, len(bc))
			success = true
		}
	}
	return smallestComboLen, success
}

func toggleAndCheckP2(buttons [][]int, joltageResult []int) bool {
	joltage := intz.Init1D(len(joltageResult), 0)
	for _, button := range buttons {
		handleButtonPressP2(joltage, button)
	}
	return fmt.Sprintf("%v", joltageResult) == fmt.Sprintf("%v", joltage)
}

func handleButtonPressP2(joltage []int, buttons []int) {
	for _, w := range buttons {
		joltage[w] += 1
	}
}

/***** Common Functions *****/

func extractButtonAsPerCombo(buttons [][]int, buttonCombo []int) [][]int {
	extractedButtons := make([][]int, 0)
	for _, bIdx := range buttonCombo {
		extractedButtons = append(extractedButtons, buttons[bIdx])
	}
	return extractedButtons
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
	// allButtonCombinations = make([][]int, 0)
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// arr = arr[:maxButtons]
	// for i := 0; i < maxButtons; i++ {
	// 	combos := GenerateCombinations(arr, i+1)
	// 	allButtonCombinations = append(allButtonCombinations, combos...)
	// }
	// fmt.Println(allButtonCombinations)
	// for i := 0; i < len(gInputIndicators); i++ {
	// 	fmt.Printf("gInputIndicators[%v] : %v\n", i, gInputIndicators[i])
	// 	fmt.Printf("gInputWiring[%v] : %v\n", i, gInputWiring[i])
	// 	fmt.Printf("gInputJoltage[%v] : %v\n", i, gInputJoltage[i])
	// }
}

func generateButtonCombinations(maxButtons int) [][]int {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arr = arr[:maxButtons]
	allCombos := make([][]int, 0)
	for i := 0; i < maxButtons; i++ {
		combos := GenerateCombinations(arr, i+1)
		allCombos = append(allCombos, combos...)
	}
	return allCombos
}

func generateButtonCombinationsP2(maxButtons int) [][]int {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arr = arr[:maxButtons]
	allCombos := make([][]int, 0)
	for i := 0; i < maxButtons; i++ {
		combos := GenerateCombinationsWithRepeats(arr, i+1)
		allCombos = append(allCombos, combos...)
	}
	return allCombos
}

func GenerateCombinations(arr []int, r int) [][]int {
	var result [][]int
	var currentCombo []int

	var recurse func(int, int)
	recurse = func(start int, depth int) {
		if depth == r {
			// Make a copy to avoid slice reference issues
			combo := make([]int, r)
			copy(combo, currentCombo)
			result = append(result, combo)
			return
		}

		for i := start; i < len(arr); i++ {
			currentCombo = append(currentCombo, arr[i])
			recurse(i+1, depth+1)                             // Recurse with the next element
			currentCombo = currentCombo[:len(currentCombo)-1] // Backtrack
		}
	}

	recurse(0, 0)
	return result
}

// GenerateCombinationsWithRepeats returns all length-r combinations where elements
// from arr can repeat (combinations with repetition). Order in each combination
// is non-decreasing relative to arr indices to avoid duplicate permutations.
func GenerateCombinationsWithRepeats(arr []int, r int) [][]int {
	var result [][]int
	var currentCombo []int

	var recurse func(int, int)
	recurse = func(start int, depth int) {
		if depth == r {
			combo := make([]int, r)
			copy(combo, currentCombo)
			result = append(result, combo)
			return
		}

		for i := start; i < len(arr); i++ {
			currentCombo = append(currentCombo, arr[i])
			recurse(i, depth+1)                               // allow repetition of arr[i]
			currentCombo = currentCombo[:len(currentCombo)-1] // Backtrack
		}
	}

	recurse(0, 0)
	return result
}
