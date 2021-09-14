package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getKeys(m map[int]int) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func myQuickSort(myList []int, l int, r int) {

	if l < r {
		i := l
		j := r
		pivot := myList[i]

		for i < j {
			for i < j && myList[j] > pivot {
				j--
			}
			if i < j {
				myList[i] = myList[j]
				i++
			}
			for i < j && myList[i] < pivot {
				i++
			}
			if i < j {
				myList[j] = myList[i]
				j--
			}
		}
		myList[i] = pivot
		myQuickSort(myList, l, i-1)
		myQuickSort(myList, i+1, r)
	}
}

func main() {

	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	var list = make([]int, 0)
	for i := 0; i < n; i++ {
		context, _ := reader.ReadString('\n')
		data := strings.Fields(context)
		temp, _ := strconv.Atoi(data[0])
		list = append(list, temp)
	}

	var mapList = make(map[int]int)
	for i := 0; i < n; i++ {
		temp := mapList[list[i]]
		mapList[list[i]] = temp + 1
	}

	// fmt.Println(mapList)
	// fmt.Println(getKeys(mapList))
	var finalList = make([]int, 0)
	keys := getKeys(mapList)
	for i := 0; i < len(keys); i++ {
		temp := keys[i]
		finalList = append(finalList, temp)
	}
	l := 0
	r := len(keys)
	myQuickSort(finalList, l, r-1)
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%d %d\n", finalList[i], mapList[finalList[i]])
	}

}
