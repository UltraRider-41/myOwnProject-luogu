package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	if n <= 1 {
		fmt.Printf("Today, I ate %d apple.", n)
	} else if n > 1 {
		fmt.Printf("Today, I ate %d apples.", n)
	}
}
