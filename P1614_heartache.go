package main

import "fmt"

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	temp := 0
	var numList = make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp)
		numList = append(numList, temp)
	}

	min := 100001
	if m == 0 {
		min = 0
	} else if n == m {
		res := 0
		for a := 0; a < n; a++ {
			res += numList[a]
		}
		min = res
	} else if n > m {
		for a := 0; a <= n-m; a++ {
			res := 0
			for b := 0; b < m; b++ {
				res += numList[a+b]
			}
			if res < min {
				min = res
			}
		}
	}

	fmt.Printf("%d", min)
}
