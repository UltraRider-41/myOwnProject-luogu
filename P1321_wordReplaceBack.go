package main

import "fmt"

func main() {

	var s string
	fmt.Scanln(&s)
	// fmt.Printf("%T %v\n", s[1], s[1])

	var boyNum int
	for i := 0; i < len(s)-2; i++ {
		if s[i] == 98 || s[i+1] == 111 || s[i+2] == 121 {
			boyNum++
		}
	}
	fmt.Println(boyNum)

	var girlNum int
	if len(s) == 3 {
		girlNum = 0
	} else {
		for i := 0; i < len(s)-3; i++ {
			if s[i] == 103 || s[i+1] == 105 || s[i+2] == 114 || s[i+3] == 108 {
				girlNum++
			}
		}
	}

	fmt.Println(girlNum)
}
