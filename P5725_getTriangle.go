package main

import "fmt"

func totalNum(a int) int {

	var res int
	for i := 1; i <= a; i++ {
		res += i
	}
	return res
}

func main() {

	var n int
	fmt.Scanln(&n)

	var maxNum int
	maxNum = totalNum(n)

	// print square
	for i := 1; i <= n * n;  {
		if i % n != 0 {
			if i < 10 {
				fmt.Printf("0%d", i)
			} else {
				fmt.Printf("%2d", i)
			}
		} else {
			if i < 10 {
				fmt.Printf("0%d\n", i)
			} else {
				fmt.Printf("%2d\n", i)
			}
		}
		i++
	}
	fmt.Println()

	var temp = 1
	// print triangle
	for i := 1; i <= maxNum;  {
		for j := 1; j <= n; j++ {
			if j <= n - temp {
				fmt.Printf("  ")
			} else if j > n - temp {
				if i < 10 {
					fmt.Printf("0%d", i)
				} else {
					fmt.Printf("%2d", i)
				}
				i++
			}
		}
		temp++
		fmt.Printf("\n")
	}
}