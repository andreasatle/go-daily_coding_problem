// Package stringshift contains Daily Coding Problem #108
//
// This problem was asked by Google.
//
// Given two strings A and B, return whether or not A can be
// shifted some number of times to get B.
//
// For example, if A is abcde and B is cdeab, return true.
// If A is abc and B is acb, return false.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package stringshift

// ShiftedEqual returns true if there exists a circular shift that makes the strings equal
func ShiftedEqual(s1, s2 string) bool {
	// Check that strings are of same length
	if len(s1) != len(s2) {
		return false
	}
	// n is the length of the two strings
	n := len(s1)

	// Check all circular shifts and return true if match is found
	for i := 0; i < n; i++ {
		if s1[:i] == s2[n-i:] && s1[i:] == s2[:n-i] {
			return true
		}
	}

	// Return false except if both strings are empty
	return n == 0
}
