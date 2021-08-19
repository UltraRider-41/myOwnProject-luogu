// one-dimensional array

package main

import "fmt"

func magicPower(matrix *[]int, x, y, r, z, n int) {

	tempTemp := *matrix
	l := 2*r + 1
	var subMatrix = make([]int, l*l)
	for i := 0; i < l*l; i++ {
		row := i / l
		col := i % l
		subMatrix[i] = tempTemp[(row+x-r-1)*n+(col+y-r-1)]
	}
	for i := 0; i < l*l; i++ {
		row := i / l
		col := i % l
		if z == 0 {
			tempTemp[(row+x-r-1)*n+(col+y-r-1)] = subMatrix[(2*r+1-col-1)*l+row]
		} else {
			tempTemp[(row+x-r-1)*n+(col+y-r-1)] = subMatrix[col*l+(2*r+1-row-1)]
		}
	}
}

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	var matrix = make([]int, 0)
	for i := 0; i < n*n; i++ {
		matrix = append(matrix, i+1)
	}
	// fmt.Println(matrix)

	var x, y, r, z int
	for i := 0; i < m; i++ {
		fmt.Scanln(&x, &y, &r, &z)
		magicPower(&matrix, x, y, r, z, n)
	}

	for i := 0; i < n*n; i++ {
		if (i+1)%n != 0 {
			fmt.Printf("%d ", matrix[i])
		} else {
			fmt.Printf("%d\n", matrix[i])
		}
	}

}
