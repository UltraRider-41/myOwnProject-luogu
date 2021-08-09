package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isLeap(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, _, _ := reader.ReadLine()

	var res, max, min int
	var letter = make([]int, 26)
	for i := 0; i < len(s); i++ {
		for j := 97; j <= 122; j++ {
			if s[i] == byte(j) {
				letter[j-97] += 1
			}
		}
	}

	max = 0
	min = 101
	for i := range letter {
		if min > letter[i] && letter[i] != 0 {
			min = letter[i]
		}
		if max < letter[i] {
			max = letter[i]
		}
	}
	res = max - min
	if isLeap(res) {
		fmt.Printf("Lucky Word\n%d", res)
	} else {
		fmt.Printf("No Answer\n0")
	}
}
