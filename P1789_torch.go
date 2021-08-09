package main

import "fmt"

func main() {

	var n, m, k int
	fmt.Scanln(&n, &m, &k)

	var field [][]int
	for x := 0; x < n; x++ {  //循环为一维长度
		arr := make([]int, n) //创建一个一维切片
		field = append(field, arr) //把一维切片,当作一个整体传入二维切片中
	}

	type point struct {
		x int
		y int
	}

	var torch = make([]point, 0)
	var fluorite = make([]point, 0)
	var temp point
	for i := 0; i < m; i++ {
		fmt.Scanln(&temp.x, &temp.y)
		torch = append(torch, temp)
	}
	for j := 0; j < k; j++ {
		fmt.Scanln(&temp.x, &temp.y)
		fluorite = append(fluorite, temp)
	}

	// torch
	// point in field -> torch's point
	var row, col int
	for i := 0; i < m; i++ {
		row = torch[i].x
		col = torch[i].y
		for a := 0; a < n; a++ {
			for b := 0; b < n; b++ {
				if (row-a-1)*(row-a-1) + (col-b-1)*(col-b-1) <= 4 {
					field[a][b] = 1
				}
			}
		}
	}
	/*
	"torch's field -> point in field" too complex, failed
	var row, col int
	for i := 0; i < m; i++ {
		row = torch[i].x
		col = torch[i].y
		for a := -2; a <= 2; a++ {
			if row+a-1 <= n-1 && row+a-1 >= 0 {
				for b := -2; b <= 2; b++ {
					if col+b-1 <= n-1 && col+b-1 >= 0 {
						if a == -2 || a == 2 {
							field[row+a-1][col+b-1] = 1
						} else if a == -1 || a == 1 {
							field[row+a-1][col+b-1] = 1

						} else if a == 0 {
							field[row+a-1][col+b-1] = 1

						}
					}
				}
			}
		}
	}
	 */

	// fluorite
	// fluorite's field -> point in field
	for j := 0; j < k; j++ {
		row = fluorite[j].x
		col = fluorite[j].y
		for a := -2; a <= 2; a++ {
			if row+a-1 <= n-1 && row+a-1 >= 0 {
				for b := -2; b <= 2; b++ {
					if col+b-1 <= n-1 && col+b-1 >= 0 {
						field[row+a-1][col+b-1] = 1
					}
				}
			}
		}
	}

	var monsterNum int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if field[i][j] == 0 {
				monsterNum++
			}
		}
	}
	fmt.Printf("%d", monsterNum)
}
