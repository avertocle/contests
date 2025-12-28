package stringz

import "fmt"

func Init2D(rows, cols int) [][]string {
	ans := make([][]string, rows)
	for i, _ := range ans {
		ans[i] = make([]string, cols)
	}
	return ans
}

func Fill2D(s2d [][]string, val string) {
	for i, s1d := range s2d {
		for j, _ := range s1d {
			s2d[i][j] = val
		}
	}
}

func PPrint2D(arr [][]string) {
	for i, row := range arr {
		fmt.Printf("%v => \n", i)
		for _, ele := range row {
			fmt.Printf("%v\n", ele)
		}
		fmt.Println()
	}
	fmt.Println()
}
