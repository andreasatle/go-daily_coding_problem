// Package minpath contains Daily Coding Problem #100
//
// This problem was asked by Google.
//
// You are in an infinite 2D grid where you can move in any of the 8 directions:
//
//  (x,y) to
//    (x+1, y  ),
//    (x-1, y  ),
//    (x  , y+1),
//    (x  , y-1),
//    (x-1, y-1),
//    (x+1, y+1),
//    (x-1, y+1),
//    (x+1, y-1)
// You are given a sequence of points and the order in which you need to cover the points.
// Give the minimum number of steps in which you can achieve it. You start from the first point.
//
// Example:
//  Input: [(0, 0), (1, 1), (1, 2)]
//  Output: 2
// It takes 1 step to move from (0, 0) to (1, 1). It takes one more step to move from (1, 1) to (1, 2).
//
// Author: Andreas Atle, atle.andreas@gmail.com
package minpath

// Point contains a 2D coordinate (X,Y)
type Point struct {
	X int
	Y int
}

// MinSteps returns the minimal number of steps passing through all points
func MinSteps(points []Point) int {
	if len(points) < 2 {
		return 0
	}
	return minSteps(points[0], points[1:])
}

// minSteps is a recursive helpfunction for MinSteps
func minSteps(first Point, rest []Point) int {
	if len(rest) == 0 {
		return 0
	}

	return distance(first, rest[0]) + minSteps(rest[0], rest[1:])
}

// distance returns the maximum distance along all the axis directions
func distance(p1, p2 Point) int {
	return max(abs(p1.X-p2.X), abs(p1.Y-p2.Y))
}

// abs returns the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// max returns the maximum value
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
