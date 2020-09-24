// Package inverttree contains Daily Coding Problem #83
//
// This problem was asked by Google.
//
// Invert a binary tree.
// For example, given the following tree:
//      a
//     / \
//    b   c
//   / \  /
//  d   e f
// should become:
//    a
//   / \
//   c  b
//   \  / \
//    f e  d
//
// Author: Andreas Atle, atle.andreas@gmail.com
package inverttree

// Node contains one node in a binary tree
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// String is a stringer routine for a tree with root-node n
func (n *Node) String() string {
	if n == nil {
		return "-"
	}

	if n.Left == nil && n.Right == nil {
		return "Leaf(" + n.Value + ")"
	}
	return "Node(" + n.Value + "," + n.Left.String() + "," + n.Right.String() + ")"
}

// Invert flips Left and Right in the whole tree
func (n *Node) Invert() {
	if n == nil {
		return
	}

	n.Left, n.Right = n.Right, n.Left
	n.Left.Invert()
	n.Right.Invert()
}
