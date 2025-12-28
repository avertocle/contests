package day11

import (
	"testing"

	"github.com/avertocle/contests/aoc/testz"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "5", "0"},
		{"input_small_02.txt", "0", "2"},
		{"input_final.txt", "508", "315116216513280"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
