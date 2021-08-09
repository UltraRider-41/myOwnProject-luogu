package main

import (
	"fmt"
	"math"
)

func main() {

	var id int
	fmt.Scanln(&id)

	switch id {
		case 1:
			fmt.Printf("I love Luogu!")
		case 2:
			fmt.Printf("6 4")
		case 3:
			fmt.Printf("3\n12\n2")
	    case 4:
	    	res := 500.0 / 3.0
	    	fmt.Printf("%.3f", res)
		case 5:
			fmt.Printf("15")
		case 6:
			var diagonal float64
			diagonal = math.Sqrt(6.0 * 6.0 + 9.0 * 9.0)
			fmt.Printf("%.4f",diagonal)
		case 7:
			fmt.Printf("110\n90\n0")
		case 8:
			var radius, perimeter, square, volume, pai float64
			radius = 5.0
			pai = 3.141593
			perimeter = 2 * pai * radius
			square = pai * radius * radius
			volume = 4.0 / 3.0 * pai * radius * radius * radius
			fmt.Printf("%.4f\n%.4f\n%.3f", perimeter, square, volume)
		case 9:
			fmt.Printf("22")
		case 10:
			fmt.Printf("9")
		case 11:
			var t float64
			t = 100.0 / 3.0
			fmt.Printf("%.4f", t)
		case 12:
			fmt.Printf("13\nR")
		case 13:
			var pai, roundVolume, edge float64
			pai = 3.141593
			r1 := 4.0
			r2 := 10.0
			roundVolume = 4.0 / 3.0 * (r1 * r1 * r1 + r2 * r2 * r2) * pai
			edge = math.Cbrt(roundVolume)
			fmt.Printf("%d", int(edge))
		case 14:
			fmt.Printf("50")
	}

}
