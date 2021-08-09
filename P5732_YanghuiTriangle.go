package main

import "fmt"

func main() {

	var n int // 1 <= n <= 20
	fmt.Scanln(&n)

	var triangle [][]int
	for x := 0; x < n; x++ {  //循环为一维长度
		arr := make([]int, n) //创建一个一维切片
		triangle = append(triangle, arr) //把一维切片,当作一个整体传入二维切片中
	}

	// var row, col int
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if j == i || j == 0 {
				triangle[i][j] = 1
			}
			if i > 0 && j > 0 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}

	for x := 0; x < n; x++ {
		for y := 0; y <= x; y++ {
			if y != x {
				fmt.Printf("%d ", triangle[x][y])
			} else {
				fmt.Printf("%d\n", triangle[x][y])
			}
		}
	}
}
