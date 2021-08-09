package main

import (
	"fmt"
)

func absolute(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func difference(a int, b int) int {
	return absolute(a - b)
}

func main() {

	// student's scores
	type student struct {
		chinese int
		math int
		english int
	}

	var n int
	fmt.Scanln(&n)

	var temp student
	var studentList = make([]student, 0)
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp.chinese, &temp.math, &temp.english)
		studentList = append(studentList, temp)
	}

	var res int
	var totalDifference int
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			totalDifference = difference((studentList[i].chinese + studentList[i].math + studentList[i].english), (studentList[j].chinese + studentList[j].math + studentList[j].english))
			if absolute(studentList[i].chinese - studentList[j].chinese) <= 5 && absolute(studentList[i].math - studentList[j].math) <= 5 && absolute(studentList[i].english - studentList[j].english) <= 5 && totalDifference <= 10 {
				res++
			}
		}
	}
	fmt.Printf("%d", res)
}
