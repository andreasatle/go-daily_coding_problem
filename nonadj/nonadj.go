// Package nonadj contains Daily Coding Problem #9
//
// This problem was asked by Airbnb.
//
// Given a list of integers, write a function that returns the largest
// sum of non-adjacent numbers. Numbers can be 0 or negative.
//
// For example,
//  [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5.
//
//  [5, 1, 1, 5] should return 10, since we pick 5 and 5.
//
// Follow-up: Can you do this in O(N) time and constant space?
//
// Note: Dynamical programming is used to fulfill the complexity
// in the Follow-up.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package nonadj

// Tuple contains an (index,value)-pair
type Tuple struct {
	index int
	sum   int
}

// Sum returns the maximal sum of non-adjacent elements in an array (slice)
func Sum(vals []int) int {
	t0 := Tuple{-3, 0}
	t1 := Tuple{-2, 0}
	t2 := Tuple{-1, 0}
	for i := 0; i < len(vals); i++ {
		if vals[i] <= 0 {
			continue
		}
		if i-t2.index > 1 {
			t0, t1, t2 = t1, t2, Tuple{i, max(t2, t1) + vals[i]}
		} else {
			t0, t1, t2 = t1, t2, Tuple{i, max(t0, t1) + vals[i]}
		}
	}

	return max(t1, t2)
}

// max is an internal function that returns the max sum in two tuples
func max(t1, t2 Tuple) int {
	if t1.sum >= t2.sum {
		return t1.sum
	}
	return t2.sum
}
