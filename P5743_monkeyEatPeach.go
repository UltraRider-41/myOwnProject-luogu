package main

import "fmt"

func eat(day int) int {
	if day == 1 {
		return 1
	} else {
		return (eat(day-1)+1)*2
	}
}

func main() {

	var n int
	fmt.Scanln(&n)

	fmt.Println(eat(n))
}
