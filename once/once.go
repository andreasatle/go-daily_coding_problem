// Package once contains Daily Coding Problem #40
//
// This problem was asked by Google.
//
// Given an array of integers where every integer occurs
// three times except for one integer, which only occurs once,
// find and return the non-duplicated integer.
// For example, given [6, 1, 3, 3, 3, 6, 6], return 1.
// Given [13, 19, 13, 13], return 19.
// Do this in O(N) time and O(1) space.
//
// Long remark:
// This is probably a useless unpractical way of solving the problem.
// It does satisfy the criteria setup in the problem-formulation.
// That is  Linear in time, and constant in space (in the array-size).
// The method will be very slow for large numbers, but still meet the criteria.
// A better version would be to sort the array (NlogN) and then identify the
// singlet in one pass, but again, then I violate the O(N) criteria in time.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package once

// Find returns the value of the singlet given one singlet and the rest triplets.
// It would be possible to undo the changes to the array by do the reverse
// ops after the recursion, but then it becomes even more complicated.
func Find(vals []int) int {
	// Continue until all element has vanished
	if maxAbs(vals) == 0 {
		return 0
	}

	// Get the remainder of the singlet
	rem := sum(vals) % 3

	// Subtract all elements with the remainder and divide by 3
	subAndDivideScalar(vals, rem, 3)

	// Return the recursive call
	return 3*Find(vals) + rem
}

// sum returns the sum of the elements in the array
func sum(vals []int) int {
	s := 0
	for _, v := range vals {
		s += v
	}
	return s
}

// subAndDivideScalar subtracts all elements with sub and divides with div
func subAndDivideScalar(vals []int, sub int, div int) {
	for i := range vals {
		vals[i] = (vals[i] - sub) / div
	}
}

// maxAbs returns the largest |element| in the array
func maxAbs(vals []int) int {
	if len(vals) == 0 {
		return 0
	}
	res := vals[0]
	for _, val := range vals {
		if aval := abs(val); aval > res {
			res = aval
		}
	}
	return res
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
