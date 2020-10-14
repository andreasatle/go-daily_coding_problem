// Package strtree contains Daily Coding Problem #3
//
// This problem was asked by Google.
//
// Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
// For example, given the following Node class
//  class Node:
//     def __init__(self, val, left=None, right=None):
//         self.val = val
//         self.left = left
//         self.right = right
// The following test should pass:
//  node = Node('root', Node('left', Node('left.left')), Node('right'))
//  assert deserialize(serialize(node)).left.left.val == 'left.left'
//
// Author: Andreas Atle, atle.andreas@gmail.com
package strtree

import (
	"errors"
	"strings"
)

// Node contains one node of a binary tree
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// Serialize is a stringer for a binary tree
func (n *Node) Serialize() string {
	if n == nil {
		return "-"
	}
	return "Node(" + n.Value + "," + n.Left.Serialize() + "," + n.Right.Serialize() + ")"
}

// Deserialize is a string to binary tree converter
func Deserialize(str string) (*Node, error) {
	if strings.HasPrefix(str, "-") {
		return nil, nil
	}
	if !strings.HasPrefix(str, "Node(") {
		return nil, errors.New("wrong prefix")
	}

	findLeftRightSplit := func(str string) int {
		split, count := 0, 0
		for i, c := range str {
			if c == '(' {
				count++
			} else if c == ')' {
				count--
			} else if c == ',' && count == 0 {
				split = i
			}
		}
		return split
	}

	split := strings.Index(str, ",")
	val := str[5:split]
	stripped := str[split+1 : len(str)-1]

	split = findLeftRightSplit(stripped)
	left := stripped[:split]
	right := stripped[split+1:]

	leftDes, leftErr := Deserialize(left)
	if leftErr != nil {
		return nil, leftErr
	}
	rightDes, rightErr := Deserialize(right)
	if rightErr != nil {
		return nil, rightErr
	}

	return &Node{val, leftDes, rightDes}, nil
}

// Cmp returns true if n and n2 are equal.
func (n *Node) Cmp(n2 *Node) bool {
	if n == nil && n2 == nil {
		return true
	}

	if (n == nil && n2 != nil) || (n != nil && n2 == nil) {
		return false
	}

	if n.Value != n2.Value {
		return false
	}

	return n.Left.Cmp(n2.Left) && n.Right.Cmp(n2.Right)
}
