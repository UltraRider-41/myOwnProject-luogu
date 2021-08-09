package main

import "fmt"

func main() {

	type point struct {
		x int
		y int
		z int
		flag int // here 0e means yes, 1 means no
	}

	var w, x ,h int
	var volume int
	fmt.Scanln(&w, &x, &h)
	volume = w * x * h

	var q int
	fmt.Scanln(&q)

	var pointList1 = make([]point, 0)
	var pointList2 = make([]point, 0)
	var wholePointList = make([]point, 0)
	var temp1, temp2, temp point
	for i := 0; i < q; i++ {
		fmt.Scanln(&temp1.x, &temp1.y, &temp1.z, &temp2.x, &temp2.y, &temp2.z)
		pointList1 = append(pointList1, temp1)
		pointList2 = append(pointList2, temp2)
	}

	for a := 1; a <= w; a++ {
		for b := 1; b <= x; b++ {
			for c := 1; c <= h; c++ {
				temp.x = a
				temp.y = b
				temp.z = c
				wholePointList = append(wholePointList, temp)
			}
		}
	}

	for i := 0; i < q; i++ {
		for a := pointList1[i].x; a <= pointList2[i].x; a++ {
			for b := pointList1[i].y; b <= pointList2[i].y; b++{
				for c := pointList1[i].z; c <= pointList2[i].z; c++{
					for j := 0; j < len(wholePointList); j++ {
						if wholePointList[j].x == a && wholePointList[j].y == b && wholePointList[j].z == c && wholePointList[j].flag == 0 {
							wholePointList[j].flag = 1
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(wholePointList); i++ {
		if wholePointList[i].flag == 1 {
			volume--
		}
	}
	fmt.Printf("%d", volume)
}
