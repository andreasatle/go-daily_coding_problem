package strtree_test

import (
	"errors"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/strtree"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	var tree, tree2, tree3 *strtree.Node
	var err error

	// Example from problem formulation
	tree = &strtree.Node{"root", &strtree.Node{"root->left", nil, nil}, &strtree.Node{"root->right", nil, &strtree.Node{"root->right->right", nil, nil}}}
	assert.True(t, tree.Cmp(tree))
	assert.False(t, tree.Cmp(nil))

	// Check Serialize and Deserialize
	tree2, err = strtree.Deserialize(tree.Serialize())
	assert.True(t, tree.Cmp(tree2))

	// Check deserialize of invalid serialization
	ser1 := "Node(root,None(root->left,-,-),Node(root->right,-,Node(root->right->right,-,-)))"
	ser2 := "Node(root,Node(root->left,-,-),None(root->right,-,Node(root->right->right,-,-)))"
	_, err1 := strtree.Deserialize(ser1)
	_, err2 := strtree.Deserialize(ser2)
	assert.Equal(t, errors.New("wrong prefix"), err1)
	assert.Equal(t, errors.New("wrong prefix"), err2)

	// Check deserialize of invalid serialization again
	tree3 = &strtree.Node{"root", &strtree.Node{"root->left", nil, nil}, &strtree.Node{"root->right", nil, &strtree.Node{"root->right->righ", nil, nil}}}
	assert.False(t, tree.Cmp(tree3))

	// Check deserialize of invalid serialization one last time
	_, err = strtree.Deserialize("Nobe(root,-,-)")
	assert.Equal(t, errors.New("wrong prefix"), err)
}
