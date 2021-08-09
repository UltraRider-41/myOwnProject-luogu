package main

import "fmt"

func aveScore(a []int) float64 {

	max := 0
	min := 11
	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
		sum += a[i]
	}

	var res float64
	res = float64(sum-max-min) / float64(len(a)-2)

	return res
}

func main()  {

	var n, m int
	fmt.Scanln(&n, &m)

	var score = make([][]int, 0)
	var temp int
	for i := 0; i < n; i++ {
		var tempList = make([]int, 0)
		for j := 0; j < m; j++ {
			fmt.Scan(&temp)
			tempList = append(tempList, temp)
		}
		score = append(score, tempList)
	}
	// fmt.Println(score)

	max := 0.0
	for i := 0; i < n; i++ {
		ave := aveScore(score[i])
		if ave > max {
			max = ave
		}
	}
	fmt.Printf("%.2f", max)
}
