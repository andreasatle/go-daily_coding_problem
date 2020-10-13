// Package anagram contains Daily Coding Problem #111
//
// This problem was asked by Google.
//
// Given a word W and a string S, find all starting indices in S
// which are anagrams of W.
//
// For example, given that W is "ab", and S is "abxaba",
// return 0, 3, and 4.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package anagram

import (
	"sort"
)

// Runes is an acronym for []rune
type Runes []rune

// Len is an interface routine for sorting
func (b Runes) Len() int { return len(b) }

// Swap is an interface routine for sorting
func (b Runes) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less is an interface routine for sorting
func (b Runes) Less(i, j int) bool { return b[i] < b[j] }

// GetIndices returns the starting index in s to all anagrams os w
func GetIndices(s, w string) []int {
	indices := []int{}
	w2 := sortString(w)
	for i := 0; i <= len(s)-len(w2); i++ {
		s2 := sortString(s[i : i+len(w2)])
		if w2 == s2 {
			indices = append(indices, i)
		}
	}
	return indices
}

// sortString sorts a string using interface routines
func sortString(s string) string {
	b := Runes(s)
	sort.Sort(Runes(b))
	return string(b)
}
