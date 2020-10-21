// Package setpts contains Daily Coding Problem #119
//
// This problem was asked by Google.
//
// Given a set of closed intervals, find the smallest set of numbers that
// covers all the intervals. If there are multiple smallest sets, return
// any of them.
//
// For example, given the intervals [0, 3], [2, 6], [3, 4], [6, 9],
// one set of numbers that covers all these intervals is {3, 6}
//
// Author: Andreas Atle, atle.andreas@gmail.com
package setpts

import (
	"sort"
)

// Interval contains the endpoints of an integer interval
type Interval struct {
	Min int
	Max int
}

// Intervals contains a slice of Interval
type Intervals []Interval

// Local datatypes
type empty struct{}
type intSet map[int]empty
type intervalSet map[Interval]empty
type intMapIntervalSet map[int]intervalSet

// FindPoints find the minimal number of points that covers all closed intervals
func (iVals Intervals) FindPoints() []int {
	vals := iVals.endValues()
	iMapIvals := iVals.newIntMapInterval(vals)
	points := reduce(iMapIvals)
	sort.Ints(points)
	return points
}

func (iVals Intervals) endValues() intSet {
	iSet := intSet{}
	for _, iVal := range iVals {
		iSet[iVal.Min] = empty{}
		iSet[iVal.Max] = empty{}
	}
	return iSet
}
func (iVals Intervals) newIntMapInterval(iSet intSet) intMapIntervalSet {
	iMapIvals := intMapIntervalSet{}
	for val := range iSet {
		for _, iVal := range iVals {
			if val >= iVal.Min && val <= iVal.Max {
				if _, ok := iMapIvals[val]; !ok {
					iMapIvals[val] = intervalSet{}
				}
				iMapIvals[val][iVal] = empty{}
			}
		}
	}
	return iMapIvals
}

func reduce(m intMapIntervalSet) []int {
	for k1, v1 := range m {
		for k2, v2 := range m {
			if k1 == k2 {
				continue
			}
			iValSet := intervalSet{}
			for ival := range v1 {
				iValSet[ival] = empty{}
			}
			for ival := range v2 {
				iValSet[ival] = empty{}
			}
			if len(v1) == len(iValSet) {
				delete(m, k2)
			}
			if len(v2) == len(iValSet) {
				delete(m, k1)
			}
		}
	}
	out := []int{}
	for val := range m {
		out = append(out, val)
	}
	return out
}
