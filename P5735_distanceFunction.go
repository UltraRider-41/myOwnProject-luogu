package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func distance(a point, b point) float64 {
	var res float64
	res = math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
	return res
}

func main() {

	var a, b, c point
	fmt.Scanln(&a.x, &a.y)
	fmt.Scanln(&b.x, &b.y)
	fmt.Scanln(&c.x, &c.y)
	// or like this:
	// fmt.Scan(&a.x, &a.y, &b.x, &b.y, &c.x, &c.y)

	var perimeter float64
	perimeter = distance(a, b) + distance(a, c) + distance(b, c)

	fmt.Printf("%.2f", perimeter)
}
