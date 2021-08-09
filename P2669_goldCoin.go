package main

import "fmt"

var temp int = 1
var i = 1

func main() {

	var k int
	fmt.Scanln(&k)
	res := 0

	j := 1
	for i = 1; i <= k; {
		if j <= temp {
			res += temp
			if j == temp {
				temp++
				j = 0
			}
		}
		j++
		i++
	}
	fmt.Printf("%d", res)
}
