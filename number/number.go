// Package number contains Daily Coding Problem #123
//
// This problem was asked by LinkedIn.
//
// Given a string, return whether it represents a number.
// Here are the different kinds of numbers:
//  "10", a positive integer
//  "-10", a negative integer
//  "10.1", a positive real number
//  "-10.1", a negative real number
//  "1e5", a number in scientific notation
// And here are examples of non-numbers:
//  "a"
//  "x 1"
//  "a -2"
//  "-"
//
// Note: This problem is trivial in golang, since there are builtin
// routines that do the decoding. Why reinvent the wheel...
//
// Author: Andreas Atle, atle.andreas@gmail.com
package number

import (
	"strconv"
)

// Check if the string argument can be decoded as a number
func Check(str string) bool {
	// Check for a float64, scientific or not, integer or not
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
