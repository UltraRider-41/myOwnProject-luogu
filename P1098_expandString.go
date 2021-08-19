package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var p1, p2, p3 int
	fmt.Scanln(&p1, &p2, &p3)

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	for i := 0; i < len(data); i++ {
		if data[i] == '-' && i != 0 && i != len(data)-1 && ((data[i-1] >= 48 && data[i-1] <= 57 && data[i+1] >= 48 && data[i+1] <= 57) || (+data[i-1] >= 97 && data[i-1] <= 122 && data[i+1] >= 97 && data[i+1] <= 122)) {
			if data[i-1]+1 == data[i+1] {

			} else if data[i-1] >= data[i+1] {
				fmt.Printf("-")
			} else if data[i-1] < data[i+1] {
				var temp = int(data[i+1] - data[i-1])
				switch p3 {
				case 1:
					switch p1 {
					case 1:
						for a := 0; a < temp-1; a++ {
							for b := 0; b < p2; b++ {
								fmt.Printf("%s", string(int(data[i-1])+a+1))
							}
						}
					case 2:
						for a := 0; a < temp-1; a++ {
							for b := 0; b < p2; b++ {
								if data[i-1] >= 48 && data[i-1] <= 57 {
									fmt.Printf("%s", string(int(data[i-1])+a+1))
								} else if data[i-1] >= 97 && data[i-1] <= 122 {
									fmt.Printf("%s", string(int(data[i-1])+a+1-32))
								}
							}
						}
					case 3:
						for a := 0; a < temp-1; a++ {
							for b := 0; b < p2; b++ {
								fmt.Printf("*")
							}
						}
					}
				case 2:
					switch p1 {
					case 1:
						for a := 0; a < temp-1; a++ {
							for b := 0; b < p2; b++ {
								fmt.Printf("%s", string(int(data[i+1])-a-1))
							}
						}
					case 2:
						for a := 0; a < temp-1; a++ {
							for b := 0; b < p2; b++ {
								if data[i-1] >= 48 && data[i-1] <= 57 {
									fmt.Printf("%s", string(int(data[i+1])-a-1))
								} else if data[i-1] >= 97 && data[i-1] <= 122 {
									fmt.Printf("%s", string(int(data[i+1])-a-1-32))
								}
							}
						}
					case 3:
						for a := 0; a < temp-1; a++ {
							for b := 0; b < p2; b++ {
								fmt.Printf("*")
							}
						}
					}
				}
			}
		} else {
			fmt.Printf("%s", string(data[i]))
		}
	}

}
