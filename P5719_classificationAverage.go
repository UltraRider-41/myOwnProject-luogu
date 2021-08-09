package main

import "fmt"

func main() {

	var n, k int
	fmt.Scanln(&n, &k)

	var aTotal, bTotal int
	var aCount, bCount int
	var aAverage, bAverage float64
	for i := 1; i <= n; i++ {
		if i % k == 0 {
			aTotal += i
			aCount ++
		} else {
			bTotal += i
			bCount++
		}
	}
	aAverage = float64(aTotal) / float64(aCount)
	bAverage = float64(bTotal) / float64(bCount)
	fmt.Printf("%.1f %.1f", aAverage, bAverage)
}
