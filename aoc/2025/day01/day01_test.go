package day01

import (
	"testing"

	"github.com/avertocle/contests/aoc/testz"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "3", "6"},
		{"input_final.txt", "1191", "6858"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
