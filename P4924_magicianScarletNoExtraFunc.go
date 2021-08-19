// one-dimensional array + no "func magic()"

package main

import "fmt"

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
		l := 2*r + 1
		var subMatrix = make([]int, l*l)
		for i := 0; i < l*l; i++ {
			row := i / l
			col := i % l
			subMatrix[i] = matrix[(row+x-r-1)*n+(col+y-r-1)]
		}
		for i := 0; i < l*l; i++ {
			row := i / l
			col := i % l
			if z == 0 {
				matrix[(row+x-r-1)*n+(col+y-r-1)] = subMatrix[(2*r+1-col-1)*l+row]
			} else {
				matrix[(row+x-r-1)*n+(col+y-r-1)] = subMatrix[col*l+(2*r+1-row-1)]
			}
		}
	}

	for i := 0; i < n*n; i++ {
		if (i+1)%n != 0 {
			fmt.Printf("%d ", matrix[i])
		} else {
			fmt.Printf("%d\n", matrix[i])
		}
	}
}
