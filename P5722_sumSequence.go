package main

import "fmt"

func main() {

	var n int // 1 <= n <= 100
	fmt.Scanln(&n)

	var res int
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Printf("%d", res)
}
