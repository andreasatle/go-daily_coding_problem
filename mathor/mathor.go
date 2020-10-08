// Package mathor contains Daily Coding Problem #85
//
// This problem was asked by Facebook.
//
// Given three 32-bit integers x, y, and b, return x if b is 1 and y if b is 0,
// using only mathematical or bit operations. You can assume b can only be 1 or 0.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package mathor

import "errors"

// Ternary returns x if b is 1 and y if b is 0.
func Ternary(x, y, b int) (int, error) {
	if b != 0 && b != 1 {
		return 0, errors.New("b has to be 0 or 1")
	}
	return b*x | (1-b)*y, nil
}
