package main

import (

	"fmt"
	"math"
)

type score struct {
	w int
	l int
}

func main() {

	win11 := 0
	win21 := 0
	lose11 := 0
	lose21 := 0
	N := 62500
	// record 保存输入的字符串
	var record = make([]byte, N)
	var result11 = make([]score, 6000)
	var result21 = make([]score, 6000)
	j := 0 // j 计数，记录result11有多少个元素,即记录打完了几场11分制比赛
	k := 0 // k 计数，记录result21有多少个元素,即记录打完了几场21分制比赛

	// 边输入边处理
	for i := 0; i < N; i++ {
		_, err := fmt.Scanf("%c", &record[i])
		if err != nil {
			return
		}
		if record[i] == 'W' {
			win11++
			win21++
		}
		if record[i] == 'L' {
			lose11++
			lose21++
		}
		if (win11 >= 11 || lose11 >= 11) && math.Abs(float64(win11 - lose11)) >= 2 {
			result11[j] = score{win11, lose11}
			win11 = 0
			lose11 = 0
			j++
		}
		if (win21 >= 21 || lose21 >= 21) && math.Abs(float64(win21 - lose21)) >= 2 {
			result21[k] = score{win21, lose21}
			win21 = 0
			lose21 = 0
			k++
		}
		if record[i] == 'E' {
			result11[j] = score{win11, lose11}
			j++
			result21[k] = score{win21, lose21}
			k++
			break
		}
	}

	// 11分制
	if j == 0 {
		fmt.Println("0:0")
	} else {
		for i := 0; i < j; i++ {
			fmt.Printf("%d:%d\n", result11[i].w, result11[i].l)
		}
	}
	fmt.Println()

	// 21分制
	if k == 0 {
		fmt.Println("0:0")
	} else {
		for i := 0; i < k; i++ {
			fmt.Printf("%d:%d\n", result21[i].w, result21[i].l)
		}
	}

}
