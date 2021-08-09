// base 10 -> base r(r is a minus number)

package main

import (
	"fmt"
)

func main() {

	var n, base int64
	var remainder int64
	var remainderList = make([]int64, 0)
	fmt.Scanln(&n, &base)
	// the remainder must be a positive number, so when it becomes negative, the quotient will plus 1
	// and recompute the remainder(this time it'll be positive)
	for quotient := n; quotient <= 0 || quotient >= -base; {
		temp := quotient/base
		if quotient - (temp * base) < 0 {
			temp += 1
			remainder = quotient - (temp * base)
			remainderList = append(remainderList, remainder)
			quotient = temp
		} else {
			remainder = quotient - (temp * base)
			remainderList = append(remainderList, remainder)
			quotient = temp
		}
		if quotient > 0 && quotient < -base {
			remainderList = append(remainderList, quotient)
		}
	}

	fmt.Printf("%d=", n)
	for i := len(remainderList)-1; i >= 0; i-- {
		//  base > 10, so we need to change some numbers to letters(e.g. 11 -> B, 15 -> F)
		if remainderList[i] > 9 {
			a := remainderList[i] - 9
			b := 64 + a
			fmt.Printf("%v", string(b))
		} else {
			fmt.Printf("%d", remainderList[i])
		}
	}
	fmt.Printf("(base%d)", base)
}