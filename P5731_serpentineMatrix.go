package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var matrix [][]int
	for x := 0; x < n; x++ {  //循环为一维长度
		arr := make([]int, n) //创建一个一维切片
		matrix = append(matrix, arr) //把一维切片,当作一个整体传入二维切片中
	}

	var num = 1
	var row, col int
	var a, b, c, d int
here: for i := 0; i < int(math.Ceil(float64(n)/2.0)); i++ {
		// right
		for a = i; a < n-i-1; a++ {

			row = i
			col = a
			matrix[row][col] = num
			if num == n*n {
				break here
			}
			num++
		}
		// down
		for b = i; b < n-i-1; b++ {

			row = b
			col = n-i-1
			matrix[row][col] = num
			if num == n*n {
				break here
			}
			num++
		}
		// left
		for c = n-i-1; c > i; c-- {
			row = n-i-1
			col = c
			matrix[row][col] = num
			if num == n*n {
				break here
			}
			num++
		}
		// up
		for d = n-i-1; d > i; d-- {

			row = d
			col = i
			matrix[row][col] = num
			if num == n*n {
				break here
			}
			num++
		}
	}
	if n%2 != 0 {
		matrix[n/2][n/2] = n*n
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if (y+1) % n == 0 {
				fmt.Printf(" %2d\n", matrix[x][y])
			} else {
				fmt.Printf(" %2d", matrix[x][y])
			}
		}
	}
}
