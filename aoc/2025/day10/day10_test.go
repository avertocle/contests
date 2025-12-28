package day10

import (
	"testing"

	"github.com/avertocle/contests/aoc/testz"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "7", "0"},
		{"input_final.txt", "396", "0"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
