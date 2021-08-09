package main

import (
	"fmt"
	"strconv"
)

func main()  {

	var n, m int
	fmt.Scanln(&n, &m)

	var mine = make([][]string, 0)
	for i := 0; i < n; i++ {
		var temp = make([]string, m)
		var s string
		fmt.Scanln(&s)
		for j := 0; j < m; j++ {
			temp[j] = string(s[j])
		}
		mine = append(mine, temp)
	}
	// fmt.Println(mine)

	type point struct {
		row int
		col int
	}

	for i := 0; i < n; i++ {
		for j  := 0; j < m; j++ {
			var num int
			if mine[i][j] != "*" {
				var lu, u, ru, l, r, ld, d, rd point // left up right down
				lu.row = i-1
				lu.col = j-1

				u.row = i-1
				u.col = j

				ru.row = i-1
				ru.col = j+1

				l.row = i
				l.col = j-1

				r.row = i
				r.col = j+1

				ld.row = i+1
				ld.col = j-1

				d.row = i+1
				d.col = j

				rd.row = i+1
				rd.col = j+1

				if lu.row >= 0 && lu.col >= 0 && lu.row < n && lu.col < m {
					if mine[lu.row][lu.col] == "*" {
						num++
					}
				}
				if u.row >= 0 && u.col >= 0 && u.row < n && u.col < m {
					if mine[u.row][u.col] == "*" {
						num++
					}
				}
				if ru.row >= 0 && ru.col >= 0 && ru.row < n && ru.col < m {
					if mine[ru.row][ru.col] == "*" {
						num++
					}
				}
				if l.row >= 0 && l.col >= 0 && l.row < n && l.col < m {
					if mine[l.row][l.col] == "*" {
						num++
					}
				}
				if r.row >= 0 && r.col >= 0 && r.row < n && r.col < m {
					if mine[r.row][r.col] == "*" {
						num++
					}
				}
				if ld.row >= 0 && ld.col >= 0 && ld.row < n && ld.col < m {
					if mine[ld.row][ld.col] == "*" {
						num++
					}
				}
				if d.row >= 0 && d.col >= 0 && d.row < n && d.col < m {
					if mine[d.row][d.col] == "*" {
						num++
					}
				}
				if rd.row >= 0 && rd.col >= 0 && rd.row < n && rd.col < m {
					if mine[rd.row][rd.col] == "*" {
						num++
					}
				}
				mine[i][j] = strconv.Itoa(num)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (j+1) % m == 0 {
				fmt.Printf("%s\n", mine[i][j])
			} else {
				fmt.Printf("%s", mine[i][j])
			}
		}
	}
}
