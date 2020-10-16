// Package runmed contains Daily Coding Problem #33
//
// This problem was asked by Microsoft.
//
// Compute the running median of a sequence of numbers. That is, given a
// stream of numbers, print out the median of the list so far on each
// new element.
//
// Recall that the median of an even-numbered list is the average of the
// two middle numbers.
//
// For example, given the sequence [2, 1, 5, 7, 2, 0, 5], your algorithm
// should print out:
//  2
//  1.5
//  2
//  3.5
//  2
//  2
//  2
//
// Author: Andreas Atle, atle.andreas@gmail.com
package runmed

// Sequence stores all inserted values
type Sequence struct {
	vals []float64
}

// NewSequence returns pointer to an empty Sequence
func NewSequence() *Sequence {
	return &Sequence{}
}

// Insert a value and return the running median
func (seq *Sequence) Insert(val float64) float64 {
	seq.insert(val)
	return seq.median()
}

// insert puts val in the correct place so slice is ordered
func (seq *Sequence) insert(val float64) {
	// Insert val last in slice
	seq.vals = append(seq.vals, val)

	// Swap backwards until val is in correct place
	for i := len(seq.vals) - 2; i >= 0; i-- {
		if seq.vals[i] <= seq.vals[i+1] {
			break
		}
		seq.vals[i], seq.vals[i+1] = seq.vals[i+1], seq.vals[i]
	}
}

// median computes the median (duh)
func (seq *Sequence) median() float64 {
	n := len(seq.vals)
	// Use that n/2 == (n-1)/2 for odd numbers (no need for if-statement)
	return 0.5 * (seq.vals[n/2] + seq.vals[(n-1)/2])
}
