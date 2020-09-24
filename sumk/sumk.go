// Package sumk contains Daily Coding Problem #1
//
// This problem was recently asked by Google.
//
// Given a list of numbers and a number k, return
// whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17,
// return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?
//
// Author: Andreas Atle, atle.andreas@gmail.com
package sumk

type empty = struct{}

// Sumk returns true if two values sum to k.
func SumK(vals []int, k int) bool {
	set := map[int]empty{}
	for _, val := range vals {
		if _, ok := set[k-val]; ok {
			return true
		}
		set[val] = empty{}
	}
	return false
}
