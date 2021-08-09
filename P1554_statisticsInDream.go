package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var m, n int
	fmt.Scanln(&m, &n)

	// store numbers between m and n(m and n included)
	var numTemp int
	var numList = make([]int, 0)
	for j := m; j <= n; j++ {
		numTemp = j
		numList = append(numList, numTemp)
	}

	// turn int to string, be convenient to count 0~9
	var stringTemp string
	var stringList = make([]string, 0)
	for j := m; j <= n; j++ {
		stringTemp = strconv.Itoa(j)
		stringList = append(stringList, stringTemp)
	}

	// count string 0~9
	var a, b, c, d, e, f, g, h, i, k int
	for j := 0; j < len(stringList); j++ {
		a += strings.Count(stringList[j], "0")
		b += strings.Count(stringList[j], "1")
		c += strings.Count(stringList[j], "2")
		d += strings.Count(stringList[j], "3")
		e += strings.Count(stringList[j], "4")
		f += strings.Count(stringList[j], "5")
		g += strings.Count(stringList[j], "6")
		h += strings.Count(stringList[j], "7")
		i += strings.Count(stringList[j], "8")
		k += strings.Count(stringList[j], "9")
	}

	fmt.Printf("%d %d %d %d %d %d %d %d %d %d", a, b, c, d, e, f, g, h, i, k)
}
