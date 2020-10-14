package validtree_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/validtree"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	var tree *validtree.Node

	// Check empty tree
	tree = nil
	assert.True(t, tree.Check())

	// Check a valid tree
	tree = &validtree.Node{5, &validtree.Node{3, nil, nil}, &validtree.Node{7, nil, nil}}
	assert.True(t, tree.Check())

	// Check an invalid tree
	tree = &validtree.Node{5, &validtree.Node{6, nil, nil}, &validtree.Node{7, nil, nil}}
	assert.False(t, tree.Check())

	// Check an invalid tree
	tree = &validtree.Node{5, &validtree.Node{3, nil, nil}, &validtree.Node{4, nil, nil}}
	assert.False(t, tree.Check())

	// Check a valid tree
	tree = &validtree.Node{5, &validtree.Node{3, nil, &validtree.Node{4, nil, nil}}, &validtree.Node{7, &validtree.Node{6, nil, nil}, nil}}
	assert.True(t, tree.Check())
}
