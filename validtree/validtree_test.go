package validtree_test

import (
	"testing"

	"github.com/andreasatle/Assorted/Daily/validtree"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	var tree *validtree.Node

	// Check empty tree
	tree = nil
	assert.Equal(t, true, tree.Check())

	// Check a valid tree
	tree = &validtree.Node{5, &validtree.Node{3, nil, nil}, &validtree.Node{7, nil, nil}}
	assert.Equal(t, true, tree.Check())

	// Check an invalid tree
	tree = &validtree.Node{5, &validtree.Node{6, nil, nil}, &validtree.Node{7, nil, nil}}
	assert.Equal(t, false, tree.Check())

	// Check a valid tree
	tree = &validtree.Node{5, &validtree.Node{3, nil, &validtree.Node{4, nil, nil}}, &validtree.Node{7, &validtree.Node{6, nil, nil}, nil}}
	assert.Equal(t, true, tree.Check())
}
