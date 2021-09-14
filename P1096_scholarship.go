package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scholarScore struct {
	chinese int
	math    int
	english int
	total   int
	no      int
}

// flag = 1:total; = 2:chinese; =3:no
func scholarSort(myList []scholarScore, l int, r int, flag int) {

	if flag == 1 {
		if l < r {

			i := l
			j := r
			pivot := myList[i]

			for i < j {
				for i < j && myList[j].total < pivot.total {
					j--
				}
				if i < j {
					myList[i] = myList[j]
					i++
				}
				for i < j && myList[i].total > pivot.total {
					i++
				}
				if i < j {
					myList[j] = myList[i]
					j--
				}
			}
			myList[i] = pivot
			scholarSort(myList, l, i-1, 1)
			scholarSort(myList, i+1, r, 1)
		}
	} else if flag == 2 {
		if l < r {

			i := l
			j := r
			pivot := myList[i]

			for i < j {
				for i < j && myList[j].chinese < pivot.chinese {
					j--
				}
				if i < j {
					myList[i] = myList[j]
					i++
				}
				for i < j && myList[i].chinese > pivot.chinese {
					i++
				}
				if i < j {
					myList[j] = myList[i]
					j--
				}
			}
			myList[i] = pivot
			scholarSort(myList, l, i-1, 2)
			scholarSort(myList, i+1, r, 2)
		}
	} else if flag == 3 {
		if l < r {

			i := l
			j := r
			pivot := myList[i]

			for i < j {
				for i < j && myList[j].no > pivot.no {
					j--
				}
				if i < j {
					myList[i] = myList[j]
					i++
				}
				for i < j && myList[i].no < pivot.no {
					i++
				}
				if i < j {
					myList[j] = myList[i]
					j--
				}
			}
			myList[i] = pivot
			scholarSort(myList, l, i-1, 3)
			scholarSort(myList, i+1, r, 3)
		}
	} else {
		panic("Wrong Flag!")
	}

}

func main() {

	var n int
	fmt.Scanln(&n)

	// process the input
	var list = make([]scholarScore, 0)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		var temp scholarScore
		context, _ := reader.ReadString('\n')
		data := strings.Fields(context)
		temp.chinese, _ = strconv.Atoi(data[0])
		temp.math, _ = strconv.Atoi(data[1])
		temp.english, _ = strconv.Atoi(data[2])
		temp.total = temp.chinese + temp.math + temp.english
		temp.no = i + 1
		list = append(list, temp)
	}

	// total score(highest)
	l := 0
	r := n
	scholarSort(list, l, r-1, 1)
	// chinese score(highest)
	for i := 0; i < n; {
		l = i
		r = i + 1
		for j := i + 1; j < n; j++ {
			if list[i].total == list[j-1].total && list[i].total != list[j].total {
				r = j
				if j-1 != i {
					scholarSort(list, l, r-1, 2)
				}
			}
		}
		i = r
	}
	// student No(lowest)
	for i := 0; i < n; {
		l = i
		r = i + 1
		for j := i + 1; j < n; j++ {
			if list[i].total == list[j-1].total && list[i].chinese == list[j-1].chinese && (list[i].total != list[j].total || list[i].chinese != list[j].chinese) {
				r = j
				if j-1 != i {
					scholarSort(list, l, r-1, 3)
				}
			}
		}
		i = r
	}

	// print the output
	for i := 0; i < 5; i++ {
		fmt.Printf("%d %d\n", list[i].no, list[i].total)
	}

}
