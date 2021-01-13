package main

import (
	"strings"
	"fmt"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	return MyString(strings.ToUpper(string(s))) == s
}

func main() {
	var a MyString = MyString("a")
	fmt.Println(a.IsUpperCase()) // logs false
}
