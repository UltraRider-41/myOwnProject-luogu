package main

import "fmt"

type thisStudent struct {
	id int
	score1 int // study score
	score2 int // quality score
	total int // total score
}

func (rec thisStudent) getTotalScore(a thisStudent) int {
	res := 0.7*float64(a.score1) + 0.3*float64(a.score2)
	return int(res)
}

func judgeExcellent(a thisStudent) bool {
	if (a.score1+a.score2) > 140 && a.total >= 80 {
		return true
	} else {
		return false
	}
}

func main() {

	var n int
	fmt.Scanln(&n)

	var temp thisStudent
	var studentList = make([]thisStudent, 0)
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp.id, &temp.score1, &temp.score2)
		temp.total = temp.getTotalScore(temp)
		studentList = append(studentList, temp)
	}

	// fmt.Println(studentList)
	for i := 0; i < len(studentList); i++ {
		if judgeExcellent(studentList[i]) {
			fmt.Println("Excellent")
		} else {
			fmt.Println("Not excellent")
		}
	}
}