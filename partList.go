/*

   Write a function partlist that gives all the ways to divide a list (an array) of at least two elements into two non-empty parts.

   Each two non empty parts will be in a pair (or an array for languages without tuples or a structin C - C: see Examples test Cases - )
   Each part will be in a string
   Elements of a pair must be in the same order as in the original array.
   Examples of returns in different languages:
   a = ["az", "toto", "picaro", "zone", "kiwi"] -->

   [["az", "toto picaro zone kiwi"], ["az toto", "picaro zone kiwi"], ["az toto picaro", "zone kiwi"], ["az toto picaro zone", "kiwi"]]

   input: slice of string
   output: string (each part has () around it), all concatenated into a single string

   - initialise an empty string (finalStr)

   - go from 0 up to len(slice) - 2
  each time take a slice from 0: the number (part1)
  take a slice from the number : the end (part2)
  
  add it to finalStr in the format
  "(%v, %v)", part1, part2


  return finalStr
*/

package main

import (
	"fmt"
	"strings"
)

func PartList(slc []string) string {
	var finalStr string
	
	for i := 0; i <= len(slc) - 2; i++ {
		var part1 string
		var part2 string
		
		part1 = strings.Join(slc[0: i + 1], " ")
		part2 = strings.Join(slc[i + 1:], " ")
		
		finalStr += fmt.Sprintf("(%v, %v)", part1, part2)
	}
	
	return finalStr
}

func main() {
	fmt.Println(PartList([]string{"I", "wish", "I", "hadn't", "come"}) == "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)")
}
