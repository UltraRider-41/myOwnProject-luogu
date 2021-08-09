package main

import "fmt"

// bubble sort
func mySimpleSort(list *[]int) {
	var numList = *list
	if numList == nil || len(numList) < 2 {
		return
	}
	for i := 0; i < len(*list) - 1; i++ {
		for j := 0; j < len(*list) - 1 - i; j++ {
			if numList[j] > numList[j + 1] {
				numList[j], numList[j + 1] = numList[j + 1], numList[j]
			}
		}
	}
}

func isTriangle(a int, b int, c int) bool {

	if a + b > c && a + c > b && b + c > a {
		return true
	} else {
		return false
	}
}

func isRight(a int, b int, c int) bool {
	if a * a + b * b == c * c {
		return true
	} else {
		return false
	}
}

func isAcute(a int, b int, c int) bool {
	if a * a + b * b > c * c {
		return true
	} else {
		return false
	}
}

func isObtuse(a int, b int, c int) bool {
	if a * a + b * b < c * c {
		return true
	} else {
		return false
	}
}

func isIsosceles(a int, b int, c int) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func isEquilateral(a int, b int, c int) bool {
	if a == b && a == c {
		return true
	} else {
		return false
	}
}

func main() {

	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	var list = []int{a, b, c}
	mySimpleSort(&list)
	a = list[0]
	b = list[1]
	c = list[2]
	if isTriangle(a, b, c) {
		if isRight(a, b, c) {
			fmt.Println("Right triangle")
		}
		if isAcute(a, b, c) {
			fmt.Println("Acute triangle")
		}
		if isObtuse(a, b, c) {
			fmt.Println("Obtuse triangle")
		}
		if isIsosceles(a, b, c) {
			fmt.Println("Isosceles triangle")
		}
		if isEquilateral(a, b, c) {
			fmt.Println("Equilateral triangle")
		}
	} else {
		fmt.Printf("Not triangle")
	}
}
