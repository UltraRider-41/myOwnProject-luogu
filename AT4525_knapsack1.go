// 0/1 knapsack problem, N means the number of items and W means the sum of items' weight

package main

import (
	"fmt"
	"math"
)

var vs = make([]int, 101) // value array
var ws = make([]int, 101) // weight array
var results = make([][]int, 0)
// recursion, will get TLE
func knapsack1(num int, weight int) int{

	result := 0
	if num == 0 || weight == 0 {
		result = 0
	} else if ws[num] > weight { // no enough space
		result = knapsack1(num-1, weight)
	} else { // compare and return the highest
		tmp1 := knapsack1(num-1, weight)
		tmp2 := knapsack1(num-1, weight-ws[num]) + vs[num]
		result = int(math.Max(float64(tmp1), float64(tmp2)))
	}

	return result
}
// dp + recursion + top-down
func knapsack2(num int, weight int) int {

	result := 0
	for a := 0; a <= num; a++ {
		b := make([]int, weight+1)
		results = append(results, b)
	}
	if results[num][weight] != 0 {
		return results[num][weight]
	}
	if num == 0 || weight == 0 {
		result = 0
	} else if ws[num] > weight {
		result = knapsack2(num-1, weight)
	} else {
		tmp1 := knapsack2(num-1, weight)
		tmp2 := knapsack2(num-1, weight-ws[num]) + vs[num]
		result = int(math.Max(float64(tmp1), float64(tmp2)))
	}
	results[num][weight] = result
	return result
}
// down-top
func knapsack3(num int, weight int) int {

	for a := 0; a <= num; a++ {
		b := make([]int, weight+1)
		results = append(results, b)
	}
	// assign 0 to the 1st row
	for m := 0; m <= num; m++ {
		results[m][0] = 0
	}
	// assign 0 to the 1st column
	for m := 0; m <= weight; m++ {
		results[0][m] = 0
	}
	// fill the table
	for m := 1; m <= num; m++ {
		for n := 1; n <= weight; n++ {
			// no enough space
			if n < ws[m] {
				results[m][n] = results[m-1][n]
			} else {
				// enough space, find and return the highest value
				results[m][n] = int(math.Max(float64(results[m-1][n]), float64(results[m-1][n - ws[m]] + vs[m])))
			}
		}
	}
	return results[num][weight]
}

func main() {

	// enter the total number & weight
	var N,W int
	fmt.Scanln(&N, &W)
	// enter the specefic item's value & weight
	for i := 1; i <= N; i++ {
		fmt.Scanln(&ws[i], &vs[i]) // NOTICE the order of '&ws[i]' and '&vs[i]'
	}

	res := knapsack1(N, W)
	fmt.Printf("%d\n", res)

	res = knapsack2(N, W)
	fmt.Printf("%d\n", res)

	res = knapsack3(N, W)
	fmt.Printf("%d", res)
}
