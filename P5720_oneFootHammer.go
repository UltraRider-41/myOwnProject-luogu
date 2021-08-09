package main

import "fmt"

func main() {

	var length, i int
	fmt.Scanln(&length)

	for i = 1; ; i++ {
		if length == 1 {
			fmt.Printf("%d", i)
			break
		}
		length = length / 2
	}
}
