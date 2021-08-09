package main

import "fmt"

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	type person struct {
		direction int // 0 means in, 1 means out
		position  string
	}

	var temp1 person
	var list = make([]person, 0)

	for i := 0; i < n; i++ {
		fmt.Scanln(&temp1.direction, &temp1.position)
		list = append(list, temp1)
	}
	// fmt.Println(list)

	type solution struct {
		a int // 0 means left, 1 means right
		s int // s means difference number
	}

	var temp2 solution
	var operation = make([]solution, 0)
	for i := 0; i < m; i++ {
		fmt.Scanln(&temp2.a, &temp2.s)
		operation = append(operation, temp2)
	}
	// fmt.Println(operation)

}
