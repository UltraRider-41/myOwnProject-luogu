package main

import "fmt"

func leap(year int) bool {
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	var year, month, days int
	fmt.Scanln(&year, &month)

	switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			days = 31
		case 4, 6 ,9 ,11:
			days = 30
		case 2:
			if leap(year) {
				days = 29
			} else {
				days = 28
			}
	}
	fmt.Printf("%d", days)
}
