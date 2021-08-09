package main

import "fmt"

func main() {

	// n <= 100
	var n int
	fmt.Scanln(&n)

	var list = make([]int, 0)
	// 0 <= temp <= 1000
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		list = append(list, temp)
	}

	min := 1001
	for i := 0; i < n; i++ {
		if min > list[i] {
			min = list[i]
		}
	}
	fmt.Printf("%d", min)

}
