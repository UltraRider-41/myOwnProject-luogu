package main

import "fmt"

func main() {

	var cola, average float64
	var classmate, bottle int

	fmt.Scanln(&cola, &classmate)

	bottle = classmate * 2
	average = cola / float64(classmate)

	fmt.Printf("%.3f\n%d", average, bottle)
}
