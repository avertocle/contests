package day16

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "1651", "1707"},
		{"input_final.txt", "1792", "0"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
