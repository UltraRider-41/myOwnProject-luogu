package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type funcInput struct {
	a, b, c int64
}

var abcMap = make(map[funcInput]int64)

func w(a, b, c int64) int64 {
	temp := funcInput{a, b, c}
	value, ok := abcMap[temp]
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	} else if a > 20 || b > 20 || c > 20 {
		if ok {
			return value
		} else {
			res := w(20, 20, 20)
			abcMap[temp] = res
			return res
		}
	} else if a < b && b < c {
		if ok {
			return value
		} else {
			res := w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
			abcMap[temp] = res
			return res
		}
	} else {
		if ok {
			return value
		} else {
			res := w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
			abcMap[temp] = res
			return res
		}
	}
}

func main() {

	// process the input
	var list = make([]funcInput, 0)
	reader := bufio.NewReader(os.Stdin)
	var countNum int64
	for i := 0; ; i++ {
		var temp funcInput
		context, _ := reader.ReadString('\n')
		data := strings.Fields(context)
		if data[0] == "-1" && data[1] == "-1" && data[2] == "-1" {
			break
		} else {
			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[1])
			z, _ := strconv.Atoi(data[2])
			temp.a = int64(x)
			temp.b = int64(y)
			temp.c = int64(z)
			list = append(list, temp)
			countNum++
		}
	}

	var i int64
	var res int64
	for i = 0; i < countNum; i++ {
		value, ok := abcMap[list[i]]
		if ok {
			//处理找到的value
			res = value
		} else {
			res = w(list[i].a, list[i].b, list[i].c)
			abcMap[list[i]] = res
		}
		fmt.Printf("w(%d, %d, %d) = %d\n", list[i].a, list[i].b, list[i].c, res)
	}

}
