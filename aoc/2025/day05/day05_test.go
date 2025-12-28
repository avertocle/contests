package day05

import (
	"testing"

	"github.com/avertocle/contests/aoc/testz"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "3", "14"},
		{"input_final.txt", "885", "348115621205535"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
