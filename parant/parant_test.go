package parant_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/parant"
	"github.com/stretchr/testify/assert"
)

func TestParantCase1(t *testing.T) {
	// Case from formulation
	assert.Equal(t, 2, parant.RemoveCount(")("))
}

func TestParantCase2(t *testing.T) {
	// Case from formulation
	assert.Equal(t, 1, parant.RemoveCount("()())()"))
}

func TestParantCase3(t *testing.T) {
	// Valid case
	assert.Equal(t, 0, parant.RemoveCount("Node(Foo(),Bar())"))
}

func TestParantCase4(t *testing.T) {
	// Added two parantheses to valid case 3
	assert.Equal(t, 2, parant.RemoveCount("Node(Foo()),Bar()))"))
}
