package levelwise_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/levelwise"
	"github.com/stretchr/testify/assert"
)

func TestLevelWise(t *testing.T) {
	n2 := &levelwise.Node{2, nil, nil}
	n4 := &levelwise.Node{4, nil, nil}
	n5 := &levelwise.Node{5, nil, nil}
	n3 := &levelwise.Node{3, n4, n5}
	n1 := &levelwise.Node{1, n2, n3}

	vals := n1.Print()
	// Check that the output length is 5
	assert.Equal(t, 5, len(vals))

	// Check that vals == []Value{1, 2, 3, 4, 5}
	assert.Equal(t, levelwise.Value(1), vals[0])
	assert.Equal(t, levelwise.Value(2), vals[1])
	assert.Equal(t, levelwise.Value(3), vals[2])
	assert.Equal(t, levelwise.Value(4), vals[3])
	assert.Equal(t, levelwise.Value(5), vals[4])

	var emptyTree *levelwise.Node
	// Check that length of output from empty tree is zero
	assert.Equal(t, 0, len(emptyTree.Print()))
}
