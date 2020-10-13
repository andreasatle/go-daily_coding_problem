// Package leaf contains Daily Coding Problem #110
//
// This problem was asked by Facebook.
//
// Given a binary tree, return all paths from the root to leaves.
//
//For example, given the tree:
//
//     1
//    / \
//   2   3
//      / \
//     4   5
// Return
//   [[1, 2], [1, 3, 4], [1, 3, 5]].
//
// Author: Andreas Atle, atle.andreas@gmail.com
package leaf

// Value is the type of the element in the binary tree
type Value int

// Node contains one node in the binary tree
type Node struct {
	Val   Value
	Left  *Node
	Right *Node
}

// Print returns the paths from the root to all leaves
func (n *Node) Print() [][]Value {
	if n.isEmpty() {
		return [][]Value{}
	}

	return n.print([]Value{}, [][]Value{})
}

// Recursive help-function for public function Print
func (n *Node) print(current []Value, all [][]Value) [][]Value {
	// Apply back-tracing (with defer)
	current = append(current, n.Val)
	defer func() { current = current[:len(current)-1] }()

	// Base-case add current to all
	if n.isLeaf() {
		return append(all, current)
	}

	// Recursion step
	all = n.Left.print(current, all)
	all = n.Right.print(current, all)

	return all
}

// isEmpty returns true if the node is empty (equal nil)
func (n *Node) isEmpty() bool {
	return n == nil
}

// isLeaf returns true if the node is a (non-empty) leaf
func (n *Node) isLeaf() bool {
	return !n.isEmpty() && n.Left == nil && n.Right == nil
}
