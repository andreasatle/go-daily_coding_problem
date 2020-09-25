// Package schedule contains Daily Coding Problem #10
//
//This problem was asked by Apple.
//
// Implement a job scheduler which takes in a function f
// and an integer n, and calls f after n milliseconds.
//
// Note: This is very ugly in golang, which is strongly typed,
// and doesn't have generics. We need plenty of duplication.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package schedule

import (
	"time"
)

// Func is an alias for a function without args or return value
type Func func()

// FuncIntToInt is an alias for a function Z => Z
type FuncIntToInt func(int) int

// Job is a job scheduler for a function Func
func Job(f Func, n int) Func {
	time.Sleep(time.Duration(n) * time.Millisecond)
	return f
}

// JobIntToInt is a job scheduler for a function FuncIntToInt
func JobIntToInt(f FuncIntToInt, n int) FuncIntToInt {
	time.Sleep(time.Duration(n) * time.Millisecond)
	return f
}
