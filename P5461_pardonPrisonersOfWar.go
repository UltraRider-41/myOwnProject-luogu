package main

import (
	"fmt"
	"math"
)

// 递归条件有问题，考虑是否加上返回值
func pardon(matrix *[][]int, length int, rStart int, cStart int) {

	m := *matrix
	l := length
	newL := l/2

	if l == 1 {
		m[rStart][cStart] = 1
		return
	} else {
		for i := 0; i < newL; i++ {
			for j := 0; j < newL; j++ {
				m[rStart+i][cStart+j] = 0
			}
		}
		pardon(matrix, newL, rStart+0, cStart+newL) // right-up
		pardon(matrix, newL, rStart+newL, cStart+0) // left-down
		pardon(matrix, newL, rStart+newL, cStart+newL) // right-down
	}
	return
}

func main() {

	// 0 <= n <= 10
	var n int
	fmt.Scanln(&n)

	length := int(math.Pow(2.0, float64(n)))
	var matrix = make([][]int, 0)
	for i := 0; i <length; i++ {
		var temp = make([]int, length)
		for j := 0; j < length; j++ {
			temp[j] = 1
		}
		matrix = append(matrix, temp)
	}
	// fmt.Println(matrix)
	pardon(&matrix, length, 0, 0)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (j+1) % length == 0 {
				fmt.Printf("%d\n", matrix[i][j])
			} else {
				fmt.Printf("%d ", matrix[i][j])
			}
		}
	}
}
