package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	s := string(data)
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res++
		} else if s[i] == 'a' || s[i] == 'd' || s[i] == 'g' || s[i] == 'j' || s[i] == 'm' || s[i] == 'p' || s[i] == 't' || s[i] == 'w' {
			res++
		} else if s[i] == 'b' || s[i] == 'e' || s[i] == 'h' || s[i] == 'k' || s[i] == 'n' || s[i] == 'q' || s[i] == 'u' || s[i] == 'x' {
			res += 2
		} else if s[i] == 'c' || s[i] == 'f' || s[i] == 'i' || s[i] == 'l' || s[i] == 'o' || s[i] == 'r' || s[i] == 'v' || s[i] == 'y' {
			res += 3
		} else if s[i] == 's' || s[i] == 'z' {
			res += 4
		}
	}
	fmt.Printf("%d", res)
}
