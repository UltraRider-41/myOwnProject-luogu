package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var s string
	fmt.Scanln(&s)

	var some int
	var res = make([]string, 0)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &some)
		switch some {
		case 1:
			var aStr string
			fmt.Scanln(&aStr)
			s = s+aStr
			res = append(res, s)
		case 2:
			var x, y int
			fmt.Scanf("%d %d\n", &x, &y)
			s = s[x:x+y]
			res = append(res, s)
		case 3:
			var start int
			var cStr string
			fmt.Scanf("%d %s\n", &start, &cStr)
			s = s[:start]+cStr+s[start:]
			res = append(res, s)
		case 4:
			var dStr string
			fmt.Scanln(&dStr)
			res = append(res, strconv.Itoa(strings.Index(s, dStr)))
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(res[i])
	}
}