package main

import "fmt"

var n int

// judge if a slice equals b slice
func equal(a [][]string, b [][]string) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// judge the situation
func judge(a [][]string, b [][]string) int {
	if equal(clockwise90(a), b) {
		return 1
	} else if equal(clockwise180(a), b) {
		return 2
	} else if equal(clockwise270(a), b) {
		return 3
	} else if equal(horizonal(a), b) {
		return 4
	} else if equal(combination90(a), b) || equal(combination180(a), b) || equal(combination270(a), b) {
		return 5
	} else if equal(a, b) {
		return 6
	} else {
		return 7
	}
}

// 1
func clockwise90(matrix [][]string) [][]string {
	var temp = make([][]string, len(matrix))
	for i := range matrix {
		temp[i] = make([]string, len(matrix[i]))
		// copy(temp[i], matrix[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp[j][n-i-1] = matrix[i][j]
		}
	}
	return temp
}

// 2
func clockwise180(matrix [][]string) [][]string {
	var temp = make([][]string, len(matrix))
	for i := range matrix {
		temp[i] = make([]string, len(matrix[i]))
		// copy(temp[i], matrix[i])
	}
	temp = clockwise90(matrix)
	temp = clockwise90(temp)
	return temp
}

// 3
func clockwise270(matrix [][]string) [][]string {
	var temp = make([][]string, len(matrix))
	for i := range matrix {
		temp[i] = make([]string, len(matrix[i]))
		// copy(temp[i], matrix[i])
	}
	temp = clockwise180(matrix)
	temp = clockwise90(temp)
	return temp
}

// 4
func horizonal(matrix [][]string) [][]string {
	var temp = make([][]string, len(matrix))
	for i := range matrix {
		temp[i] = make([]string, len(matrix[i]))
		// copy(temp[i], matrix[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp[i][j] = matrix[i][n-j-1]
		}
	}
	return temp
}

// 5.90
func combination90(matrix [][]string) [][]string {
	var temp = make([][]string, len(matrix))
	for i := range matrix {
		temp[i] = make([]string, len(matrix[i]))
		// copy(temp[i], matrix[i])
	}
	temp = horizonal(matrix)
	temp = clockwise90(temp)
	return temp
}
// 5.180
func combination180(matrix [][]string) [][]string {
	var temp = make([][]string, len(matrix))
	for i := range matrix {
		temp[i] = make([]string, len(matrix[i]))
		// copy(temp[i], matrix[i])
	}
	temp = horizonal(matrix)
	temp = clockwise180(temp)
	return temp
}
// 5.270
func combination270(matrix [][]string) [][]string {
	var temp = make([][]string, len(matrix))
	for i := range matrix {
		temp[i] = make([]string, len(matrix[i]))
		// copy(temp[i], matrix[i])
	}
	temp = horizonal(matrix)
	temp = clockwise270(temp)
	return temp
}

func main() {

	// handle inputs
	fmt.Scanln(&n)

	var primeBlock = make([][]string, 0)
	for i := 0; i < n; i++ {
		var arr = make([]string, n)
		primeBlock = append(primeBlock, arr)
	}
	var temp1 string
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp1)
		for j := 0; j < n; j++ {
			primeBlock[i][j] = string(temp1[j])
		}
	}

	var resBlock = make([][]string, 0)
	for i := 0; i < n; i++ {
		var ar = make([]string, n)
		resBlock = append(resBlock, ar)
	}
	var temp2 string
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp2)
		for j := 0; j < n; j++ {
			resBlock[i][j] = string(temp2[j])
		}
	}

	// fmt.Printf("%T %v\n", primeBlock, primeBlock)
	// fmt.Printf("%T %v\n", primeBlock[0][0], primeBlock[0][0])

	result := judge(primeBlock, resBlock)
	fmt.Printf("%d", result)
}
