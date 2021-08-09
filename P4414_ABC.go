package main

import "fmt"

func myBubbleSort(list *[]int) {
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

func main() {

	var x, y, z int
	var order string
	fmt.Scanln(&x, &y, &z)
	fmt.Scanln(&order)

	var list = []int{x, y, z}
	myBubbleSort(&list)
	x = list[0]
	y = list[1]
	z = list[2]

	switch order {
		case "ABC":
			fmt.Printf("%d %d %d", x, y, z)
		case "ACB":
			fmt.Printf("%d %d %d", x, z, y)
		case "BAC":
			fmt.Printf("%d %d %d", y, x, z)
		case "BCA":
			fmt.Printf("%d %d %d", y, z, x)
		case "CAB":
			fmt.Printf("%d %d %d", z, x, y)
		case "CBA":
			fmt.Printf("%d %d %d", z, y, x)
	}
}
