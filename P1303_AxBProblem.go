// high-precision A*B

package main

import (
	"fmt"
	"math/big"
)

func main() {

	var temp1, temp2 string

	fmt.Scanln(&temp1)
	fmt.Scanln(&temp2)
	a, _ := new(big.Int).SetString(temp1, 10)
	b, _ := new(big.Int).SetString(temp2, 10)

	// fmt.Printf("%T %v", a, b)
	res := big.NewInt(1)
	res = res.Mul(a, b)

	fmt.Printf("%v", res)

}
