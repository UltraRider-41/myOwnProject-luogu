package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var a, b, c, d, e, f string
	fmt.Scanln(&a, &b, &c, &d, &e, &f)
	var s = make([]string, 6)
	s[0] = a
	s[1] = b
	s[2] = c
	s[3] = d
	s[4] = e
	s[5] = f
	// fmt.Printf("%T %#v\n", s, s)

	var notExist int
	var num = make([]int, 6)
	for i := 0; i < 6; i++ {
		switch s[i] {
			case "one", "a", "first", "another" :
				num[i] = 1
			case "two", "second" :
				num[i] = 2
			case "three", "third" :
				num[i] = 3
			case "four" :
				num[i] = 4
			case "five" :
				num[i] = 5
			case "six" :
				num[i] = 6
			case "seven" :
				num[i] = 7
			case "eight" :
				num[i] = 8
			case "nine" :
				num[i] = 9
			case "ten" :
				num[i] = 10
			case "eleven" :
				num[i] = 11
			case "twelve" :
				num[i] = 12
			case "thirteen" :
				num[i] = 13
			case "fourteen" :
				num[i] = 14
			case "fifteen" :
				num[i] = 15
			case "sixteen" :
				num[i] = 16
			case "seventeen" :
				num[i] = 17
			case "eighteen" :
				num[i] = 18
			case "nineteen" :
				num[i] = 19
			case "twenty" :
				num[i] = 20
			default:
				notExist++
		}
	}

	for i := 0; i < 6; i++ {
		num[i] = (num[i]*num[i]) % 100
	}

	sort.Ints(num)
	// fmt.Printf("%T %#v\n", num, num)
	 var res string
	 for i := 0; i < 6; i++ {
	 	if num[i] / 10 != 0 {
			temp := strconv.Itoa(num[i])
			res = res + temp
		} else if num[i] / 10 == 0 {
			temp := strconv.Itoa(num[i])
			temp = "0" + temp
			res = res + temp
		}
	}
	flag := 0
	for i := 0; i < len(res); i++ {
		if res[i] != 48 && flag == 0 {
			fmt.Printf("%v", res[i]-48)
			flag = 1
		} else if flag == 1 {
			fmt.Printf("%v", res[i]-48)
		}
	}
	// be careful of the non-special-word-exist situation
	if notExist == 6 {
		fmt.Printf("0")
	}
}
