// Package exprod contains Daily Coding Problem #2
//
// This problem was asked by Uber.
//
// Given an array of integers, return a new array such that each element
// at index i of the new array is the product of all the numbers in the
// original array except the one at i.
//
// For example, if our input was
//  [1, 2, 3, 4, 5],
// the expected output would be
//  [120, 60, 40, 30, 24].
// If our input was
//  [3, 2, 1],
// the expected output would be
//  [2, 3, 6].
//
// Follow-up: what if you can't use division?
//
// Note: This problem is solved with a divide and conquer algorithm.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package exprod

func Product(vals []int) []int {
	prods := make([]int, len(vals))
	product(vals, prods, 0, len(vals))
	return prods
}

func product(vals, prods []int, first, last int) int {
	if first+1 == last {
		prods[first] = 1
		return vals[first]
	}
	mid := (first + last + 1) / 2
	prod1 := product(vals, prods, first, mid)
	prod2 := product(vals, prods, mid, last)
	for i := first; i < mid; i++ {
		prods[i] *= prod2
	}
	for i := mid; i < last; i++ {
		prods[i] *= prod1
	}
	return prod1 * prod2
}
