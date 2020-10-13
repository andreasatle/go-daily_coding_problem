// package unival contains Daily Coding Problem #8
//
// This problem was asked by Google.
//
// A unival tree (which stands for "universal value")
// is a tree where all nodes under it have the same value.
// Given the root to a binary tree, count the number of unival subtrees.
// For example, the following tree has 5 unival subtrees:
//    0
//   / \
//  1   0
//     / \
//    1   0
//   / \
//  1   1
//
// Author: Andreas Atle, atle.andreas@gmail.com
package unival

type Value int

type Node struct {
	Val   Value
	Left  *Node
	Right *Node
}

func (n *Node) Count() int {
	if n == nil {
		return 0
	}

	_, _, out := n.count(0)
	return out
}

func (n *Node) count(val Value) (bool, Value, int) {
	if n == nil {
		return true, val, 0
	}

	leftFlag, leftVal, leftCnt := n.Left.count(n.Val)
	rightFlag, rightVal, rightCnt := n.Right.count(n.Val)

	nodeCnt := leftCnt + rightCnt
	if leftFlag && rightFlag && leftVal == n.Val && rightVal == n.Val {
		return true, n.Val, nodeCnt + 1
	}
	return false, n.Val, nodeCnt
}
