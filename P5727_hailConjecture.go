package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	var list = make([]int, 0)
	list = append(list, n)
	for i := 0; ; i++ {
		if n == 1 { // notice: need to determine if n == 1 first
			break
		}
		if n % 2 != 0 {
			n = 3 * n + 1
		} else {
			n /= 2
		}
		list = append(list, n)
	}
	for i := len(list) - 1; i >= 0; i-- {
		fmt.Printf("%d ", list[i])
	}

}