package main

import "fmt"

func leapYear(year int) int {
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {

	var year int
	fmt.Scanln(&year)

	res := leapYear(year)
	fmt.Printf("%d", res)
}
