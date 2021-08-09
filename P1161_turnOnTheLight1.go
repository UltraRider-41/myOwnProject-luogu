// violence, not xor. xor method is more complex than violence method and takes more time

package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var light = make([]int, 2000001)
	var aTemp float64
	var tTemp int
	var aList = make([]float64, 0)
	var tList = make([]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scanln(&aTemp, &tTemp)
		aList = append(aList, aTemp)
		tList = append(tList, tTemp)
		for j := 1; j <= tTemp; j++ {
			light[int(float64(j)*aTemp)-1] = (light[int(float64(j)*aTemp)-1]+1) % 2
		}
	}

	// fmt.Println(light)
	for i := 0; ; i++ {
		if light[i] == 1 {
			fmt.Printf("%d", i+1)
			break
		}
	}
}
