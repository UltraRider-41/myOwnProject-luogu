package main

import "fmt"

func total(a int) int {

	var res int
	for i := 1; i <= a; i++ {
		res += i
	}
	return res
}

func main() {

	var n int // 1 <= n <= 13
	fmt.Scanln(&n)

	var maxNum int
	maxNum = total(n)

	for i := 1; i <= maxNum;  {
		for j := n; j >= 1; j-- {
			if i < 10 {
				fmt.Printf("0%d", i)
			} else {
				fmt.Printf("%2d", i)
			}
			i++
		}
		n--
		fmt.Printf("\n")
	}
}
