package main

import (
	"fmt"
)

func Divide(weight int) bool {
	var remainder int = weight % 2

	if remainder != 0 || weight == 2 {
		return false
	} else {
		return true
	}
}


func main() {
	var box1 int = 1
	var box2 int = 4

	fmt.Println(Divide(box1)) // logs false
	fmt.Println(Divide(box2)) // logs true
}
