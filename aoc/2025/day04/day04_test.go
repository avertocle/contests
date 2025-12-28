package day04

import (
	"testing"

	"github.com/avertocle/contests/aoc/testz"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "13", "43"},
		{"input_final.txt", "1489", "8890"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
