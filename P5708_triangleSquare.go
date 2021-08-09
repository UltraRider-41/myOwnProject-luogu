package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, p, s float64 // 0 <= a, b, c <= 1000
	fmt.Scanln(&a, &b, &c)
	p = (a + b + c) / 2
	s = math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Printf("%.1f", s)

}
