package main

import (
	"fmt"
)

func absoluteValue(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func getDifference(a int, b int) int {
	return absoluteValue(a - b)
}

func main() {

	// student's scores
	type student struct {
		name string
		chinese int
		math int
		english int
	}

	var n int
	fmt.Scanln(&n)

	var temp1 student
	var studentList = make([]student, 0)
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp1.name, &temp1.chinese, &temp1.math, &temp1.english)
		studentList = append(studentList, temp1)
	}

	var totalgetDifference int
	type opponent struct {
		code1 int
		code2 int
	}
	var temp2 opponent
	var codeList = make([]opponent, 0)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			totalgetDifference = getDifference((studentList[i].chinese + studentList[i].math + studentList[i].english), (studentList[j].chinese + studentList[j].math + studentList[j].english))
			if absoluteValue(studentList[i].chinese - studentList[j].chinese) <= 5 && absoluteValue(studentList[i].math - studentList[j].math) <= 5 && absoluteValue(studentList[i].english - studentList[j].english) <= 5 && totalgetDifference <= 10 {
				temp2.code1 = i
				temp2.code2 = j
				codeList = append(codeList, temp2)
			}
		}
	}

	for i := 0; i < len(codeList); i++ {
		fmt.Printf("%s %s\n", studentList[codeList[i].code1].name, studentList[codeList[i].code2].name)
	}
}
