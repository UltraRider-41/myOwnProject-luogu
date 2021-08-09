package main

import "fmt"

func getRow(a int, n int) int {
	return a/n + 1
}

func getCol(a int, n int) int {
	return a%n + 1
}

func main() {

	var n int
	fmt.Scanln(&n)

	var square = make([]int, n*n)

	for i := 1; i <= n*n; i++ {
		if i == 1 {
			square[n/2] = 1
		} else {
			for j := 0; j < n*n; j++ {
				if square[j] == i - 1 {
					row := getRow(j, n)
					col := getCol(j, n)
					if row == 1 && col != n {
						row = n
						col += 1
					} else if row != 1 && col == n {
						row -= 1
						col = 1
					} else if row == 1 && col == n {
						row += 1
						col += 0
					} else if row != 1 && col != n {
						if square[n*(row-1-1) + (col+1-1)] == 0 {
							row -= 1
							col += 1
						} else {
							row += 1
						}
					}
					square[n*(row-1) + (col-1)] = i
				}
			}
		}
	}

	// print the magic square
	for i := 0; i < n * n; i++ {
		if (i+1) % n != 0 {
			fmt.Printf("%d ", square[i])
		} else {
			fmt.Printf("%d\n", square[i])
		}
	}
}
