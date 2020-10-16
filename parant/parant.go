// Package parant contains Daily Coding Problem #86
//
// This problem was asked by Google.
//
// Given a string of parentheses, write a function to compute the minimum
// number of parentheses to be removed to make the string valid
// (i.e. each open parenthesis is eventually closed).
//
// For example, given the string "()())()", you should return 1. Given the
// string ")(", you should return 2, since we must remove all of them.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package parant

// RemoveCount returns the number of paranthesis to be removed to get a
// valid expression
func RemoveCount(str string) int {
	removeCnt := 0
	paraCnt := 0
	for _, s := range str {
		if s == '(' {
			paraCnt++
		}
		if s == ')' {
			if paraCnt == 0 {
				paraCnt = 0
				removeCnt++
				continue
			}
			paraCnt--
		}
	}
	return removeCnt + paraCnt
}
