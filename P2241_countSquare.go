package main

import "fmt"

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	var min, max int
	if n < m {
		min = n
		max = m
	} else {
		min = m
		max = n
	}

	var rectangle, square int
	// count square
	var squareL, squareW int
	for r := 1; r <= min; r++ {
		squareL = max - r + 1
		squareW = min - r + 1
		square += squareL * squareW
	}

	// count rectangle
	var recL, recW int
	for l := 1; l <= max; l++ {
		recL = max - l + 1
		for w := 1; w <= min; w++ {
			recW = min - w + 1
			if l != w {
				rectangle += recL * recW
			}
		}
	}

	fmt.Println(square, rectangle)
}
