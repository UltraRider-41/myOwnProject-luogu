package main

import (
	"fmt"
)

var N int

type member struct {
	name string
	year int
	month int
	day int
	ymd int
}

// 插入排序
func insertionSort(list *[]member) {
	myList := *list
	for i:= 1; i < N; i++ {
		j := i
		for j > 0 {
			// 普通的插排这里不应该有 = ，因为要满足插排是稳定排序的条件
			// 但在这里加上 = ，是因为要满足题干条件：人物同生日（相同ymd）时后输入的先输出
			if myList[j].ymd <= myList[j-1].ymd {
				myList[j], myList[j-1] = myList[j-1], myList[j]
				j--
			} else {
				break
			}
		}
	}
}

// 注意同生日的情况！！！
func main() {

	fmt.Scanln(&N)

	var group = make([]member, 100)
	var a, b, c int

	for i := 0; i < N; i++ {
		fmt.Scanln(&group[i].name, &a, &b, &c)
		group[i].year = a
		group[i].month = b
		group[i].day = c
		group[i].ymd = a*10000 + b*100 + c
	}

	/*
	    for i := 0; i < N; i++ {
			fmt.Println(group[i])
		}
	*/

	insertionSort(&group)

	for i := 0; i < N; i++ {
		if (group[i].ymd != 0) && i == N-1 {
			fmt.Printf("%s", group[i].name)
		} else if (group[i].ymd != 0) && i < N-1 {
			fmt.Printf("%s\n", group[i].name)
		}
	}

}
