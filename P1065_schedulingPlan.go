package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// represents occupied points in timelines
type occupy struct {
	start int
	end   int
}

type occupyRow []occupy

// custom function which can scan a line
func scanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	return string(b)
}

// row means which timeline/machine
func check(occupyPoint *[][]occupy, start int, time int, row int) int {

	var newEnd int
	end := start + time - 1
	myOccupy := *occupyPoint
	length := len(myOccupy[row])
	for i := 0; i < length; i++ {
		if end < myOccupy[row][i].start {
			break
		} else if start > myOccupy[row][i].end || end < myOccupy[row][i].start {

		} else {
			start = myOccupy[row][i].end + 1
			end = start + time - 1
		}
	}
	newEnd = end
	return newEnd
}

func (myOccupy occupyRow) Swap(i, j int) {
	myOccupy[i], myOccupy[j] = myOccupy[j], myOccupy[i]
}

func (myOccupy occupyRow) Len() int {
	return len(myOccupy)
}

func (myOccupy occupyRow) Less(i, j int) bool {
	return myOccupy[i].start < myOccupy[j].start
}

func main() {

	type element struct {
		itemId        int // 工件编号 just 1 2 3 or more
		inputOrder    int // 输入顺序 0 1 2 3 4 5 6...
		eachSerialNum int // 当前工件的工序编号
		machineId     int // 当前工件工序所用机器编号
		time          int // 当前工件工序所用时间
		localStart    int // 当前工件工序的开始时间点
		localEnd      int // 当前工件工序的结束时间点
		last          int // 当前工件工序的上一步工序在输入顺序中的编号
	}

	// process inputs
	var m, n int
	fmt.Scanln(&m, &n)

	var occupyPoint = make([][]occupy, 0)
	for i := 0; i < m; i++ {
		var occupyTemp = make([]occupy, n*m)
		for j := 0; j < n*m; j++ {
			occupyTemp[j].start = 99999999
			occupyTemp[j].end = -1
		}
		occupyPoint = append(occupyPoint, occupyTemp)
	}

	data := scanLine()
	orderTemp := strings.Fields(data)
	var order = make([]element, 0)
	for i := 0; i < n*m; i++ {
		var eleTemp element
		eleTemp.itemId, _ = strconv.Atoi(orderTemp[i])
		eleTemp.inputOrder = i
		order = append(order, eleTemp)
	}

	for i := 1; i <= n; i++ {
		temp := 0
		for j := 0; j < n*m; j++ {
			if order[j].itemId == i {
				temp++
				order[j].eachSerialNum = temp
			}
		}
	}

	var machineOrder = make([][]int, n)
	var processTime = make([][]int, n)
	var matrix = make([][]int, 0)

	reader2 := bufio.NewReader(os.Stdin)
	for i := 0; i < 2*n; i++ {
		data2, _ := reader2.ReadString('\n')
		context2 := string(data2)
		eachRow := make([]int, m)
		rowTemp := strings.Fields(context2)
		for j := 0; j < m; j++ {
			eleTemp, _ := strconv.Atoi(rowTemp[j])
			eachRow[j] = eleTemp
		}
		matrix = append(matrix, eachRow)
		if i >= 0 && i <= n-1 {
			machineOrder[i] = matrix[i]
		} else {
			processTime[i-n] = matrix[i]
		}
	}

	for i := 0; i < n*m; i++ {
		order[i].machineId = machineOrder[order[i].itemId-1][order[i].eachSerialNum-1]
		order[i].time = processTime[order[i].itemId-1][order[i].eachSerialNum-1]
		if order[i].eachSerialNum == 1 {
			order[i].last = -1
		}
	}

	for i := 1; i < n*m; i++ {
		for j := 0; j < i; j++ {
			if order[j].itemId == order[i].itemId && order[j].eachSerialNum == order[i].eachSerialNum-1 {
				order[i].last = order[j].inputOrder
			}
		}
	}

	for i := 0; i < n*m; i++ {
		if i == 0 {
			order[i].localStart = 0
			order[i].localEnd = order[i].localStart + order[i].time - 1
		} else {
			if order[i].last == -1 { // 是当前序号工件的第一道工序
				order[i].localStart = 0
				order[i].localEnd = check(&occupyPoint, order[i].localStart, order[i].time, order[i].machineId-1)
				order[i].localStart = order[i].localEnd - order[i].time + 1
			} else { // 不是当前序号工件的第一道工序
				order[i].localStart = order[order[i].last].localEnd + 1
				order[i].localEnd = check(&occupyPoint, order[i].localStart, order[i].time, order[i].machineId-1)
				order[i].localStart = order[i].localEnd - order[i].time + 1
			}
		}
		occupyPoint[order[i].machineId-1][i].start = order[i].localStart
		occupyPoint[order[i].machineId-1][i].end = order[i].localEnd
		// sort
		var neverMind occupyRow
		neverMind = occupyPoint[order[i].machineId-1]
		sort.Sort(neverMind)
	}

	// print the result
	max := -1
	for i := 0; i < n*m; i++ {
		if order[i].localEnd > max {
			max = order[i].localEnd
		}
	}
	fmt.Println(max + 1)
}
