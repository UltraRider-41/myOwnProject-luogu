package main

import (
	"fmt"
	"sort"
)

func main() {

	var n, b int
	fmt.Scanln(&n, &b)

	var cow = make([]int, 0)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp)
		cow = append(cow, temp)
	}

	min := 0
	sum := 0
	sort.Ints(cow)
	for i := n - 1; i >= 0; i-- {
		sum += cow[i]
		min += 1
		if sum >= b {
			break
		}
	}
	fmt.Println(min)
}
