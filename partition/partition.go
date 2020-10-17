// Package partition contains Daily Coding Problem #60
//
// This problem was asked by Facebook.
//
// Given a multiset of integers, return whether it can be partitioned into
// two subsets whose sums are the same.
//
// For example, given the multiset {15, 5, 20, 10, 35, 15, 10}, it would
// return true, since we can split it up into {15, 5, 10, 15, 10} and
// {20, 35}, which both add up to 55.
//
// Given the multiset {15, 5, 20, 10, 35}, it would return false, since we
// can't split it up into two subsets that add up to the same sum.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package partition

// SameSum returns true if the values can be partitioned in two subsets
// with the same sum.
func SameSum(vals []int) bool {
	sum := intSum(vals)

	// Partition not possible for odd sum
	if sum%2 == 1 {
		return false
	}

	// sum is even, from here on it is the half sum
	sum /= 2

	return subsum(vals, sum)
}

func intSum(vals []int) int {
	acc := 0
	for _, val := range vals {
		acc += val
	}
	return acc
}

func subsum(vals []int, sum int) bool {
	if sum == 0 {
		return true
	}
	if len(vals) == 0 {
		return false
	}
	return subsum(vals[1:], sum-vals[0]) || subsum(vals[1:], sum)
}
