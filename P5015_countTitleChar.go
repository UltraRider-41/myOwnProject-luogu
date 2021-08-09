package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	var s string
	s = string(data)

	var num int
	for i := range s {
		if s[i] != ' ' {
			num++
		}
	}
	fmt.Printf("%d", num)
}
