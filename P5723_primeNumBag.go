package main

import (
	"fmt"
	"math"
)

func isPrimeNum(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {

	var L int
	fmt.Scanln(&L)

	var j, res int
	for i := 2; ; i++ {
		if isPrimeNum(i) {
			if L - res >= i {
				res += i
				fmt.Println(i)
				j++
			} else {
				fmt.Printf("%d", j)
				break
			}
		}
	}
}