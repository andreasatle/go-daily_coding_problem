// Package validtree contains Daily Coding Problem #89
//
// This problem was asked by LinkedIn.
//
// Determine whether a tree is a valid binary search tree.
// A binary search tree is a tree with two children, left and right,
// and satisfies the constraint that the key in the left child must
// be less than or equal to the root and the key in the right child
// must be greater than or equal to the root.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package validtree

// Node contains one node in a binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Check return true if the tree with root n is an ordered binary tree.
func (n *Node) Check() bool {
	if n == nil {
		return true
	}

	if n.Left != nil && n.Left.Value > n.Value {
		return false
	}

	if n.Right != nil && n.Right.Value < n.Value {
		return false
	}

	return n.Left.Check() && n.Right.Check()
}
