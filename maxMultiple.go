/*
   Given a Divisor and a Bound , Find the largest integer N , Such That ,

   Conditions :
   N is divisible by divisor

   N is less than or equal to bound

   N is greater than 0.

   Notes
   The parameters (divisor, bound) passed to the function are only positive values .
   It's guaranteed that a divisor is Found .

   start at one minus the bound and iterate down to the divisor + 1
   if that value (of the current iteration) % divisor == 0, then return that value
*/

package main

import "fmt"

func MaxMultiple(d, b int) int {
	for i := b; i > d; i -- {
		if i % d == 0 {
			return i 
		}
	}
	return d
}

func main() {
	fmt.Println(MaxMultiple(2, 7) == 6)
}
