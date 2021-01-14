/*
   Given a string and an array of integers representing indices, capitalize all letters at the given indices.

   For example:

   capitalize("abcdef",[1,2,5]) = "aBCdeF"
   capitalize("abcdef",[1,2,5,100]) = "aBCdeF". There is no index 100.
   The input will be a lowercase string with no spaces and an array of digits.

*/

package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string, arr []int) string {
	var sliceSt = strings.Split(st, "")
	
	for _, v := range arr {
		
		if v < len(st) {
			sliceSt[v] = strings.ToUpper(sliceSt[v])
		}
		
	}
	
	return strings.Join(sliceSt, "")
}

func main() {
	fmt.Println(Capitalize("abcdef",[]int{1,2,5, 100}) == "aBCdeF") // should be true
}
