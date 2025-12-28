package day11

import (
	"fmt"
	"strings"

	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2025/day11"

var gInputGraph map[string][]string

func SolveP1() string {
	ans := 0
	pathCount := new(int)
	countAllPathsP1("you", "out", make(map[string]bool), pathCount)
	ans = *pathCount
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	segments := [][]string{
		{"svr", "dac"},
		{"dac", "fft"},
		{"fft", "out"},
		{"svr", "fft"},
		{"fft", "dac"},
		{"dac", "out"},
	}
	pathCounts := make([]int, len(segments))
	for i, s := range segments {
		dp := make(map[string]int)
		visited := make(map[string]bool)
		countAllPathsP2(s[0], s[1], visited, dp)
		pathCounts[i] = dp[s[0]+"-"+s[1]]
	}
	ans1 := pathCounts[0] * pathCounts[1] * pathCounts[2]
	ans2 := pathCounts[3] * pathCounts[4] * pathCounts[5]
	ans = ans1 + ans2
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func countAllPathsP1(start, end string, vis map[string]bool, pathCount *int) {
	if start == end {
		*pathCount++
		return
	}
	vis[start] = true
	nbrs := gInputGraph[start]
	for _, nbr := range nbrs {
		if vis[nbr] {
			continue
		}
		countAllPathsP1(nbr, end, vis, pathCount)
	}
	vis[start] = false
}

/***** P2 Functions *****/

func countAllPathsP2(start, end string, vis map[string]bool, dp map[string]int) {
	if vis[start] || start == end {
		return
	}
	vis[start] = true
	nbrs := gInputGraph[start]
	for _, nbr := range nbrs {
		incrementDp(start, nbr, 1, dp)
	}
	for _, nbr := range nbrs {
		countAllPathsP2(nbr, end, vis, dp)
	}
	for _, nbr := range nbrs {
		incrementDp(start, end, dp[nbr+"-"+end], dp)
	}
}

func incrementDp(k1, k2 string, val int, dp map[string]int) {
	k := k1 + "-" + k2
	if _, ok := dp[k]; !ok {
		dp[k] = 0
	}
	dp[k] = dp[k] + val
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputGraph = make(map[string][]string)
	// assumes no duplicate vertices in first column
	for _, line := range lines {
		tokens := strings.Fields(line)
		v := tokens[0][:len(tokens[0])-1]
		gInputGraph[v] = tokens[1:]
	}
}

func PrintDP(dp map[string]int) {
	for k, v := range dp {
		fmt.Printf("%v=%v | ", k, v)
	}
	fmt.Println()
}
