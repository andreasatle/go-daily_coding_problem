// Package missing contains Daily Coding Problem #4
// This problem was asked by Stripe.
// Given an array of integers, find the first missing positive integer
// in linear time and constant space. In other words, find the lowest
// positive integer that does not exist in the array. The array can contain
// duplicates and negative numbers as well.
// For example, the input [3, 4, -1, 1] should give 2.
// The input [1, 2, 0] should give 3.
// You can modify the input array in-place.
package missing

import "sort"

// FirstMissingPosInt returns the smallest positive integer not in slice.
func FirstMissingPosInt(vals []int) int {
	sort.Ints(vals)
	for i, val := range vals {
		if val < 1 {
			continue
		}
		return FirstMissingNatural(vals[i:])
	}

	// Special case where no positive numbers in array
	return 1
}

// FirstMissingNatural returns the first missing natural number in slice
func FirstMissingNatural(vals []int) int {
	if vals[0] != 1 {
		return 1
	}
	for i := range vals {
		if vals[i] != i+1 {
			return i + 1
		}
	}
	return len(vals) + 1
}
