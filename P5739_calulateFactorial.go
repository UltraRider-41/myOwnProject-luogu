package main

import "fmt"

func calFac(a int) int {
	var res int
	if a == 1 {
		res = 1
	} else {
		res = a * calFac(a-1)
	}

	return res
}

func main()  {

	var n int
	fmt.Scanln(&n)

	res := calFac(n)
	fmt.Printf("%d", res)
}
