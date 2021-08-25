// Junk problem description!!! Waste so much time

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type gangMem struct {
	name     string
	position string
	contr    int
	lv       int
	rank     int
}

// condition 1 means sort contr, 2 means sort lv
// sort contribution first, if two members own same position, then sort their levels; if two members also own same
// level, then sort their rank
func insertionSortGang(list *[]gangMem, n int, condition int) {

	myList := *list
	if condition == 1 {
		for i := 1; i < n; i++ {
			j := i
			for j > 0 {
				if myList[j].contr > myList[j-1].contr {
					myList[j], myList[j-1] = myList[j-1], myList[j]
					j--
				} else {
					break
				}
			}
		}
		for i := 1; i < n; i++ {
			j := i
			for j > 0 {
				if myList[j].contr == myList[j-1].contr && myList[j].rank < myList[j-1].rank {
					myList[j], myList[j-1] = myList[j-1], myList[j]
					j--
				} else {
					break
				}
			}
		}
	} else if condition == 2 {
		var start, end int
		start = 0
		for i := 0; i < n; i++ {
			end = i
			if myList[start].position != myList[end].position {
				samePosList := myList[start:end]
				for x := 1; x < len(samePosList); x++ {
					j := x
					for j > 0 {
						if samePosList[j].lv > samePosList[j-1].lv {
							samePosList[j], samePosList[j-1] = samePosList[j-1], samePosList[j]
							j--
						} else {
							break
						}
					}
				}
				for x := 1; x < len(samePosList); x++ {
					j := x
					for j > 0 {
						if samePosList[j].lv == samePosList[j-1].lv && samePosList[j].rank < samePosList[j-1].rank {
							samePosList[j], samePosList[j-1] = samePosList[j-1], samePosList[j]
							j--
						} else {
							break
						}
					}
				}
				start = end
				// till the end
			} else if end == n-1 && myList[start].position == myList[end].position {
				end += 1
				samePosList := myList[start:end]
				for x := 1; x < len(samePosList); x++ {
					j := x
					for j > 0 {
						if samePosList[j].lv > samePosList[j-1].lv {
							samePosList[j], samePosList[j-1] = samePosList[j-1], samePosList[j]
							j--
						} else {
							break
						}
					}
				}
				for x := 1; x < len(samePosList); x++ {
					j := x
					for j > 0 {
						if samePosList[j].lv == samePosList[j-1].lv && samePosList[j].rank < samePosList[j-1].rank {
							samePosList[j], samePosList[j-1] = samePosList[j-1], samePosList[j]
							j--
						} else {
							break
						}
					}
				}
				start = end
			}
		}
	}
}

func main() {

	var n int
	fmt.Scanln(&n)

	// process the input
	var memList = make([]gangMem, 0)
	var temp gangMem
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		data, _, _ := reader.ReadLine()
		context := strings.Fields(string(data))
		temp.name = context[0]
		temp.position = context[1]
		temp.contr, _ = strconv.Atoi(context[2])
		temp.lv, _ = strconv.Atoi(context[3])
		temp.rank = i
		memList = append(memList, temp)
	}

	var bangZhuList = make([]gangMem, 3)
	var otherMemList = make([]gangMem, 0)
	/* it seems we needn't sort BangZhu and FuBangZhu
	if bangZhuList[1].lv < bangZhuList[2].lv {
		bangZhuList[1], bangZhuList[2] = bangZhuList[2], bangZhuList[1]
	}
	*/

	insertionSortGang(&memList, n, 1)

	countBangzhu := 0
	countFu := 0
	// count BangZhu & FuBangZhu numbers
	for i := 0; i < n; i++ {
		if memList[i].position == "BangZhu" {
			countBangzhu++
			bangZhuList[0] = memList[i]
		} else if memList[i].position == "FuBangZhu" && countFu != 1 {
			countFu++
			bangZhuList[1] = memList[i]
		} else if memList[i].position == "FuBangZhu" && countFu == 1 {
			countFu++
			bangZhuList[2] = memList[i]
		}
		// all found, break loop
		if countBangzhu == 1 && countFu == 2 {
			break
		}
	}

	// deal different situations of BangZhu & FuBangZhu numbers
	if countBangzhu == 1 && countFu == 0 {
		bangZhuList[1] = memList[0]
		bangZhuList[2] = memList[1]
	} else if countBangzhu == 1 && countFu == 1 {
		bangZhuList[2] = memList[0]
	} else if countBangzhu == 0 && countFu == 0 {
		bangZhuList[0] = memList[0]
		bangZhuList[1] = memList[1]
		bangZhuList[2] = memList[2]
	} else if countBangzhu == 0 && countFu == 1 {
		bangZhuList[0] = memList[0]
		bangZhuList[2] = memList[1]
	} else if countBangzhu == 0 && countFu == 2 {
		bangZhuList[0] = memList[0]
	}
	// adjust the order of FuBangZhu based on rank element
	if bangZhuList[1].rank > bangZhuList[2].rank {
		bangZhuList[1], bangZhuList[2] = bangZhuList[2], bangZhuList[1]
	}

	for i := 0; i < n; i++ {
		if memList[i].name != bangZhuList[0].name && memList[i].name != bangZhuList[1].name && memList[i].name != bangZhuList[2].name {
			otherMemList = append(otherMemList, memList[i])
		}
	}
	// assigned position
	for i := 0; i < len(otherMemList); i++ {
		if i+1 == 1 || i+1 == 2 {
			otherMemList[i].position = "HuFa"
		} else if i+1 >= 3 && i+1 <= 6 {
			otherMemList[i].position = "ZhangLao"
		} else if i+1 >= 7 && i+1 <= 13 {
			otherMemList[i].position = "TangZhu"
		} else if i+1 >= 14 && i+1 <= 38 {
			otherMemList[i].position = "JingYing"
		} else if i+1 >= 39 {
			otherMemList[i].position = "BangZhong"
		}
	}

	insertionSortGang(&otherMemList, len(otherMemList), 2)

	// print results
	for i := 0; i < 3; i++ {
		if i != 2 {
			fmt.Printf("%v %v %v\n", bangZhuList[i].name, bangZhuList[i].position, bangZhuList[i].lv)
		} else {
			fmt.Printf("%v %v %v", bangZhuList[i].name, bangZhuList[i].position, bangZhuList[i].lv)
		}
	}
	for i := 0; i < n-3; i++ {
		if i == 0 && i != n-4 {
			fmt.Printf("\n%v %v %v\n", otherMemList[i].name, otherMemList[i].position, otherMemList[i].lv)
		} else if i != 0 && i != n-4 {
			fmt.Printf("%v %v %v\n", otherMemList[i].name, otherMemList[i].position, otherMemList[i].lv)
		} else if i != 0 && i == n-4 {
			fmt.Printf("%v %v %v", otherMemList[i].name, otherMemList[i].position, otherMemList[i].lv)
		} else if i == 0 && i == n-4 {
			fmt.Printf("\n%v %v %v", otherMemList[i].name, otherMemList[i].position, otherMemList[i].lv)
		}
	}
}
