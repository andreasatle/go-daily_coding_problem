// Package levelwise contains Daily Coding Problem #107
//
// This problem was asked by Microsoft.
//
// Print the nodes in a binary tree level-wise. For example, the following
// should print 1, 2, 3, 4, 5.
//
//     1
//    / \
//   2   3
//      / \
//     4   5
//
// Author: Andreas Atle, atle.andreas@gmail.com
package levelwise

// Value contains the type of element in the binary tree
type Value int

// Node contains one node of the binary tree
type Node struct {
	Val   Value
	Left  *Node
	Right *Node
}

// Print returns the values of the tree in a breadth-first order
func (n *Node) Print() []Value {
	if n == nil {
		return []Value{}
	}
	return n.print([]Value{n.Val})
}

// print is a recursive help-function for Print
func (n *Node) print(vals []Value) []Value {
	if n == nil {
		return vals
	}

	hasLeft := n.Left != nil
	hasRight := n.Right != nil

	if hasLeft {
		vals = append(vals, n.Left.Val)
	}

	if hasRight {
		vals = append(vals, n.Right.Val)
	}

	if hasLeft {
		vals = n.Left.print(vals)
	}

	if hasRight {
		vals = n.Right.print(vals)
	}

	return vals
}
