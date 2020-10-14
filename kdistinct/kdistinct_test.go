package kdistinct_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/kdistinct"
	"github.com/stretchr/testify/assert"
)

func TestKDistinct(t *testing.T) {
	// Case from problem formulation
	assert.Equal(t, 3, kdistinct.LongestSubstring("abcba", 2))

	// Check that full string is used
	assert.Equal(t, 5, kdistinct.LongestSubstring("abcba", 3))

	// Check when k is larger than size of string
	assert.Equal(t, 3, kdistinct.LongestSubstring("abc", 7))
}
