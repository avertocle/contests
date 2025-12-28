package day06

import (
	"testing"

	"github.com/avertocle/contests/aoc/testz"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "4277556", "3263827"},
		{"input_final.txt", "6299564383938", "11950004808442"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
