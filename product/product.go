// Package product contains Daily Coding Problem #69
// This problem was asked by Facebook.
// Given a list of integers, return the largest product that
// can be made by multiplying any three integers.
// For example, if the list is [-10, -10, 5, 2], we should return 500,
// since that's -10 * -10 * 5.
// You can assume the list has at least three integers.
package product

import (
	"log"
	"sort"
)

func LargestProductOfThree(u []int) int {
	n := len(u)
	v := make([]int, n)
	copy(v, u)

	// Check that we have atleast 3 ints
	if n < 3 {
		log.Fatalf("Need atleast 3 integers, given: %v", n)
	}

	// Sort the slice of ints
	sort.Ints(v)

	// Two cases: either the largest int is non-negative or not
	if v[n-1] >= 0 {
		return v[n-1] * max(v[0]*v[1], v[n-3]*v[n-2])
	}

	return v[n-3] * v[n-2] * v[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
