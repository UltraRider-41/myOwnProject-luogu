package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ratio = 2^(n-1), n is the number of set's elements
func calRatio(a int) int {

	var res int
	if a == 1 {
		res = 1
	} else {
		res = int(math.Pow(2.0, float64(a-1)))
	}

	return res
}

func main() {

	var set = make([]int, 0)
	var s = make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	ar, _ := reader.ReadString('\n')
	s = strings.Fields(ar)
	// fmt.Println(s[0])

	for i := 0; i < len(s); i++ {
		temp, _ := strconv.Atoi(s[i])
		set = append(set, temp)
	}

	// fmt.Printf("%T %#v", set[0], set[1])

	var ratio int
	var res int
	ratio = calRatio(len(set))
	// ratio * sum(set)
	for i := 0; i < len(set); i++ {
		res += ratio * set[i]
	}
	fmt.Printf("%d", res)
}
