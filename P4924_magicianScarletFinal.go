// 1.Optimize the input: fmt.Scanln -> bufio.NewReader(os.Stdin)
// 2.Dynamic allocation of array size to static allocation of array size: make([]xxx, 0)+append -> make([]xxx, constNum)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func finalMagic(matrix *[][]int, x, y, r, z, n int) {

	tempTemp := *matrix
	if z == 0 {
		var subMatrix = make([][]int, 501)
		for i := 0; i < 2*r+1; i++ {
			subMatrix[i] = make([]int, 501)
		}
		for i := 0; i < 2*r+1; i++ {
			for j := 0; j < 2*r+1; j++ {
				subMatrix[i][j] = tempTemp[i+x-r-1][j+y-r-1]
			}
		}

		for i := 0; i < 2*r+1; i++ {
			for j := 0; j < 2*r+1; j++ {
				tempTemp[i+x-r-1][j+y-r-1] = subMatrix[2*r+1-j-1][i]
			}
		}

	} else {
		var subMatrix = make([][]int, 501)
		for i := 0; i < 2*r+1; i++ {
			subMatrix[i] = make([]int, 501)
		}
		for i := 0; i < 2*r+1; i++ {
			for j := 0; j < 2*r+1; j++ {
				subMatrix[i][j] = tempTemp[i+x-r-1][j+y-r-1]
			}
		}

		for i := 0; i < 2*r+1; i++ {
			for j := 0; j < 2*r+1; j++ {
				tempTemp[i+x-r-1][j+y-r-1] = subMatrix[j][2*r+1-i-1]
			}
		}
	}
}

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	var matrix = make([][]int, 501)
	num := 1
	for i := 0; i < n; i++ {
		var tempMatrix = make([]int, 501)
		for j := 0; j < n; j++ {
			tempMatrix[j] = num
			num++
		}
		matrix[i] = tempMatrix
	}
	// fmt.Println(matrix)

	var x, y, r, z int
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {
		data, _, _ := reader.ReadLine()
		var context = make([]string, 0)
		context = strings.Fields(string(data))
		x, _ = strconv.Atoi(context[0])
		y, _ = strconv.Atoi(context[1])
		r, _ = strconv.Atoi(context[2])
		z, _ = strconv.Atoi(context[3])
		finalMagic(&matrix, x, y, r, z, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

}
