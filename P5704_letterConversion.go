package main

import (
	"fmt"
)

func letterToInt(letter string) int {

	var res int
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 26 ; i++ {
		if string(alphabet[i]) == letter {
			res = i + 65
		}
	}
	return res
}

func main() {

	var letter string
	var temp int

	fmt.Scanln(&letter)
	temp = letterToInt(letter)
	fmt.Printf("%v", string(temp))
}
