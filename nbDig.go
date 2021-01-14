/*
   Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer.
   Square all numbers k (0 <= k <= n) between 0 and n.
   Count the numbers of digits d used in the writing of all the k**2.
   Call nb_dig (or nbDig or ...) the function taking n and d as parameters and returning this count.
*/

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func NbDig(n int, d int) int {
	var count int
	
	for i := 0; i <= n; i++ {
		count += strings.Count(strconv.Itoa(i * i), strconv.Itoa(d))
		
	}

	return count
}

func main() {
	fmt.Println(NbDig(25, 1)) // should log 11
}
