package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	local := 5 * n
	luogu := 3 * n + 11
	if local < luogu {
		fmt.Printf("Local")
	} else {
		fmt.Printf("Luogu")
	}

}
