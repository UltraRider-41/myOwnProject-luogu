package main

import "fmt"

// RemoveRepByLoop 通过两重循环过滤重复元素
func RemoveRepByLoop(slc *[]int) []int {
	var result []int // 存放结果
	list := *slc
	for i := range list{
		flag := true
		for j := range result{
			if list[i] == result[j] {
				flag = false  // 存在重复元素，标识为false
				break
			}
		}
		if flag && list[i] != 0 {  // 标识为false，不添加进结果
			result = append(result, list[i])
		}
	}
	return result
}

// quicksort 快排
func quicksort(myList []int, l int, r int) {

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
		quicksort(myList, l, i-1)
		quicksort(myList, i+1, r)
	}
}

func main() {
	var N int
	fmt.Scanln( &N)

	var list = make([]int, 100)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &list[i])
	}

	// fmt.Println(list)
	list = RemoveRepByLoop(&list)
	// fmt.Println(list)

	quicksort(list, 0, len(list)-1)

	fmt.Printf("%d\n", len(list))
	for i := 0; i < len(list); i++ {
		if i < len(list)-1 {
			fmt.Printf("%d ",list[i])
		} else {
			fmt.Printf("%d", list[i])
		}
	}

}
