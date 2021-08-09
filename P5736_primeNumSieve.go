package main

import (

	"fmt"
	"math"
)

func prime(num int) bool {
	if num == 1 {
		return false
	}
	// must sqrt, or this function will lead to "TLE"
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func main()  {

	var n int
	fmt.Scanln(&n)

	var temp int
	var list = make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		list = append(list, temp)
	}
	// fmt.Printf("%T %#v\n", list[0], list[0])

	for i := 0; i < n; i++ {
		if prime(list[i]) {
			fmt.Printf("%d ", list[i])
		}
	}
}
