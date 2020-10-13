// Package jumps contains Daily Coding Problem #106
//
// This problem was asked by Pinterest.
//
// Given an integer list where each number represents the number of hops you
// can make, determine whether you can reach to the last index starting at
// index 0.
//
// For example, [2, 0, 1, 0] returns True while [1, 1, 0, 1] returns False.
//
// Remark: We don't look for cyclic paths, but will terminate after len(array) steps.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package jumps

// FirstToLast returns true if we can jump from first to last index
// given the offset in the slice.
func FirstToLast(offsets []int) bool {
	last := len(offsets) - 1

	for i, cnt := 0, 0; cnt <= last; i, cnt = i+offsets[i], cnt+1 {
		// Return false if index is out of range
		if i < 0 || i > last {
			return false
		}

		// Return true if last index is found
		if i == last {
			return true
		}
	}
	return false
}
