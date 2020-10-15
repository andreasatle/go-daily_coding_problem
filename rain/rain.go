// Package rain contains Daily Coding Problem #30
//
// This problem was asked by Facebook.
//
// You are given an array of non-negative integers that represents a
// two-dimensional elevation map where each element is unit-width wall
// and the integer is the height. Suppose it will rain and all spots
// between two walls get filled up.
//
// Compute how many units of water remain trapped on the map
// in O(N) time and O(1) space.
//
// For example, given the input [2, 1, 2], we can hold 1 unit of water
// in the middle.
//
// Given the input [3, 0, 1, 3, 0, 5], we can hold 3 units in the
// first index, 2 in the second, and 3 in the fourth index (we cannot
// hold 5 since it would run off to the left), so we can trap 8 units of water.
//
// Note: The problem is solved with a forward sweep (fillFwd) that will detect
// the last highest peak it the topography. I finish with a backward sweep
// (fillBwd) where we know that the first entry is the largest. Therefore
// the recursion terminates in two steps.
// Atleast it works for several non-trivial test-cases, and should
// satisfy the complexity requirements O(N) time, O(1) space.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package rain

// Fill returns the amount of rainwater that can fill the topography.
func Fill(topo []int) int {
	return fillFwd(topo)
}

// fillFwd is a forward sweep help function to Fill.
func fillFwd(topo []int) int {
	n := len(topo)
	if n <= 2 {
		return 0
	}

	output := 0
	iMax, iCurr := 0, 1
	for ; iCurr < n; iCurr++ {
		if topo[iCurr] >= topo[iMax] {
			output += accTopo(topo[iMax+1:iCurr], topo[iMax])
			iMax = iCurr
		}
	}

	return output + fillBwd(topo[iMax:iCurr])
}

// fillBwd is a backward sweep help function to Fill.
func fillBwd(topo []int) int {
	n := len(topo)
	if n <= 2 {
		return 0
	}

	output := 0
	iMax, iCurr := n-1, n-2
	for ; iCurr >= 0; iCurr-- {
		if topo[iCurr] >= topo[iMax] {
			output += accTopo(topo[iCurr+1:iMax], topo[iMax])
			iMax = iCurr
		}
	}

	return output
}

// accTopo returns the water volume assuming that
// topo is the waterbottom, with water at max.
func accTopo(topo []int, max int) int {
	acc := 0
	for _, t := range topo {
		acc += max - t
	}
	return acc
}
