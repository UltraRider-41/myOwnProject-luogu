package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var s = make([]string, 0)
	var temp string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	for i := 0; i < n; i++ {
		s = append(s, string(data[i]))
		temp = temp + s[i]
	}
	// fmt.Printf("%T %#v\n", s[0], s[0])

	max := 0
	// replace to V""
	for j := 0; j < n; j++ {
		var tempLetter string
		tempLetter = s[j]
		s[j] = "V"
		var res int
		for i := 0; i < n-1; i++ {
			if s[i] == "V" && s[i+1] == "K" {
				res++
			}
			if max < res {
				max = res
			}
		}
		s[j] = tempLetter
	}
	// replace to "K"
	for j := 0; j < n; j++ {
		var tempLetter string
		tempLetter = s[j]
		s[j] = "K"
		var res int
		for i := 0; i < n-1; i++ {
			if s[i] == "V" && s[i+1] == "K" {
				res++
			}
			if max < res {
				max = res
			}
		}
		s[j] = tempLetter
	}
	fmt.Printf("%d", max)

}
