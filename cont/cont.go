// Package cont contains Daily Coding Problem #102
//
// This problem was asked by Lyft.
//
// Given a list of integers and a number K, return which contiguous elements
// of the list sum to K.
//
// For example, if the list is [1, 2, 3, 4, 5] and K is 9, then it should return
// [2, 3, 4], since 2 + 3 + 4 = 9.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package cont

// FindTerms returns a contiguous subslice that sums to sum
func FindTerms(vals []int, sum int) []int {
	if len(vals) == 0 {
		return []int{}
	}

	res, ok := findSum(vals, sum)

	if ok {
		return res
	}

	return FindTerms(vals[1:], sum)
}

func findSum(vals []int, key int) ([]int, bool) {
	sum := 0
	for i, val := range vals {
		sum += val
		if sum == key {
			return vals[:i+1], true
		}
	}
	return []int{}, false
}
