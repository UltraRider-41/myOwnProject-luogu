package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	var code string
	code = string(data)
	var context = make([]string, 0)
	context = strings.Fields(code)

	var temp int
	var list = make([]int, 0)
	for i := 0; i < len(context); i++ {
		temp, _ = strconv.Atoi(context[i])
		list = append(list, temp)
	}

	n := list[0]
	var res = make([]int, 0)
	for i := 1; i < len(list); i++ {
		if i % 2 != 0 {
			for j := 0; j < list[i]; j++ {
				res = append(res, 0)
			}
		} else if i % 2 == 0 {
			for j := 0; j < list[i]; j++ {
				res = append(res, 1)
			}
		}
	}

	for i := 0; i < len(res); i++ {
		if (i+1) % n == 0 {
			fmt.Printf("%d\n", res[i])
		} else if (i+1) % n != 0 {
			fmt.Printf("%d", res[i])
		}
	}
}