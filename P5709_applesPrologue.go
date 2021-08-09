// there are traps:
// (1)t = 0		(2)m - temp < 0		(3)integer, not float

package main

import "fmt"

func main() {

	var m, t, s, res int
	var temp int
	fmt.Scanln(&m, &t, &s)

	if t == 0 {
		res = 0
	} else {
		temp = s / t
		if s % t != 0 {
			temp += 1
		}
		res = m - temp
		if res < 0 {
			res = 0
		}
	}
	fmt.Printf("%v", res)
}