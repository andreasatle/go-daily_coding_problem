// Package subtree contains Daily Coding Problem #115
//
// This problem was asked by Google.
//
// Given two non-empty binary trees s and t, check whether tree t has
// exactly the same structure and node values with a subtree of s. A
// subtree of s is a tree consists of a node in s and all of this
// node's descendants. The tree s could also be considered as a
// subtree of itself.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package subtree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// IsSubtree returns true if t is a subtree of s
func (s *Node) IsSubtree(t *Node) bool {
	return s.isSubtree(t, t)
}

func (s *Node) isSubtree(tRoot, t *Node) bool {
	if s == nil {
		return t == nil
	}
	if t == nil {
		return false
	}

	// Check is values match and both subtrees match
	valuesMatch := s.Value == t.Value && s.Left.isSubtree(tRoot, t.Left) && s.Right.isSubtree(tRoot, t.Right)

	// Go through each possible subtree
	checkSubtrees := s.Left.isSubtree(tRoot, tRoot) || s.Right.isSubtree(tRoot, tRoot)

	return valuesMatch || checkSubtrees
}
