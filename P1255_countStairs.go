// do not use recursion, because it'll waste too much time

package main

import (
	"fmt"
	"math/big"
)

func main() {

	var n int64
	fmt.Scanln(&n)

	var list = make([]*big.Int, 0)
	var i int64
	if n > 0 {
		for i = 0; i < n; i++ {
			if i == 0 {
				list = append(list, big.NewInt(1))
			} else if i == 1 {
				list = append(list, big.NewInt(2))
			} else {
				temp := big.NewInt(0)
				temp.Add(temp, list[i-1])
				temp.Add(temp, list[i-2])
				list = append(list, temp)
			}
		}
		var res *big.Int
		res = list[n-1]
		fmt.Println(res)
	} else {
		fmt.Println(0)
	}

}
