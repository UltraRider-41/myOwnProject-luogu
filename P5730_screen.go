package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)

	var list string
	fmt.Scanln(&list)

	type num struct {
		a string
		b string
		c string
		d string
		e string
	}
	var screenNum = make([]num, 0)
	var temp num
	for i := 0; i <= 9; i++ {
		switch i {
			case 0:
				temp.a = "XXX"
				temp.b = "X.X"
				temp.c = "X.X"
				temp.d = "X.X"
				temp.e = "XXX"
			case 1:
				temp.a = "..X"
				temp.b = "..X"
				temp.c = "..X"
				temp.d = "..X"
				temp.e = "..X"
			case 2:
				temp.a = "XXX"
				temp.b = "..X"
				temp.c = "XXX"
				temp.d = "X.."
				temp.e = "XXX"
			case 3:
				temp.a = "XXX"
				temp.b = "..X"
				temp.c = "XXX"
				temp.d = "..X"
				temp.e = "XXX"
			case 4:
				temp.a = "X.X"
				temp.b = "X.X"
				temp.c = "XXX"
				temp.d = "..X"
				temp.e = "..X"
			case 5:
				temp.a = "XXX"
				temp.b = "X.."
				temp.c = "XXX"
				temp.d = "..X"
				temp.e = "XXX"
			case 6:
				temp.a = "XXX"
				temp.b = "X.."
				temp.c = "XXX"
				temp.d = "X.X"
				temp.e = "XXX"
			case 7:
				temp.a = "XXX"
				temp.b = "..X"
				temp.c = "..X"
				temp.d = "..X"
				temp.e = "..X"
			case 8:
				temp.a = "XXX"
				temp.b = "X.X"
				temp.c = "XXX"
				temp.d = "X.X"
				temp.e = "XXX"
			case 9:
				temp.a = "XXX"
				temp.b = "X.X"
				temp.c = "XXX"
				temp.d = "..X"
				temp.e = "XXX"
		}
		screenNum = append(screenNum, temp)
	}
	var aList = make([]string, 0)
	var bList = make([]string, 0)
	var cList = make([]string, 0)
	var dList = make([]string, 0)
	var eList = make([]string, 0)
	for i := 0; i <= 9; i++ {
		aList = append(aList, screenNum[i].a)
		bList = append(bList, screenNum[i].b)
		cList = append(cList, screenNum[i].c)
		dList = append(dList, screenNum[i].d)
		eList = append(eList, screenNum[i].e)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < n; j++ {
			if j != n-1 {
				switch i {
				case 0:
					fmt.Printf("%s.", aList[list[j]-48])
				case 1:
					fmt.Printf("%s.", bList[list[j]-48])
				case 2:
					fmt.Printf("%s.", cList[list[j]-48])
				case 3:
					fmt.Printf("%s.", dList[list[j]-48])
				case 4:
					fmt.Printf("%s.", eList[list[j]-48])
				}
			} else {
				switch i {
				case 0:
					fmt.Printf("%s", aList[list[j]-48])
				case 1:
					fmt.Printf("%s", bList[list[j]-48])
				case 2:
					fmt.Printf("%s", cList[list[j]-48])
				case 3:
					fmt.Printf("%s", dList[list[j]-48])
				case 4:
					fmt.Printf("%s", eList[list[j]-48])
				}
			}
		}
		fmt.Println()
	}
}
