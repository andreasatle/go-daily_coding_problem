package palin_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/palin"
	"github.com/stretchr/testify/assert"
)

func TestPalinTrue(t *testing.T) {
	list := palin.NewDoubleLinkedList()
	list.InsertLast(3)
	list.InsertLast(4)
	list.InsertLast(5)
	list.InsertLast(4)
	list.InsertLast(3)
	assert.Equal(t, true, list.Palindrome())
	assert.Equal(t, true, list.PalindromeSingle())
}

func TestPalinFalse(t *testing.T) {
	list := palin.NewDoubleLinkedList()
	list.InsertLast(3)
	list.InsertLast(4)
	list.InsertLast(5)
	list.InsertLast(3)
	assert.Equal(t, false, list.Palindrome())
	assert.Equal(t, false, list.PalindromeSingle())
}
