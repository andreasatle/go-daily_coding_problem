package strtree_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/strtree"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	var tree, tree2, tree3 *strtree.Node

	tree = &strtree.Node{"root", &strtree.Node{"root->left", nil, nil}, &strtree.Node{"root->right", nil, &strtree.Node{"root->right->right", nil, nil}}}
	tree2 = strtree.Deserialize(tree.Serialize())
	tree3 = &strtree.Node{"root", &strtree.Node{"root->left", nil, nil}, &strtree.Node{"root->right", nil, &strtree.Node{"root->right->righ", nil, nil}}}

	// Check Cmp function
	assert.Equal(t, false, tree.Cmp(tree3))
	assert.Equal(t, true, tree.Cmp(tree))

	// Check Serialize and Deserialize
	assert.Equal(t, true, tree.Cmp(tree2))

	tree = nil
	tree2 = strtree.Deserialize(tree.Serialize())
	assert.Equal(t, true, tree.Cmp(tree2))
}
