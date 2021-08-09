package main

import (
	"fmt"
	"math"
)

func newPrimeDamn(num int) int {
	if num == 1 {
		return 0
	}
	// must sqrt, or this function will lead to "TLE"
	// notice: <= instead of <
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			return 0
		}
	}
	return 1
}

func main()  {

	var n int
	fmt.Scanln(&n)

	for i := 4; i <= n; i += 2 {
		for j := 2; j < i; j++ {
			if newPrimeDamn(i-j) == 1 && newPrimeDamn(j) == 1 {
				fmt.Printf("%d=%d+%d\n", i, j, i-j)
				break
			}
		}
	}
}
