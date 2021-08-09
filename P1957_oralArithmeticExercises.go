package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var exercise = make([][]string, 0)
	for i := 0; i < n; i++ {
		var temp = make([]string, 3)
		var x, y, z string
		fmt.Scanln(&x, &y, &z)
		temp[0] = x
		temp[1] = y
		temp[2] = z
		exercise = append(exercise, temp)
	}
	// fmt.Printf("%#v", exercise[0][0])

	var flag string
	flag = exercise[0][0]
	for i := 0; i < n; i++ {
		switch exercise[i][0] {
		case "a":
			flag = "a"
			num1, _ := strconv.Atoi(exercise[i][1])
			num2, _ := strconv.Atoi(exercise[i][2])
			res := num1 + num2
			s := fmt.Sprintf("%d+%d=%d", num1, num2, res)
			fmt.Printf("%s\n%d\n", s, len(s))
		case "b":
			flag = "b"
			num1, _ := strconv.Atoi(exercise[i][1])
			num2, _ := strconv.Atoi(exercise[i][2])
			res := num1 - num2
			s := fmt.Sprintf("%d-%d=%d", num1, num2, res)
			fmt.Printf("%s\n%d\n", s, len(s))
		case "c":
			flag = "c"
			num1, _ := strconv.Atoi(exercise[i][1])
			num2, _ := strconv.Atoi(exercise[i][2])
			res := num1 * num2
			s := fmt.Sprintf("%d*%d=%d", num1, num2, res)
			fmt.Printf("%s\n%d\n", s, len(s))
		default:
			switch flag {
			case "a":
				num1, _ := strconv.Atoi(exercise[i][0])
				num2, _ := strconv.Atoi(exercise[i][1])
				res := num1 + num2
				s := fmt.Sprintf("%d+%d=%d", num1, num2, res)
				fmt.Printf("%s\n%d\n", s, len(s))
			case "b":
				num1, _ := strconv.Atoi(exercise[i][0])
				num2, _ := strconv.Atoi(exercise[i][1])
				res := num1 - num2
				s := fmt.Sprintf("%d-%d=%d", num1, num2, res)
				fmt.Printf("%s\n%d\n", s, len(s))
			case "c":
				num1, _ := strconv.Atoi(exercise[i][0])
				num2, _ := strconv.Atoi(exercise[i][1])
				res := num1 * num2
				s := fmt.Sprintf("%d*%d=%d", num1, num2, res)
				fmt.Printf("%s\n%d\n", s, len(s))
			}
		}
	}
}