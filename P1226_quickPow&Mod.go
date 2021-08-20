package main

import (
	"fmt"
	"math/big"
)

func specialQuickPow(x *big.Int, n int64, mod *big.Int) *big.Int {

	if x == big.NewInt(0) {
		return big.NewInt(0)
	} else {
		if n == 0 {
			temp := big.NewInt(1)
			return temp.Mod(temp, mod)
		} else if n%2 == 1 {
			temp := specialQuickPow(x, n-1, mod)
			temp.Mul(temp, x)
			return temp.Mod(temp, mod)
		} else if n%2 == 0 {
			temp := specialQuickPow(x, n/2, mod)
			temp.Mul(temp, temp)
			return temp.Mod(temp, mod)
		} else {
			return nil
		}
	}
}

func main() {

	var x, y, z int64
	fmt.Scanln(&x, &y, &z)

	a := big.NewInt(x)
	p := big.NewInt(z)

	res := specialQuickPow(a, y, p)
	fmt.Printf("%v^%v mod %v=%v", a, y, p, res)
}
