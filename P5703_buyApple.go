package main

import "fmt"

func main() {
	var a, b, res int
	fmt.Scanln(&a, &b)
	res = a * b
	fmt.Printf("%d", res)
}
