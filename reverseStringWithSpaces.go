/*

   In this Kata, we are going to reverse a string while maintaining the spaces (if any) in their original place.

   For example:

   solve("our code") = "edo cruo"
   -- Normal reversal without spaces is "edocruo". 
   -- However, there is a space at index 3, so the string becomes "edo cruo"

   solve("your code rocks") = "skco redo cruoy". 
   solve("codewars") = "srawedoc"

*/

package main

import (
	"fmt"
	"strings"
)

func Reverse(slc []string) (reversed []string) {
	for i := len(slc) - 1; i >= 0; i-- {
		reversed = append(reversed, slc[i]) 
	}
	
	return
}

func DeSpace(slc []string) []string {
	var noSpaceString []string
	
	for _, v := range slc {
		if v != " " {
			noSpaceString = append(noSpaceString, v) 
		}
	}
	
	return noSpaceString
}

func Solve(s string) (final string) {
	splitString := strings.Split(s, "")
	reversedString := Reverse(splitString)
	reversedChars := DeSpace(reversedString)
	nextVal := 0
	
	for i := 0; i < len(reversedString); i++ {
		if splitString[i] == " " {
			final += " " 
		} else {
			final += reversedChars[nextVal]
			nextVal += 1
		}
	}
	
	return
}

func main() {
	fmt.Println(Solve("your code rocks") == "skco redo cruoy") // should be true
}
