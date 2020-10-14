// Package kdistinct contains Daily Coding Problem #13
//
// This problem was asked by Amazon.
//
// Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters.
//
// For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".
//
// Author: Andreas Atle, atle.andreas@gmail.com
package kdistinct

// Empty implements a byte-set
type Empty struct{}

// BytSet is a set of bytes
type ByteSet map[byte]Empty

// LongestSubString returns the length of the longest substring with atmost k characters.
func LongestSubstring(str string, k int) int {
	n := len(str)
	length := min(k, len(str))
	for i := 0; i <= n-k; i++ {
		m := ByteSet{}
		for j := i; ; j++ {
			if j < n {
				m[str[j]] = Empty{}
			}
			if len(m) > k || j == n {
				length = max(length, j-i)
				break
			}
		}
	}
	return length
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
