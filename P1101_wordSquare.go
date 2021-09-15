package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkRowRight(square [][]string, i int, j int, n int) bool {

	if n-j < 7 {
		return false
	}
	if square[i][j] != "y" || square[i][j+1] != "i" || square[i][j+2] != "z" || square[i][j+3] != "h" || square[i][j+4] != "o" || square[i][j+5] != "n" || square[i][j+6] != "g" {
		return false
	}

	return true
}

func checkRowLeft(square [][]string, i int, j int, n int) bool {

	if j < 6 {
		return false
	}
	if square[i][j] != "y" || square[i][j-1] != "i" || square[i][j-2] != "z" || square[i][j-3] != "h" || square[i][j-4] != "o" || square[i][j-5] != "n" || square[i][j-6] != "g" {
		return false
	}
	return true
}

func checkColDown(square [][]string, i int, j int, n int) bool {

	if n-i < 7 {
		return false
	}
	if square[i][j] != "y" || square[i+1][j] != "i" || square[i+2][j] != "z" || square[i+3][j] != "h" || square[i+4][j] != "o" || square[i+5][j] != "n" || square[i+6][j] != "g" {
		return false
	}
	return true
}

func checkColUp(square [][]string, i int, j int, n int) bool {

	if i < 6 {
		return false
	}
	if square[i][j] != "y" || square[i-1][j] != "i" || square[i-2][j] != "z" || square[i-3][j] != "h" || square[i-4][j] != "o" || square[i-5][j] != "n" || square[i-6][j] != "g" {
		return false
	}
	return true
}

func checkSlashRightDown(square [][]string, i int, j int, n int) bool {

	if n-i < 7 || n-j < 7 {
		return false
	}
	if square[i][j] != "y" || square[i+1][j+1] != "i" || square[i+2][j+2] != "z" || square[i+3][j+3] != "h" || square[i+4][j+4] != "o" || square[i+5][j+5] != "n" || square[i+6][j+6] != "g" {
		return false
	}
	return true
}

func checkSlashLeftDown(square [][]string, i int, j int, n int) bool {

	if n-i < 7 || j < 6 {
		return false
	}
	if square[i][j] != "y" || square[i+1][j-1] != "i" || square[i+2][j-2] != "z" || square[i+3][j-3] != "h" || square[i+4][j-4] != "o" || square[i+5][j-5] != "n" || square[i+6][j-6] != "g" {
		return false
	}
	return true
}

func checkSlashRightUp(square [][]string, i int, j int, n int) bool {

	if n-j < 7 || i < 6 {
		return false
	}
	if square[i][j] != "y" || square[i-1][j+1] != "i" || square[i-2][j+2] != "z" || square[i-3][j+3] != "h" || square[i-4][j+4] != "o" || square[i-5][j+5] != "n" || square[i-6][j+6] != "g" {
		return false
	}
	return true
}

func checkSlashLeftUp(square [][]string, i int, j int, n int) bool {

	if j < 6 || i < 6 {
		return false
	}
	if square[i][j] != "y" || square[i-1][j-1] != "i" || square[i-2][j-2] != "z" || square[i-3][j-3] != "h" || square[i-4][j-4] != "o" || square[i-5][j-5] != "n" || square[i-6][j-6] != "g" {
		return false
	}
	return true
}

