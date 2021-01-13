package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	var text []string = strings.Split(strings.ToLower(str), "")
	var reverseText []string

	for i := len(text) - 1; i >= 0; i-- {
		reverseText = append(reverseText, text[i])
	}

	return strings.Join(text, "") == strings.Join(reverseText, "")
}

func main() {
	var text string = "Yay"
	fmt.Println(IsPalindrome(text))
}
