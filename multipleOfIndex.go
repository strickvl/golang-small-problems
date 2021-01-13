package main

import (
	"fmt"
)

func multipleOfIndex (ints []int) []int {
	var newSlice []int
	for i, v := range ints {
		if i == 0 {
			continue
		} else if v % i == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func main() {
	var ex1 = []int{22, -6, 32, 82, 9, 25}
	fmt.Println(multipleOfIndex(ex1)) // [-6, 32, 25]

	var ex2 = []int{68, -1, 1, -7, 10, 10}
	fmt.Println(multipleOfIndex(ex2)) // [-1, 10]

	var ex3 = []int{-56,-85,72,-26,-14,76,-27,72,35,-21,-67,87,0,21,59,27,-92,68}
	fmt.Println(multipleOfIndex(ex3)) // [-85, 72, 0, 68]	
}
