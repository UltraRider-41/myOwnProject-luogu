package main

import "fmt"

// actually this function can reverse any string, not just float number
func reverseFloatNum(num string) string {

	var res string
	var temp = make([]byte, 0)
	for i := len(num)-1; i >= 0; i-- {
		temp = append(temp, num[i])
	}
	res = string(temp)
	return res
}

func main() {

	var num string
	fmt.Scanln(&num)
	num = reverseFloatNum(num)
	fmt.Printf("%v", num)

}
