package main

import "fmt"

func main() {

	var a string
	fmt.Scanln(&a)

	for i := range []byte(a) {
		if a[i] <= 'z' && a[i] >= 'a' {
			x := a[i]
			x = x + 'A' - 'a'
			fmt.Printf("%c", x)
		} else {
			fmt.Printf("%c", a[i])
		}
	}
}
