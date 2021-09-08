package main

import (
	"bufio"
	"fmt"
	"os"
)

// white -> blue -> red
func main() {

	var n, m int
	var grid int
	fmt.Scanln(&n, &m)
	if n < 3 {
		panic("N must > 3!")
	}

	reader := bufio.NewReader(os.Stdin)
	var flag = make([][]string, 0)
	for i := 0; i < n; i++ {
		var tempList = make([]string, m)
		context, _ := reader.ReadString('\n')
		for j := 0; j < m; j++ {
			tempList[j] = string(context[j])
		}
		flag = append(flag, tempList)
	}

	// head, white
	for i := 0; i < m; i++ {
		if flag[0][i] != "W" {
			grid++
			// flag[0][i] = "W"
		}
	}

	// tail, red
	for i := 0; i < m; i++ {
		if flag[n-1][i] != "R" {
			grid++
			// flag[n-1][i] = "R"
		}
	}

	// body
	var countGrid = make([][]int, 0)
	for i := 0; i < n-2; i++ {
		var wbr = make([]int, 3)
		for j := 0; j < m; j++ {
			switch flag[i+1][j] {
			case "W":
				wbr[0]++
			case "B":
				wbr[1]++
			case "R":
				wbr[2]++
			}
		}
		countGrid = append(countGrid, wbr)
	}

	max := 0
	num := 0
	for i := 0; i < n-2; i++ {
		num = 0 // notice: must initialize num in each new loop
		bStart := i
		for x := 0; x < bStart; x++ {
			// count white
			num += countGrid[x][0]
		}
		maxTemp := 0
		for j := bStart; j < n-2; j++ {
			numTemp := 0
			bEnd := j
			if bEnd > bStart {
				for z := bEnd + 1; z < n-2; z++ {
					// count red
					numTemp += countGrid[z][2]
				}
				for y := bStart; y <= bEnd; y++ {
					// count blue
					numTemp += countGrid[y][1]
				}
				if numTemp > maxTemp {
					maxTemp = numTemp
				}
			} else if bEnd == bStart {
				// count blue
				numTemp += countGrid[bEnd][1]
				for z := bEnd + 1; z < n-2; z++ {
					// count red
					numTemp += countGrid[z][2]
				}
				if numTemp > maxTemp {
					maxTemp = numTemp
				}
			}
		}
		num += maxTemp
		if num > max {
			max = num
		}
	}
	num = (n-2)*m - max
	grid += num
	fmt.Println(grid)
}