func main() {

	type information struct {
		row  int
		col  int
		kind int // 1 means rowRight, 2 means rowLeft, 3 means colDown, 4 means colUp, 5 means slashRightDown, 6 means slashLeftDown, 7 means slashRightUp, 8 means slashLeftUp
	}

	// 7 <= n <= 100
	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	var square = make([][]string, 0)
	for i := 0; i < n; i++ {
		var temp = make([]string, n)
		context, _ := reader.ReadString('\n')
		for j := 0; j < n; j++ {
			temp[j] = string(context[j])
		}
		square = append(square, temp)
	}

	var informationList = make([]information, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if square[i][j] == "y" {
				if checkRowRight(square, i, j, n) {
					temp := information{row: i, col: j, kind: 1}
					informationList = append(informationList, temp)
				}
				if checkRowLeft(square, i, j, n) {
					temp := information{row: i, col: j, kind: 2}
					informationList = append(informationList, temp)
				}
				if checkColDown(square, i, j, n) {
					temp := information{row: i, col: j, kind: 3}
					informationList = append(informationList, temp)
				}
				if checkColUp(square, i, j, n) {
					temp := information{row: i, col: j, kind: 4}
					informationList = append(informationList, temp)
				}
				if checkSlashRightDown(square, i, j, n) {
					temp := information{row: i, col: j, kind: 5}
					informationList = append(informationList, temp)
				}
				if checkSlashLeftDown(square, i, j, n) {
					temp := information{row: i, col: j, kind: 6}
					informationList = append(informationList, temp)
				}
				if checkSlashRightUp(square, i, j, n) {
					temp := information{row: i, col: j, kind: 7}
					informationList = append(informationList, temp)
				}
				if checkSlashLeftUp(square, i, j, n) {
					temp := information{row: i, col: j, kind: 8}
					informationList = append(informationList, temp)
				}
			}
		}
	}

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			square[x][y] = "*"
		}
	}

	for x := 0; x < len(informationList); x++ {
		switch informationList[x].kind {
		case 1:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row][informationList[x].col+1] = "i"
			square[informationList[x].row][informationList[x].col+2] = "z"
			square[informationList[x].row][informationList[x].col+3] = "h"
			square[informationList[x].row][informationList[x].col+4] = "o"
			square[informationList[x].row][informationList[x].col+5] = "n"
			square[informationList[x].row][informationList[x].col+6] = "g"
		case 2:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row][informationList[x].col-1] = "i"
			square[informationList[x].row][informationList[x].col-2] = "z"
			square[informationList[x].row][informationList[x].col-3] = "h"
			square[informationList[x].row][informationList[x].col-4] = "o"
			square[informationList[x].row][informationList[x].col-5] = "n"
			square[informationList[x].row][informationList[x].col-6] = "g"
		case 3:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row+1][informationList[x].col] = "i"
			square[informationList[x].row+2][informationList[x].col] = "z"
			square[informationList[x].row+3][informationList[x].col] = "h"
			square[informationList[x].row+4][informationList[x].col] = "o"
			square[informationList[x].row+5][informationList[x].col] = "n"
			square[informationList[x].row+6][informationList[x].col] = "g"
		case 4:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row-1][informationList[x].col] = "i"
			square[informationList[x].row-2][informationList[x].col] = "z"
			square[informationList[x].row-3][informationList[x].col] = "h"
			square[informationList[x].row-4][informationList[x].col] = "o"
			square[informationList[x].row-5][informationList[x].col] = "n"
			square[informationList[x].row-6][informationList[x].col] = "g"
		case 5:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row+1][informationList[x].col+1] = "i"
			square[informationList[x].row+2][informationList[x].col+2] = "z"
			square[informationList[x].row+3][informationList[x].col+3] = "h"
			square[informationList[x].row+4][informationList[x].col+4] = "o"
			square[informationList[x].row+5][informationList[x].col+5] = "n"
			square[informationList[x].row+6][informationList[x].col+6] = "g"
		case 6:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row+1][informationList[x].col-1] = "i"
			square[informationList[x].row+2][informationList[x].col-2] = "z"
			square[informationList[x].row+3][informationList[x].col-3] = "h"
			square[informationList[x].row+4][informationList[x].col-4] = "o"
			square[informationList[x].row+5][informationList[x].col-5] = "n"
			square[informationList[x].row+6][informationList[x].col-6] = "g"
		case 7:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row-1][informationList[x].col+1] = "i"
			square[informationList[x].row-2][informationList[x].col+2] = "z"
			square[informationList[x].row-3][informationList[x].col+3] = "h"
			square[informationList[x].row-4][informationList[x].col+4] = "o"
			square[informationList[x].row-5][informationList[x].col+5] = "n"
			square[informationList[x].row-6][informationList[x].col+6] = "g"
		case 8:
			square[informationList[x].row][informationList[x].col] = "y"
			square[informationList[x].row-1][informationList[x].col-1] = "i"
			square[informationList[x].row-2][informationList[x].col-2] = "z"
			square[informationList[x].row-3][informationList[x].col-3] = "h"
			square[informationList[x].row-4][informationList[x].col-4] = "o"
			square[informationList[x].row-5][informationList[x].col-5] = "n"
			square[informationList[x].row-6][informationList[x].col-6] = "g"
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s", square[i][j])
			if (j+1)%n == 0 {
				fmt.Println()
			}
		}
	}
}
