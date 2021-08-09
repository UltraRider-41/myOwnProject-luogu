package main

import (
	"fmt"
	"strconv"
)

// reverse a non-negative number
func reverse(num string) string {

	var res string
	var temp = make([]byte, 0)
	for i := len(num)-1; i >= 0; i-- {
		temp = append(temp, num[i])
	}
	res = string(temp)
	return res
}

// calculate orderNum + reverseOrderNum
func sum(a string, b string, base int) string {

	var res string
	var temp = make([]byte, 0)
	// quotient is carry
	// remainder is local
	var quotient int
	var remainder int


	// a and b have a same digit, here it can also be len(b)-1
	for i := len(a)-1; i >= -1; i-- {
		if base != 16 {
			if i > -1 {
				var aNum int // int version of a's element
				var bNum int // int version of b's element
				aNum,_ = strconv.Atoi(string(a[i]))
				bNum,_ = strconv.Atoi(string(b[i]))
				remainder = (aNum + bNum + quotient) % base
				quotient = (aNum + bNum + quotient) / base
				temp = append(temp, byte(remainder + 48)) // the ascii code of number 0~9 : 48~57
			} else if i == -1 {
				if quotient >= 1 {
					temp = append(temp, byte(quotient + 48))
				}
			}
		} else if base == 16 {
			if i > -1 {
				var aNum int64 // int version of a's element
				var bNum int64 // int version of b's element
				aNum, _ = strconv.ParseInt(string(a[i]), base, 64)
				bNum, _ = strconv.ParseInt(string(b[i]), base, 64)
				remainder = (int(aNum + bNum) + quotient) % base
				quotient = (int(aNum + bNum) + quotient) / base
				temp = append(temp, []byte((strconv.FormatInt(int64(remainder), 16)))[0]) // the ascii code of number 0~9 : 48~57
			} else if i == -1 {
				if quotient >= 1 {
					temp = append(temp, []byte((strconv.FormatInt(int64(quotient), 16)))[0])
				}
			}
		}
	}

	res = string(temp)  // haven't reverse the num string, but it's ok in this problem
	return res
}

// determine if the number is a palindrome
func isPalindrome(num string) int {

	flag := 0
	if reverse(num) == num {
		flag = 1
	}

	return flag
}

/*
// n-base number -> 10-base number
func NtoTen(num string, base int) string {
	 var res string


	 return res
}

// 10-base number -> n-base number
func TentoN(num string, base int) string {
	var res string

	// cannot just convert like this, the number will overflow!!!
	numInt, _ := strconv.Atoi(num)
	res = strconv.FormatInt(int64(numInt), base)
	// append ...

	return res
}
 */

func main() {

	// N base number M
	// 2 <= N <= 10 or N = 16
	var N int
	var M string
	fmt.Scanln(&N)
	fmt.Scanln(&M)

	// if M is a minus number, we need to get its absolute value, in other words, we should delete the '-' element
	if M[0] == '-' {
		M = M[1:]
	}

	var reverseM string
	var res = M
	flag := 0

	for i := 0; i < 30; i++ {
		flag = isPalindrome(res)
		if flag != 0 {
			fmt.Printf("STEP=%d", i)
			break
		}
		reverseM = reverse(res)
		/*
		var convertM string
		var convertReverseM string
		convertM = NtoTen(res, N)
		convertReverseM = NtoTen(reverseM, N)
		 */
		res = sum(res, reverseM, N)
		/*
		res = TentoN(res, N)
		 */
	}
	if flag == 0 {
		fmt.Printf("Impossible!")
	}

}
