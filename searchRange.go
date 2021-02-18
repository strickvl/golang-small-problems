/*
   Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
   If target is not found in the array, return [-1, -1].

EXAMPLE:

      [5,7,7,8,8,10], target = 8
      returns [3, 4]

THE PROBLEM:

	- binary search only gets us part of the way there
	- we still have to find the left-most copy of that value, and the right-most copy of that value

APPROACH:

   1. find our value in the slice (anywhere / not necessarily the first occurrence of that value)
   2. use that as a way of splitting the slice into two (i.e. as a pivot point)
   3. find the leftmost value in the left slice of our values (using modified binary search)
   4. find the rightmost value in the right slice of our values (using modified binary search)
   5. return the index values of the boundaries (or -1 if they weren't found) in [x, y] format

------------------------------------------------------------------------------------------------------------------------------------
   FIND PIVOT POINT (---- TARGET is 8)

    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
    l     m        h

    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
             l  m  h
                        RETURN 3 because we found our value

------------------------------------------------------------------------------------------------------------------------------------
   FIND LEFT BOUNDS (---- TARGET is 8, pivot is 3)

    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
    l  m     h

   we move low to midpoint + 1 since our midpoint is still less than our target value
    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
          l  h

    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
          l  h
          m

   we move low to midpoint + 1 since our midpoint  was less than the target value
    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
             h
             l
             m

   we move high to midpoint - 1 since the midpoint value was equal to our target value
    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
          h
             l
             m
                        RETURN 3 because we break out of our binary search loop ('low' is greater than 'high')

------------------------------------------------------------------------------------------------------------------------------------
   FIND RIGHT BOUNDS (---- TARGET is 8, pivot is 3)

    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
             l  m  h

   We move the lower value to the midpoint + 1
    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
                   h
                   l
   the midpoint is now also at the same location
    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
                   h
                   l
                   m
    we move the high value to midpoint -1 since the value at the midpoint is greater than our target value
    0  1  2  3  4  5
   [5, 7, 7, 8, 8, 10]
                h
                   l
                   m

                      RETURN 4 because we break out of our binary search loop ('high' is lower than 'low')
------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {
	example1 := []int{5, 7, 7, 8, 8, 10}
	example2 := []int{}
	example3 := []int{0}

	fmt.Println(searchRange(example1, 8))  // logs [3, 4]
	fmt.Println(searchRange(example1, 12)) // logs [-1, -1] because value not found
	fmt.Println(searchRange(example2, 8))  // logs [-1, -1] because empty slice given as input
	fmt.Println(searchRange(example3, 0))  // logs [0, 0] because the single value is both the lower and upper bound
}

func searchRange(nums []int, target int) []int {
	pivot := findPivot(nums, target)

	if pivot == -1 {
		return []int{-1, -1}
	}

	left := findLeftBounds(nums, pivot)
	right := findRightBounds(nums, pivot)

	return []int{left, right}
}

func findPivot(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		midpoint := low + (high-low)/2

		if nums[midpoint] == target {
			return midpoint
		} else if nums[midpoint] < target {
			low = midpoint + 1
		} else if nums[midpoint] > target {
			high = midpoint - 1
		}
	}

	return -1
}

func findLeftBounds(nums []int, pivot int) int {
	low := 0
	high := pivot

	for low <= high {
		midpoint := low + (high-low)/2

		if nums[midpoint] < nums[pivot] {
			low = midpoint + 1
		} else if nums[midpoint] >= nums[pivot] {
			high = midpoint - 1
		}
	}

	return low
}

func findRightBounds(nums []int, pivot int) int {
	low := pivot
	high := len(nums) - 1

	for low <= high {
		midpoint := low + (high-low)/2

		if nums[midpoint] > nums[pivot] {
			high = midpoint - 1
		} else if nums[midpoint] <= nums[pivot] {
			low = midpoint + 1
		}

	}

	return high
}
