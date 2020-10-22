// Package levelsum contains Daily Coding Problem #117
//
// This problem was asked by Facebook.
//
// Given a binary tree, return the level of the tree with minimum sum.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package levelsum

import "errors"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// LevelOfMinSum returns the level with the minimum sum.
// If the tree is empty an error "empty tree" is flagged.
func (n *Node) LevelOfMinSum() (int, error) {
	nLevels := n.depth(0)
	if nLevels == 0 {
		return -1, errors.New("empty tree")
	}
	levelSum := make([]int, nLevels)
	n.levelOfMinSum(levelSum, 0)
	indOut := 0
	maxSum := levelSum[0]
	for i := 1; i < nLevels; i++ {
		if levelSum[i] < maxSum {
			maxSum = levelSum[i]
			indOut = i
		}
	}
	return indOut, nil
}

func (n *Node) levelOfMinSum(maxSum []int, level int) {
	if n == nil {
		return
	}

	maxSum[level] += n.Value
	n.Left.levelOfMinSum(maxSum, level+1)
	n.Right.levelOfMinSum(maxSum, level+1)
}

func (n *Node) depth(level int) int {
	if n == nil {
		return level
	}

	level++
	return max(n.Left.depth(level), n.Right.depth(level))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
