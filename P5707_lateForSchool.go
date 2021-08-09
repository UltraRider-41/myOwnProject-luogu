package main

import "fmt"

func main() {

	var s, v int
	fmt.Scanln(&s, &v)

	var t, h, m int
	var temp int
	m = 480
	t = s / v
	if s % v != 0 {
		t++
	}
	t += 10

	if t <= 480 {
		h = (m - t) / 60
		m -= h * 60 + t
	} else if t > 480 {
		temp = (t - 480) / 60
		h = 24 - temp
		if (t - 480) % 60 != 0 {
			h -= 1
		}
		m = t - m
		m -= temp * 60
		m = (60 - m) % 60
	}

	fmt.Printf("%02d:%02d", h, m)
}
