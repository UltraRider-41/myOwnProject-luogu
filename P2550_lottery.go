package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanln(&n)

	type number struct {
		a int
		b int
		c int
		d int
		e int
		f int
		g int
	}

	var a, b, c, d, e, f, g int

	var prize number
	fmt.Scan(&prize.a, &prize.b, &prize.c, &prize.d, &prize.e, &prize.f, &prize.g)

	var lotteryNum = make([]number, 0)
	var res = make([]int, n) // Store the number of winning digits for each lottery ticket
	for i := 0; i < n; i++ {
		var temp number
		fmt.Scan(&temp.a, &temp.b, &temp.c, &temp.d, &temp.e, &temp.f, &temp.g)
		lotteryNum = append(lotteryNum, temp)
	}

	var prizeList = []int{prize.a, prize.b, prize.c, prize.d, prize.e, prize.f, prize.g}
	for i := 0; i < n; i++ {
		var temp = []int{
			lotteryNum[i].a,
			lotteryNum[i].b,
			lotteryNum[i].c,
			lotteryNum[i].d,
			lotteryNum[i].e,
			lotteryNum[i].f,
			lotteryNum[i].g,
			}
		for j := 0; j < 7; j++ {
			for k := 0; k < 7; k++ {
				if temp[j] == prizeList[k] {
					res[i]++
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		switch res[i] {
			case 7:
				a++
			case 6:
				b++
			case 5:
				c++
			case 4:
				d++
			case 3:
				e++
			case 2:
				f++
			case 1:
				g++
		}
	}
	fmt.Printf("%d %d %d %d %d %d %d", a, b, c, d, e, f, g)
}