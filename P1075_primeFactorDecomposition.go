package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
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

func main() {

	// n <= 2 * 10^9
	var n int
	fmt.Scanln(&n)
	j := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			j = n / i
			if isPrime(j) && i * j == n {
				fmt.Printf("%d", j)
				break
			}
		}
	}
}
