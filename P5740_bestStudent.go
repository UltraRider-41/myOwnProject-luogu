package main

import "fmt"

type student struct {
	name string
	chinese int
	math int
	english int
	total int
}

func sumScore(a student) int {
	res := a.chinese + a.math + a.english
	return res
}

func main() {

	 var n int
	 fmt.Scanln(&n)

	var list = make([]student, 0)
	var temp student
	for i := 0; i < n; i++ {
		fmt.Scanln(&temp.name, &temp.chinese, &temp.math, &temp.english)
		temp.total = sumScore(temp)
		list = append(list, temp)
	}
	// fmt.Println(list)

	max := -1
	var code int
	for i := 0; i < len(list); i++ {
		if list[i].total > max {
			max = list[i].total
			code = i
		}
	}

	fmt.Printf("%s %d %d %d", list[code].name, list[code].chinese, list[code].math, list[code].english)
}
