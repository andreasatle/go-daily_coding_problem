package inverttree_test

import (
	"testing"

	"github.com/andreasatle/Assorted/Daily/inverttree"
	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	myTree := &inverttree.Node{
		Value: "a",
		Left: &inverttree.Node{
			Value: "b",
			Left: &inverttree.Node{
				Value: "d",
				Left:  nil,
				Right: nil,
			},
			Right: &inverttree.Node{
				Value: "e",
				Left:  nil,
				Right: nil,
			},
		},
		Right: &inverttree.Node{
			Value: "c",
			Left: &inverttree.Node{
				Value: "f",
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}}

	expected := "Node(a,Node(c,-,Leaf(f)),Node(b,Leaf(e),Leaf(d)))"
	myTree.Invert()
	actual := myTree.String()
	assert.Equal(t, expected, actual, "Inversion of tree failed.")
}
