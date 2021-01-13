package main

import (
	"fmt"
)

func Xor(a, b bool) bool {
  if (a == false && b == true) || (a == true && b == false) {
    return true
  }
  return false
}

func main() {
	var x bool = false
	var y bool = true
	fmt.Println(Xor(x, y) == true) // logs true
}
