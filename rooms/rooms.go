// Package rooms contains Daily Coding Problem #21
//
// This problem was asked by Snapchat.
//
// Given an array of time intervals (start, end) for classroom lectures
// (possibly overlapping), find the minimum number of rooms required.
//
// For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package rooms

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

type tuple struct {
	time  int
	value int
}

type tuples []tuple

// Minimum returns the minimum number of classrooms that has to be booked,
// given the interval of each class.
func Minimum(ivals Intervals) int {
	tups := make(tuples, 2*len(ivals))
	for i, ival := range ivals {
		tups[2*i] = tuple{ival.Start, 1}
		tups[2*i+1] = tuple{ival.End, -1}
	}
	bubbleSort(tups)
	return maxAccSum(tups)
}

// bubbleSort is a help function for Minimum, to avoid writing interface routines for package sort
func bubbleSort(tups tuples) {
	n := len(tups) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if tups[j].time > tups[j+1].time {
				tups[j], tups[j+1] = tups[j+1], tups[j]
			}
		}
	}
}

// maxAccSum returns the maximum of the accumulated sum of the tuple-values
func maxAccSum(tups tuples) int {
	max := 0
	acc := 0
	for _, tup := range tups {
		acc += tup.value
		if acc > max {
			max = acc
		}
	}
	return max
}
