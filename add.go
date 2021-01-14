/*
   Create a function add(n)/Add(n) which returns a function that always adds n to any number
*/

package main

import (
	"fmt"
)

func Add(addVal int) func(int)int {
	return func(val int) int {
		return val + addVal
	}
}

func main() {
	fmt.Println(Add(1)(3) == 4) // should be true
}
