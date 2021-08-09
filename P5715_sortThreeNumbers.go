package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	if a <= b && b <= c {
		fmt.Printf("%d %d %d", a, b, c)
	} else if a <= c && c <= b {
		fmt.Printf("%d %d %d", a, c, b)
	} else if b <= a && a <= c {
		fmt.Printf("%d %d %d", b, a, c)
	} else if b <= c && c <= a {
		fmt.Printf("%d %d %d", b, c, a)
	} else if c <= a && a <= b {
		fmt.Printf("%d %d %d", c, a, b)
	} else if c <= b && b <= a {
		fmt.Printf("%d %d %d", c, b, a)
	}

}
