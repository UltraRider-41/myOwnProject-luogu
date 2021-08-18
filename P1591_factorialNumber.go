// high-precision

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {

	var t int

	fmt.Scanln(&t)

	type num struct {
		n int64
		a int64
	}
	var temp num
	var list = make([]num, 0)
	for i := 0; i < t; i++ {
		fmt.Scanln(&temp.n, &temp.a)
		list = append(list, temp)
	}
	// fmt.Println(list)
	for i := 0; i < t; i++ {
		facRes := big.NewInt(1)
		for j := 1; j <= int(list[i].n); j++ {
			facRes.Mul(facRes, big.NewInt(int64(j)))
		}
		// fmt.Println(facRes)
		facStr := facRes.String()
		subStr := strconv.Itoa(int(list[i].a))
		res := strings.Count(facStr, subStr)
		fmt.Println(res)
	}

}
