package leaf_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/leaf"
	"github.com/stretchr/testify/assert"
)

func TestLeaf(t *testing.T) {
	n2 := &leaf.Node{2, nil, nil}
	n4 := &leaf.Node{4, nil, nil}
	n5 := &leaf.Node{5, nil, nil}
	n3 := &leaf.Node{3, n4, n5}
	tree := &leaf.Node{1, n2, n3}

	out := tree.Print()

	assert.Equal(t, leaf.Value(1), out[0][0])
	assert.Equal(t, leaf.Value(2), out[0][1])

	assert.Equal(t, leaf.Value(1), out[1][0])
	assert.Equal(t, leaf.Value(3), out[1][1])
	assert.Equal(t, leaf.Value(4), out[1][2])

	assert.Equal(t, leaf.Value(1), out[2][0])
	assert.Equal(t, leaf.Value(3), out[2][1])
	assert.Equal(t, leaf.Value(5), out[2][2])

	// Add test with empty tree to get full test coverage
	var tree2 *leaf.Node
	out2 := tree2.Print()
	assert.Equal(t, 0, len(out2))
}
