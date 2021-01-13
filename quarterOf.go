package main

import (
	"fmt"
)

func QuarterOf(month int) int {
	if month <= 3 {
		return 1
	} else if month <= 6 {
		return 2
	} else if month <= 9 {
		return 3
	} else {
		return 4
	}
}

func main() {
	var month1 int = 3
	var month2 int = 6
	fmt.Println(QuarterOf(month1)) // logs 1
	fmt.Println(QuarterOf(month2)) // logs 2
}
