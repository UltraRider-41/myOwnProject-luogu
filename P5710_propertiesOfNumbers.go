package main

import "fmt"

func isEven(a int) int {
	if a % 2 == 0 {
		return 1
	} else {
		return 0
	}
}

func numCompare(a int) int {
	if a > 4 && a <= 12 {
		return 1
	} else {
		return 0
	}
}

func main() {

	var x int
	fmt.Scanln(&x)
	m := isEven(x)
	n := numCompare(x)
	var i int
	i = m + n
	switch i {
		case 0:
			fmt.Printf("0 0 0 1")
		case 1:
			fmt.Printf("0 1 1 0")
		case 2:
			fmt.Printf("1 1 0 0")
	}
}
