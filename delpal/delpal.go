// Package delpal contains Daily Coding Problem #121
//
// This problem was asked by Google.
//
// Given a string which we can delete at most k, return whether you can
// make a palindrome.
//
// For example, given 'waterrfetawx' and a k of 2, you could delete f and x
// to get 'waterretaw'.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package delpal

func IsKPalindrome(str string, k int) bool {

	// Too many chars deleted
	if k < 0 {
		return false
	}

	// Palindrome detected
	if len(str) <= 1 {
		return true
	}

	n := len(str)

	// No deletion needed
	if str[0] == str[n-1] {
		return IsKPalindrome(str[1:n-1], k)
	}

	// delete first and last char respectively
	return IsKPalindrome(str[1:], k-1) || IsKPalindrome(str[:n-1], k-1)
}
