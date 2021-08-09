// ks means knapsack
// rslice means result slice

package main

import (
	"fmt"
	"math"
)

var vslice = []int{0, 6, 5, 6, 6, 5, 7}
var wslice = []int{0, 5, 6, 4, 6, 5, 2}
// Be careful of the difference of result and rslice!
var rslice [7][16]int

//
func ks1(num int, weight int) int{

	result := 0
	if num == 0 || weight == 0 {
		result = 0
	} else if wslice[num] > weight {
		result = ks1(num-1, weight)
	} else {
		tmp1 := ks1(num-1, weight)
		tmp2 := ks1(num-1, weight-wslice[num]) + vslice[num]
		result = int(math.Max(float64(tmp1), float64(tmp2)))
	}
	return result
}

//
func ks2(num int, weight int) int{

	result := 0
	if rslice[num][weight] != 0 {
		return rslice[num][weight]
	}
	if num == 0 || weight == 0 {
		result = 0
	} else if wslice[num] > weight {
		result = ks2(num-1, weight)
	} else {
		tmp1 := ks2(num-1, weight)
		tmp2 := ks2(num-1, weight-wslice[num]) + vslice[num]
		result = int(math.Max(float64(tmp1), float64(tmp2)))
	}
	rslice[num][weight] = result
	return result
}

//
func ks3(num int, weight int) int {

	// assign 0 to the 1st row
	for m := 0; m <= num; m++ {
		rslice[m][0] = 0
	}
	// assign 0 to the 1st column
	for m := 0; m <= weight; m++ {
		rslice[0][m] = 0
	}
	// fill the table
	for m := 1; m <= num; m++ {
		for n := 1; n <= weight; n++ {
			// not enough space
			if n < wslice[m] {
				rslice[m][n] = rslice[m-1][n]
			} else {
				// enough space, find and return the highest value
				rslice[m][n] = int(math.Max(float64(rslice[m-1][n]), float64(rslice[m-1][n - wslice[m]] + vslice[m])))
			}
		}
	}
	return rslice[num][weight]
}

func main() {
	N := 6
	W := 15
	result := ks1(N, W)
	fmt.Println("ks1:", result)
	result = ks2(N, W)
	fmt.Println("ks2:", result)
	result = ks3(N, W)
	fmt.Println("ks3:", result)
}
