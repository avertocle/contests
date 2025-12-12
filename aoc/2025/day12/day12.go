package day12

import (
	"fmt"
	"strings"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2025/day12"

var gInputBlocks [][][]byte
var gInputGridSize [][]int
var gInputGridConfigs [][]int

func SolveP1() string {
	ans := 0
	for i, gridSize := range gInputGridSize {
		totalBlocks := intz.Sum1D(gInputGridConfigs[i])
		if (gridSize[0]/3)*(gridSize[1]/3) >= totalBlocks {
			ans++
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputBlocks = make([][][]byte, 0)
	for i := 0; i < 30; i += 5 {
		chunk := lines[i+1 : i+4]
		block := make([][]byte, 0)
		for _, c := range chunk {
			block = append(block, []byte(c))
		}
		gInputBlocks = append(gInputBlocks, block)
	}

	gInputGridConfigs = make([][]int, 0)
	gInputGridSize = make([][]int, 0)
	for i := 30; i < len(lines); i++ {
		tokens := strings.Fields(lines[i])
		sizeTokens := strings.Split(tokens[0][0:len(tokens[0])-1], "x")
		size := []int{stringz.AtoI(sizeTokens[0], -1), stringz.AtoI(sizeTokens[1], -1)}
		gInputGridSize = append(gInputGridSize, size)
		configTokens := tokens[1:]
		config := make([]int, 0)
		for _, c := range configTokens {
			config = append(config, stringz.AtoI(c, -1))
		}
		gInputGridConfigs = append(gInputGridConfigs, config)
	}

	// fmt.Println(gInputBlocks)
	// fmt.Println(gInputGridConfigs)
	// fmt.Println(gInputGridSize)
}
