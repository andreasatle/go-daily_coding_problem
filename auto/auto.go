// Package auto contains Daily Coding Problem #11
//
// This problem was asked by Twitter.
//
// Implement an autocomplete system. That is, given a query string s and a set of all
// possible query strings, return all strings in the set that have s as a prefix.
//
// For example, given the query string de and the set of strings [dog, deer, deal],
// return [deer, deal].
//
// Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package auto

// AutoComplete is an acronym for a map used in autocomplete
type AutoComplete map[string][]string

// NewAutoComplete returns a map used in autocomplete
func NewAutoComplete(words []string) AutoComplete {
	output := AutoComplete{}
	for _, word := range words {
		for i := 0; i <= len(word); i++ {
			w := word[:i]
			if _, ok := output[w]; !ok {
				output[w] = []string{}
			}
			output[w] = append(output[w], word)
		}
	}
	return output
}

// Complete returns all words in corpus that start with prefix
func (a AutoComplete) Complete(prefix string) []string {
	return a[prefix]
}
