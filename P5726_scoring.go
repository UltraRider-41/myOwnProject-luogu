package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	var temp int
	var totalScore int
	var scoreList = make([]int, 0)
	min := 10
	max := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		scoreList = append(scoreList, temp)
		totalScore += temp
		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp
		}
	}

	var res float64
	res = float64(totalScore - min - max) / float64(n - 2)
	fmt.Printf("%.2f", res)
}
