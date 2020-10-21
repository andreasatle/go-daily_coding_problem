package kthlast_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/kthlast"
	"github.com/stretchr/testify/assert"
)

func TestKthlastCase1(t *testing.T) {
	// Remove 2nd from last
	ll := &kthlast.LinkedList{&kthlast.Node{1, &kthlast.Node{2, &kthlast.Node{3, &kthlast.Node{4, &kthlast.Node{5, nil}}}}}}
	out := ll.RemoveKthLast(2).Slice()
	assert.Equal(t, 4, len(out))
	assert.Equal(t, 1, out[0])
	assert.Equal(t, 2, out[1])
	assert.Equal(t, 3, out[2])
	assert.Equal(t, 5, out[3])
}

func TestKthlastCase2(t *testing.T) {
	// Remove first
	ll := &kthlast.LinkedList{&kthlast.Node{1, &kthlast.Node{2, &kthlast.Node{3, &kthlast.Node{4, &kthlast.Node{5, nil}}}}}}
	out := ll.RemoveKthLast(5).Slice()
	assert.Equal(t, 4, len(out))
	assert.Equal(t, 2, out[0])
	assert.Equal(t, 3, out[1])
	assert.Equal(t, 4, out[2])
	assert.Equal(t, 5, out[3])
}

func TestKthlastCase3(t *testing.T) {
	// Remove out of range
	ll := &kthlast.LinkedList{&kthlast.Node{1, &kthlast.Node{2, &kthlast.Node{3, &kthlast.Node{4, &kthlast.Node{5, nil}}}}}}
	out := ll.RemoveKthLast(0).Slice()
	assert.Equal(t, 5, len(out))
	assert.Equal(t, 1, out[0])
	assert.Equal(t, 2, out[1])
	assert.Equal(t, 3, out[2])
	assert.Equal(t, 4, out[3])
	assert.Equal(t, 5, out[4])
}

func TestKthlastCase4(t *testing.T) {
	// Remove out of range
	ll := &kthlast.LinkedList{&kthlast.Node{1, &kthlast.Node{2, &kthlast.Node{3, &kthlast.Node{4, &kthlast.Node{5, nil}}}}}}
	out := ll.RemoveKthLast(6).Slice()
	assert.Equal(t, 5, len(out))
	assert.Equal(t, 1, out[0])
	assert.Equal(t, 2, out[1])
	assert.Equal(t, 3, out[2])
	assert.Equal(t, 4, out[3])
	assert.Equal(t, 5, out[4])
}

func TestKthlastCase5(t *testing.T) {
	// Empty list
	ll := &kthlast.LinkedList{}
	out := ll.RemoveKthLast(6).Slice()
	assert.Equal(t, 0, len(out))
}
