package main

import "fmt"

func main() {

	var n, temp, res int
	fmt.Scanln(&n)

	var list = make([]int, 0)
	var min, max int
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		list = append(list, temp)
	}

	min = 1001
	max = -1
	for i := 0; i < n; i++ {
		if list[i] > max {
			max = list[i]
		}
		if list[i] < min {
			min = list[i]
		}
	}

	res = max - min
	fmt.Printf("%d", res)
}