package main

import (
	"fmt"
)

func main() {

	var s string
	fmt.Scanln(&s)
	// fmt.Printf("%T %v\n", s[3], s[3])

	var aRes, bRes, cRes int
	for i := 0; i < len(s); i++ {
		// notice brackets!!!
		if (i+1) % 5 == 1 {
			switch string(s[i]) {
				case "a" :
					if s[i+3] >= 48 && s[i+3] <= 57 {
						aRes = int(s[i+3])-48
					} else if string(s[i+3]) == "b" {
						aRes = bRes
					} else if string(s[i+3]) == "c" {
						aRes = cRes
					}
				case "b" :
					if s[i+3] >= 48 && s[i+3] <= 57 {
						bRes = int(s[i+3])-48
					} else if string(s[i+3]) == "a" {
						bRes = aRes
					} else if string(s[i+3]) == "c" {
						bRes = cRes
					}
				case "c" :
					if s[i+3] >= 48 && s[i+3] <= 57 {
						cRes = int(s[i+3])-48
					} else if string(s[i+3]) == "b" {
						cRes = bRes
					} else if string(s[i+3]) == "a" {
						cRes = aRes
					}
			}
		}
	}

	fmt.Println(aRes, bRes, cRes)
}
