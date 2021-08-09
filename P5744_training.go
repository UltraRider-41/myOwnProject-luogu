package main

import "fmt"

type student5744 struct {
	name string
	age int
	score int
}

func incScore(a student5744) int {
	res := int(float64(a.score) * 1.2)
	if res >= 600 {
		return 600
	} else {
		return res
	}
}

func main() {

	var n int
	fmt.Scanln(&n)

	var temp student5744
	var studentList = make([]student5744, 0)
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp.name, &temp.age, &temp.score)
		temp.score = incScore(temp)
		temp.age++
		studentList = append(studentList, temp)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", studentList[i].name, studentList[i].age, studentList[i].score)
	}
}
