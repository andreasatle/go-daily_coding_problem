// Package color contains Daily Coding Problem #19
//
// This problem was asked by Facebook.
//
// A builder is looking to build a row of N houses that can be of K different
// colors. He has a goal of minimizing cost while ensuring that no two
// neighboring houses are of the same color.
//
// Given an N by K matrix where the nth row and kth column represents the cost
// to build the nth house with kth color, return the minimum cost which
// achieves this goal.
//
// Note: This is a dynamic programming problem. The colors are also tracked
// as a bonus. This makes the code slightly more complicated. The colors
// for the optimal cost is not always unique. We return one of the possible
// solutions.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package color

import (
	"math"
)

// DPTuple contains a (cost,index)-tuple
type DPTuple struct {
	Cost  float64
	Index int
}

// Color returns the minimal cost and one possible color-combination.
func Color(cost [][]float64) (float64, []int) {
	n := len(cost)
	if n == 0 {
		return 0, []int{}
	}

	k := len(cost[0])
	dp := newDPTuple(cost)
	for house := 1; house < n; house++ {
		for color := 0; color < len(dp[house]); color++ {

			i := minIndex(dp[house-1], color)
			newCost := dp[house-1][i].Cost + cost[house][color]
			dp[house][color] = DPTuple{Cost: newCost, Index: i}
		}
	}

	index := make([]int, n)
	finalCost := dp[n-1][0].Cost
	for j := 1; j < k; j++ {
		if dp[n-1][j].Cost < finalCost {
			finalCost = dp[n-1][j].Cost
			index[n-1] = j
		}
	}
	for house := n - 2; house >= 0; house-- {
		index[house] = dp[house+1][index[house+1]].Index
	}
	return finalCost, index
}

// newDPTuple returns a matrix o
func newDPTuple(cost [][]float64) [][]DPTuple {
	dp := make([][]DPTuple, len(cost))
	for j := 0; j < len(cost); j++ {
		dp[j] = make([]DPTuple, len(cost[0]))
	}
	for j := 0; j < len(cost[0]); j++ {
		dp[0][j] = DPTuple{Cost: cost[0][j], Index: -1}
	}
	return dp
}

func minIndex(dp []DPTuple, except int) int {
	minDp := DPTuple{Cost: math.MaxFloat64, Index: -1}
	for i, tup := range dp {
		if i == except {
			continue
		}
		if tup.Cost < minDp.Cost {
			minDp = DPTuple{Cost: tup.Cost, Index: i}
		}
	}
	return minDp.Index
}
