// high-precision test, use big package

package main

import (
	"fmt"
	"math/big"
)

func main() {

	var n int64
	fmt.Scanln(&n)

	res := big.NewInt(0)
	x := big.NewInt(1)
	inc := big.NewInt(1)
	var i int64 = 1

	for ; i <= n; i++ {
		var j int64 = 1
		N := big.NewInt(i)
		for ; j < i; j++ {
			N = N.Mul(N, x)
			x = x.Add(x, inc)
		}
		res = res.Add(res, N)
		x = big.NewInt(1)
	}
	fmt.Printf("%v", res)
}
