package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n, m int
	fmt.Scanln(&n, &m)

	type person struct {
		direction int // 0 means in, 1 means out
		position  string
	}

	var temp1 person
	var list = make([]person, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&temp1.direction, &temp1.position)
		list[i] = temp1
	}

	local := 0
	var tempA, tempS int

	// cannot use fmt.scanf, will cause TLE
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		tempA, _ = strconv.Atoi(split[0])
		tempS, _ = strconv.Atoi(split[1])

		if tempA == 0 {
			switch list[local].direction {
			case 0:
				local = local - tempS
				if local < 0 {
					local += n
				}
			case 1:
				local = local + tempS
				if local >= n {
					local -= n
				}
			}
		} else if tempA == 1 {
			switch list[local].direction {
			case 1:
				local = local - tempS
				if local < 0 {
					local += n
				}
			case 0:
				local = local + tempS
				if local >= n {
					local -= n
				}
			}
		}
	}

	fmt.Println(list[local].position)
}
