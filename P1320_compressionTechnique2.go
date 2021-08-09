package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	s := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			s = append(s, text)
		} else {
			break
		}
	}

	var str string
	str = strings.Join(s, str)
	// fmt.Printf("%T %#v", str, str)

	var code = make([]int, 0)
	var num0, num1 int
	var flag0, flag1 int // actually flag0 is useless, just need flag1
	code = append(code, len(str)/len(s))

	for i := 0; i < len(str); i++ {

		if str[i] == '0' {
			if num0 == 0 && flag1 != 0 {
				code = append(code, num1)
			}
			num1 = 0
			num0++
			flag0++
			if i == len(str)-1 {
				code = append(code, num0)
			}
		} else if str[i] == '1' {
			if num1 == 0 {
				code = append(code, num0)
			}
			num0 = 0
			num1++
			flag1++
			if i == len(str)-1 {
				code = append(code, num1)
			}
		}
	}

	for i := 0; i < len(code); i++ {
		fmt.Printf("%d ", code[i])
	}
}
