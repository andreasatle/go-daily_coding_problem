package delpal_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/delpal"
	"github.com/stretchr/testify/assert"
)

func TestDelpal(t *testing.T) {

	// Case from problem formulation
	assert.True(t, delpal.IsKPalindrome("waterrfetawx", 2))

	// Reduce number of deletions
	assert.False(t, delpal.IsKPalindrome("waterrfetawx", 1))

	// Reduce number of deletions
	assert.False(t, delpal.IsKPalindrome("waterrfetawx", -1))

	// Empty string with no deletions
	assert.True(t, delpal.IsKPalindrome("", 0))

	// Empty string with one deletion
	assert.True(t, delpal.IsKPalindrome("", 1))

	// Case with one deletion
	assert.True(t, delpal.IsKPalindrome("qw", 1))

	// Case with one deletion
	assert.False(t, delpal.IsKPalindrome("qw", 0))
}
