/*
Generic main.go for all aoc packages
Just change the import
*/
package main

import (
	"fmt"
	prob "github.com/avertocle/contests/aoc/2023/day25"
	"github.com/avertocle/contests/io/clr"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"os"
	"path"
	"strings"
)

func main() {
	dirPath := prob.DirPath
	args := os.Args
	displayPrettyHeader(dirPath, args)
	problems := makeProblems(dirPath, args)
	runAll(problems)
	displayResults(problems)
}

func displayResults(problems []*problem) {
	m := make(map[string]map[int]*problem)
	for _, p := range problems {
		if _, ok := m[p.fname]; !ok {
			m[p.fname] = make(map[int]*problem)
		}
		m[p.fname][p.part] = p
	}
	for ifName, parts := range m {
		p1, p2 := parts[1], parts[2]
		displayPrettyResult(ifName, p1.ans, p2.ans)
	}
}

func runAll(problems []*problem) {
	for _, p := range problems {
		prob.ParseInput(p.inputFilePath())
		if p.part == 1 {
			p.ans = prob.SolveP1()
		} else if p.part == 2 {
			p.ans = prob.SolveP2()
		} else {
			errz.HardAssert(false, "invalid part : %v", p.part)
		}
	}
}

func makeProblems(dpath string, args []string) []*problem {
	inputFileNames, err := iutils.GetInputFileList(dpath)
	errz.HardAssert(err == nil, "error fetching input file : dir(%v) | %v", dpath, err)
	problems := make([]*problem, 0)
	//inputFileNames = []string{"input_small.txt"}
	for _, fname := range inputFileNames {
		problems = append(problems, newProblem(dpath, fname, 1, "na"))
		problems = append(problems, newProblem(dpath, fname, 2, "na"))
	}
	return problems
}

/***** Interfaces *****/

type problem struct {
	dpath string
	fname string
	part  int
	ans   string
}

func newProblem(dirPath, inputFile string, part int, ans string) *problem {
	return &problem{dpath: dirPath, fname: inputFile, part: part, ans: ans}
}

func (p *problem) inputFilePath() string {
	return path.Join(p.dpath, p.fname)
}

/***** Display Functions *****/

func displayPrettyHeader(dirPath string, args []string) {
	line := clr.Str(fmt.Sprintf("Solving %v | %v", dirPath, args[1:]), clr.Green)
	fmt.Printf("\n%v\n%v\n\n", horLine(), line)
}

func displayPrettyResult(ifName, ansP1, ansP2 string) {
	line := fmt.Sprintf("%v : ans-P1 = %v : ans-P2 = %v",
		clr.Str(ifName, clr.Yellow),
		clr.Str(ansP1, clr.Green),
		clr.Str(ansP2, clr.Green))
	fmt.Printf("\n%v\n%v\n%v\n\n", horLine(), line, horLine())
}

func horLine() string {
	return clr.Str(strings.Repeat("~-", 30)+"~", clr.Yellow)
}
