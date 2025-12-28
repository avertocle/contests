package bytez

import "fmt"

func Init3D(rows, cols, depth int, b byte) [][][]byte {
	ans := make([][][]byte, rows)
	for i := 0; i < rows; i++ {
		ans[i] = Init2D(cols, depth, b)
	}
	return ans
}

func Count3D(arr [][][]byte, v byte) int {
	ctr := 0
	for _, a2d := range arr {
		ctr += Count2D(a2d, v)
	}
	return ctr
}

func PPrint3D(arr [][][]byte) {
	for i, a2d := range arr {
		fmt.Printf("%02v => ", i)
		for _, row := range a2d {
			fmt.Printf("[%v] ", string(row))
		}
		fmt.Println()
	}
	fmt.Println()
}
