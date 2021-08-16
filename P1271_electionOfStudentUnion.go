package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type intList []int

func (elect intList) Len() int {
	return len(elect)
}

func (elect intList) Less(a int, b int) bool {
	return elect[a] < elect[b]
}

func (elect intList) Swap(a, b int) {
	elect[a], elect[b] = elect[b], elect[a]
}

func main() {

	var n, m int
	fmt.Scanln(&n, &m)
	// cannot use scanln, or it will be TLE
	reader := bufio.NewReader(os.Stdin)
	// cannot use ReadLine, or it'll be RE, maybe the capacity of list is too low
	data, _ := reader.ReadString('\n')

	var context = make([]string, 0)
	context = strings.Fields(data)

	var elect = make(intList, 0)
	var temp int
	for i := 0; i < len(context); i++ {
		temp, _ = strconv.Atoi(context[i])
		elect = append(elect, temp)
	}

	sort.Sort(elect)
	// sort.Ints(elect)
	for i := 0; i < m; i++ {
		fmt.Printf("%d ", elect[i])
	}

}
