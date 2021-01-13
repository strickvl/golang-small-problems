/*

In this Kata, you will remove the left-most duplicates from a list of integers and return the result.
// Remove the 3's at indices 0 and 3
// followed by removing a 4 at index 1
solve([3, 4, 4, 3, 6, 3]); // => [4, 6, 3]

input: slice
output: slice

rules:
  - remove left most duplicates
  
algorithm:
  - iterate from right to left
    - start from len(slice) - 1
    
  
  
  - for each int
    - check to see if it is in the result slice
      - if is not, then append to result
      
  - find 
    - iterate through result 
    - if int is in result, return true
    - else return false
    
  - reverse slice
    - iterate through slice in reverse
    - append each value to new slice
      

result [3, 6, 4]

*/

package main

import (
  "fmt"
)

func Solve(sliceVals []int) []int{
  uniques := []int{}
  
  for i := len(sliceVals) -1; i >= 0; i-- {
    if !Includes(uniques, sliceVals[i]) {
      uniques = append(uniques,sliceVals[i])
    }
  }
    
  return Reverse(uniques)
}

func Includes(uniques []int, num int) bool{
  for _, v := range uniques {
    if num == v {
      return true
    }
  }
  
  return false
}

func Reverse(slc []int) []int {
  var result []int
  
  for i := len(slc) - 1; i >= 0; i-- {
    result = append(result, slc[i])  
  }
  
  return result
}

func main() {
  fmt.Println(Solve([]int{3,4,4,3,6,3})) // []int{4,6,3}
  fmt.Println(Solve([]int{1,2,1,2,1,2,3})) // []int{1, 2, 3}
} 


