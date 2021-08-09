package main

import "fmt"

func qsort(myList []int, l int, r int) {

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
		qsort(myList, l, i-1)
		qsort(myList, i+1, r)
	}
}


func main() {
	// 输入数量N
	var N int
	fmt.Scanln(&N)

	// 输入N个数字于数组中
	var numList = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &numList[i])
	}

	// 对数组进行快排
	l := 0
	r := len(numList)
	qsort(numList, l, r-1)

	// 输出排序后的数组
	for i, v := range numList {
		if i < N-1 {
			fmt.Printf("%d ", v)
		}else {
			fmt.Printf("%d\n", v)
		}
	}
}
