package main

import (
	"fmt"
)

func bubbleSort(list *[]int) int {
	var numList = *list
	var times = 0
	if numList == nil || len(numList) < 2 {
		return 0
	}
	for i := 0; i < len(*list) - 1; i++ {
		for j := 0; j < len(*list) - 1 - i; j++ {
			if numList[j] > numList[j + 1] {
				numList[j], numList[j + 1] = numList[j + 1], numList[j]
				times++
			}
		}
	}
	return times
}

func main() {

	var N int // N <= 10000
	var res int
	var train = make([]int, 0)
	temp := 0
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&temp)
		train = append(train, temp)
	}

	res = bubbleSort(&train)
	fmt.Printf("%d", res)
}
