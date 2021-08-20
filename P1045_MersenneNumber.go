package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

// To prevent timeouts(because digits may be very large), every recursion the function "quickPow1()" will
// mod 10^500, it will reduce operational volume significantly
var modNum = quickPow2(big.NewInt(10), 500)

// recursive quick power
func quickPow1(x *big.Int, n int) *big.Int {

	if n == 0 {
		return big.NewInt(1)
	} else if n%2 == 1 {
		temp := quickPow1(x, n-1)
		temp.Mul(temp, big.NewInt(2))
		return temp.Mod(temp, modNum)
	} else if n%2 == 0 {
		temp := quickPow1(x, n/2)
		res := temp.Mul(temp, temp)
		return res.Mod(res, modNum)
	} else {
		return nil
	}
}

// loop quick power
func quickPow2(x *big.Int, n int) *big.Int {

	res := big.NewInt(1)
	for n != 0 {
		if n&1 != 0 {
			res.Mul(res, x)
		}
		x.Mul(x, x)
		n >>= 1
	}
	return res
}

func main() {

	// process the input data
	var p int
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	p, _ = strconv.Atoi(string(data))

	var digits int
	mersenne := quickPow1(big.NewInt(2), p)
	mersenne.Sub(mersenne, big.NewInt(1))
	// fmt.Println("mersenne is: ", mersenne)
	temp := mersenne.String()
	digits = int(float64(p)*math.Log10(float64(2))) + 1
	fmt.Println(digits)
	digits = len(temp)

	// print the result
	if digits-500 < 0 {
		for i := 0; i < 500; i++ {
			if i < 500-digits {
				fmt.Printf("0")
			} else {
				fmt.Printf("%v", temp[i-(500-digits)]-48)
			}
			if (i+1)%50 == 0 {
				fmt.Println()
			}
		}
	} else {
		for i := digits - 500; i < digits; i++ {
			fmt.Printf("%v", temp[i]-48)
			if (i+1-(digits-500))%50 == 0 {
				fmt.Println()
			}
		}
	}

}
