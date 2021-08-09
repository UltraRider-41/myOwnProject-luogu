package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	var res = make([]int, a*b*c)
	for x := 1; x <= a; x++ {
		for y := 1; y <= b; y++ {
			for z := 1; z <= c; z++ {
				res[(x+y+z)-1] += 1
			}
		}
	}

	var max = -1
	var j = 0
	for i := 0; i < a*b*c; i++ {
		if max < res[i] {
			max = res[i]
			j = i
		}
	}
	fmt.Printf("%d", j+1)

}
