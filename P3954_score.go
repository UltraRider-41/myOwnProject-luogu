package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	var res int
	res = int(float64(a) * 0.2 + float64(b) * 0.3 + float64(c) * 0.5)
	fmt.Printf("%d", res)
}
