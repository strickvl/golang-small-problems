/*

   In this Kata, you will be given an array of integers whose elements have both a negative and a positive value, except for one integer that is either only negative or only positive. Your task will be to find that integer.

Examples:

   [1, -1, 2, -2, 3] => 3

   3 has no matching negative appearance

*/

package main

import (
	"fmt"
	"sort"
)

func Contains(arr []int, item int) bool {
	for _, v := range arr {
		if item == v {
			return true 
		}
	}
	
	return false
}

func Solve(arr []int) int {
	sort.Ints(arr)
	
	for _, v := range arr {
		if Contains(arr, -v) == false {
			return v 
		}
	}
	
	return 0
}

func main() {
	fmt.Println(Solve([]int{1,-1,2,-2,3}) == 3) // should be true
}
