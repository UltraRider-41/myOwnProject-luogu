package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
	z int
}

type coorList []coordinate

func (myCoor coorList) Swap(i, j int) {
	myCoor[i], myCoor[j] = myCoor[j], myCoor[i]
}

func (myCoor coorList) Less(i, j int) bool {
	return myCoor[i].z < myCoor[j].z
}

func (myCoor coorList) Len() int {
	return len(myCoor)
}

func myPow(a float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else if n == 1 {
		return a
	} else {
		for i := 1; i < n; i++ {
			a = a * a
		}
		return a
	}
}

func countDis(i, j coordinate) float64 {
	return math.Sqrt(myPow(float64(i.x-j.x), 2) + myPow(float64(i.y-j.y), 2) + myPow(float64(i.z-j.z), 2))
}

func main() {

	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	var list coorList
	for i := 0; i < n; i++ {
		var temp coordinate
		context, _ := reader.ReadString('\n')
		data := strings.Fields(context)
		temp.x, _ = strconv.Atoi(data[0])
		temp.y, _ = strconv.Atoi(data[1])
		temp.z, _ = strconv.Atoi(data[2])
		list = append(list, temp)
	}

	sort.Sort(list)

	var distance float64
	for i := 0; i < n-1; i++ {
		distance += countDis(list[i], list[i+1])
	}
	fmt.Printf("%.3f\n", distance)

}
