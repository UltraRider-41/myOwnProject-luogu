package main

import (
	"fmt"
	"strconv"
)

func main() {

	var weight int
	var height float64
	fmt.Scanln(&weight, &height)

	var res float64
	var floatString string
	res = float64(weight) / (height * height)
	if res < 18.5 {
		fmt.Printf("Underweight")
	} else if res >= 18.5 && res < 24 {
		fmt.Printf("Normal")
	} else {
		floatString = strconv.FormatFloat(res, 'g', 6, 64)
		fmt.Printf("%s\nOverweight", floatString)
	}
}
