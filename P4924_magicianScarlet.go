// two-dimensional array

package main

import "fmt"

func magic(matrix *[][]int, x, y, r, z int) {

	tempTemp := *matrix

	var subMatrix = make([][]int, 2*r+1)
	for i := 0; i < 2*r+1; i++ {
		subMatrix[i] = make([]int, 2*r+1)
	}
	for i := 0; i < 2*r+1; i++ {
		for j := 0; j < 2*r+1; j++ {
			subMatrix[i][j] = tempTemp[i+x-r-1][j+y-r-1]
		}
	}
	for i := 0; i < 2*r+1; i++ {
		for j := 0; j < 2*r+1; j++ {
			if z == 0 {
				tempTemp[i+x-r-1][j+y-r-1] = subMatrix[2*r+1-j-1][i]
			} else {
				tempTemp[i+x-r-1][j+y-r-1] = subMatrix[j][2*r+1-i-1]
			}
		}
	}
}

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	var matrix = make([][]int, 0)
	num := 1
	for i := 0; i < n; i++ {
		var tempMatrix = make([]int, n)
		for j := 0; j < n; j++ {
			tempMatrix[j] = num
			num++
		}
		matrix = append(matrix, tempMatrix)
	}
	// fmt.Println(matrix)

	var x, y, r, z int
	for i := 0; i < m; i++ {
		fmt.Scanln(&x, &y, &r, &z)
		magic(&matrix, x, y, r, z)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

}
