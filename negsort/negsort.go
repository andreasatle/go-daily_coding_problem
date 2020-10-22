// Package negsort contains Daily Coding Problem #118
//
// This problem was asked by Google.
//
// Given a sorted list of integers,
// square the elements and give the output in sorted order.
//
// Note: The trivial solution is to just square, and re-sort
// the full slice. This becomes an O(N*LogN) algorithm.
// I sense that Google wants us to do something more clever,
// but we could call that premature optimization;)
//
// Note: The more clever algorithm could include reversing the negative
// part of the slice, and then merge the two parts, as for merge-sort.
// I have never implemented a merge-sort without extra storage,
// but know that it is possible.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package negsort

import (
	"sort"
)

// SortSquares squares and re-sorts the values in the slice
func SortSquares(vals []int) {
	for i := range vals {
		vals[i] = vals[i] * vals[i]
	}
	sort.Ints(vals)
}
