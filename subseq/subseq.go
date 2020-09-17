// Daily Coding Problem #75
// This problem was asked by Microsoft.
// Given an array of numbers, find the length of the longest increasing subsequence
// in the array. The subsequence does not necessarily have to be contiguous.
// For example, given the array
// [0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15],
// the longest increasing subsequence has length 6: it is 0, 2, 6, 9, 11, 15.

package subseq

// Compute the longest subsequence recursively.
// The first recursive call does not include the first entry in the array
// The second recursive call includes the first entry in the subsequence if
// it is larger than the last entry in the current subsequence.
// The subsequence is copied before each recursion; this could likely be optimized.
func Subsequence(array, subseq []int) []int {
	if len(array) == 0 {
		return subseq
	}

	var subseq1, subseq2 []int

	// make a copy of subseq
	subseq1 = make([]int, len(subseq))
	copy(subseq1, subseq)

	// Check if a second recursion is needed where array[0] is appended to subseq
	flagSecondRecursion := len(subseq) == 0 || array[0] > subseq[len(subseq)-1]

	// Do recursion with the rest of the array
	if len(array) > 1 {
		subseq1 = Subsequence(array[1:], subseq1)

		// Avoid calling two identical recursions
		if flagSecondRecursion {
			subseq2 = make([]int, len(subseq))
			copy(subseq2, subseq)
			subseq2 = Subsequence(array[1:], append(subseq2, array[0]))
		}
	}

	if flagSecondRecursion && len(subseq2) > len(subseq1) {
		return subseq2
	}
	return subseq1
}
