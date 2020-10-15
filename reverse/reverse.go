// Package reverse contains Daily Coding Problem #113
//
// This problem was asked by Google.
//
// Given a string of words delimited by spaces, reverse the words in string.
// For example, given "hello world here", return "here world hello"
//
// Follow-up: given a mutable string representation, can you perform
// this operation in-place?
//
// Note: Since golang strings are immutable, we instead use []rune that are
// mutable.
// The slice is first reversed, so that all words appear at the correct location
// but reversed.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package reverse

// Reverse reverses the words in the mutable []rune that is similar to a string
// (that is immutable).
func Reverse(words []rune) {
	n := len(words)
	// Start by reversing the whole slice. Each word is reversed but backwards.
	reverse(words)

	// Reverse each word separately. Variable last has to be declared
	// in the first loop, since it is used to update first.
	for first, last := 0, 0; first < n; first = last + 1 {

		// Make sure that last is the first index after the current word.
		for last = first; last < n && words[last] != ' '; last++ {
		}

		reverse(words[first:last])
	}
}

func reverse(r []rune) {
	for first, last := 0, len(r)-1; first < last; first, last = first+1, last-1 {
		r[first], r[last] = r[last], r[first]
	}
}
