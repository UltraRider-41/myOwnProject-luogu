package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var r, c, k int
	fmt.Scanln(&r, &c, &k)

	var court = make([][]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < r; i++ {
		var temp = make([]string, c)
		context, _ := reader.ReadString('\n')
		for j := 0; j < c; j++ {
			temp[j] = string(context[j])
		}
		court = append(court, temp)
	}

	flag := 1
	var horizon, vertical, res int
	// count horizontal spare space
	for x := 0; x < r; x++ {
		for y := 0; y <= c-k; y++ {
			for z := 0; z < k; z++ {
				if court[x][y+z] != "." {
					flag = 0
					break
				}
			}
			if flag == 1 {
				horizon++
			}
			flag = 1
		}
	}
	// count vertical spare space
	for x := 0; x < c; x++ {
		for y := 0; y <= r-k; y++ {
			for z := 0; z < k; z++ {
				if court[y+z][x] != "." {
					flag = 0
					break
				}
			}
			if flag == 1 {
				vertical++
			}
			flag = 1
		}
	}

	res = horizon + vertical
	if k == 1 { // notice: if k = 1, just need count once(horizontal or vertical)
		res /= 2
	}
	fmt.Println(res)
}
