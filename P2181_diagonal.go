package main

import (
	"fmt"
)

func  main() {

	var n uint64
	fmt.Scanln(&n)
	var res uint64

	res =  n * (n - 1) / 2 * (n - 2) / 3 * (n - 3) / 4
	fmt.Printf("%v", res)

}
