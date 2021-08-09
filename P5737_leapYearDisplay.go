package main

import "fmt"

func isLeapYear(year int) int {
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {

	var num int
	var x, y int
	fmt.Scanln(&x, &y)

	for i := x; i <= y; i++ {
		if isLeapYear(i) == 1 {
			num++
		}
	}
	fmt.Println(num)
	for i := x; i <= y; i++ {
		if isLeapYear(i) == 1 {
			fmt.Printf("%d ", i)
		}
	}
}
