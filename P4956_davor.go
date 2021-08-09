package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	var x, k int
	for x = 100; x >= 1; x-- {
		k = ((n / 52) - 7 * x) / 21
		if k > 0 && 52 * (7 * x + 21 * k) == n {
			fmt.Printf("%d\n%d", x, k)
			break
		}
	}
}
