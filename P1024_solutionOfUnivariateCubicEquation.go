package main

import "fmt"

func main() {

	var a, b, c, d float64
	var A, B, C ,D float64
	fmt.Scanln(&a, &b, &c, &d)
	A = 100.0 * a
	B = 100.0 * b
	C = 100.0 * c
	D = 100.0 * d

	i := 0

	for num := -10000; num <= 10000; num += 1 {
		if (float64(int(A)) * float64(num * num * num) / 1000000.0) + (float64(int(B)) * float64(num * num) / 10000.0) + (float64(int(C)) * float64(num) / 100.0) + D == 0.0 {
			if i != 2 {
				fmt.Printf("%.2f ", float64(num) / 100.0)
			} else {
				fmt.Printf("%.2f", float64(num) / 100.0)
			}
			i++
		}
	}
}